package server

import (
	"context"
	"strconv"
	"time"

	api2 "github.com/djcopley/zing/internal/api"
	model2 "github.com/djcopley/zing/internal/model"
	service2 "github.com/djcopley/zing/internal/service"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var _ api2.ZingServer = &Server{}

func NewServer(logger *zap.Logger, authService *service2.AuthenticationService, messageService *service2.MessageService) *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			newLoggingInterceptor(logger),
			newAuthInterceptor(authService),
		),
	)
	zingServer := &Server{
		authService:    authService,
		messageService: messageService,
		logger:         logger,
	}
	api2.RegisterZingServer(grpcServer, zingServer)
	return grpcServer
}

type Server struct {
	authService    *service2.AuthenticationService
	messageService *service2.MessageService
	logger         *zap.Logger
	api2.UnimplementedZingServer
}

func (s *Server) ClearMessages(ctx context.Context, request *api2.ClearMessagesRequest) (*api2.ClearMessagesResponse, error) {
	user := getUserFromContext(ctx)
	if err := s.messageService.ClearMessages(user.Username); err != nil {
		return &api2.ClearMessagesResponse{}, err
	}
	return &api2.ClearMessagesResponse{}, nil
}

func (s *Server) Login(ctx context.Context, request *api2.LoginRequest) (*api2.LoginResponse, error) {
	username := request.Username
	password := request.Password
	token, err := s.authService.Login(username, password)
	if err != nil {
		return &api2.LoginResponse{}, err
	}
	return &api2.LoginResponse{Token: token}, nil
}

func (s *Server) Logout(ctx context.Context, request *api2.LogoutRequest) (*api2.LogoutResponse, error) {
	token := ctx.Value("token").(string)
	err := s.authService.Logout(token)
	if err != nil {
		return &api2.LogoutResponse{}, err
	}
	return &api2.LogoutResponse{}, nil
}

func (s *Server) SendMessage(ctx context.Context, request *api2.SendMessageRequest) (*api2.SendMessageResponse, error) {
	user := getUserFromContext(ctx)
	message := request.GetMessage()
	to := request.GetTo()

	msg := &model2.Message{
		Content: message.Content,
		Metadata: model2.MessageMetadata{
			Id: uuid.New(),
			To: model2.User{
				Username: to.Username,
			},
			From: model2.User{
				Username: user.Username,
			},
			Timestamp: time.Now(),
		},
	}

	if err := s.messageService.CreateMessage(msg); err != nil {
		return nil, err
	}

	return &api2.SendMessageResponse{}, nil
}

func (s *Server) ListMessages(ctx context.Context, request *api2.ListMessagesRequest) (*api2.ListMessagesResponse, error) {
	user := getUserFromContext(ctx)

	messages, err := s.messageService.GetMessages(user.Username)
	if err != nil {
		return nil, err
	}

	// Pagination parameters
	pageSize := int(request.GetPageSize())
	if pageSize <= 0 {
		pageSize = 50
	}
	if pageSize > 1000 {
		pageSize = 1000
	}

	// page_token is a simple integer offset encoded as a string
	start := 0
	if tok := request.GetPageToken(); tok != "" {
		if off, err := strconv.Atoi(tok); err == nil && off >= 0 {
			start = off
		}
	}
	if start > len(messages) {
		start = len(messages)
	}
	end := start + pageSize
	if end > len(messages) {
		end = len(messages)
	}

	page := messages[start:end]
	apiMessages := translateMessages(page)

	nextToken := ""
	if end < len(messages) {
		nextToken = strconv.Itoa(end)
	}

	response := &api2.ListMessagesResponse{
		Messages:      apiMessages,
		NextPageToken: nextToken,
	}
	return response, nil
}

func translateMessages(messages []*model2.Message) []*api2.Message {
	var apiMessages []*api2.Message
	for _, message := range messages {
		apiMessages = append(apiMessages, message.ToProto())
	}
	return apiMessages
}
