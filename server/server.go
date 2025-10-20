package server

import (
	"context"
	"strconv"
	"time"

	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/model"
	"github.com/djcopley/zing/service"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var _ api.ZingServer = &Server{}

func NewServer(logger *zap.Logger, authService *service.AuthenticationService, messageService *service.MessageService) *grpc.Server {
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
	api.RegisterZingServer(grpcServer, zingServer)
	return grpcServer
}

type Server struct {
	authService    *service.AuthenticationService
	messageService *service.MessageService
	logger         *zap.Logger
	api.UnimplementedZingServer
}

func (s *Server) ClearMessages(ctx context.Context, request *api.ClearMessagesRequest) (*api.ClearMessagesResponse, error) {
	user := getUserFromContext(ctx)
	if err := s.messageService.ClearMessages(user.Username); err != nil {
		return &api.ClearMessagesResponse{}, err
	}
	return &api.ClearMessagesResponse{}, nil
}

func (s *Server) Login(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	username := request.Username
	password := request.Password
	token, err := s.authService.Login(username, password)
	if err != nil {
		return &api.LoginResponse{}, err
	}
	return &api.LoginResponse{Token: token}, nil
}

func (s *Server) Logout(ctx context.Context, request *api.LogoutRequest) (*api.LogoutResponse, error) {
	token := ctx.Value("token").(string)
	err := s.authService.Logout(token)
	if err != nil {
		return &api.LogoutResponse{}, err
	}
	return &api.LogoutResponse{}, nil
}

func (s *Server) SendMessage(ctx context.Context, request *api.SendMessageRequest) (*api.SendMessageResponse, error) {
	user := getUserFromContext(ctx)
	message := request.GetMessage()
	to := request.GetTo()

	msg := &model.Message{
		Content: message.Content,
		Metadata: model.MessageMetadata{
			Id: uuid.New(),
			To: model.User{
				Username: to.Username,
			},
			From: model.User{
				Username: user.Username,
			},
			Timestamp: time.Now(),
		},
	}

	if err := s.messageService.CreateMessage(msg); err != nil {
		return nil, err
	}

	return &api.SendMessageResponse{}, nil
}

func (s *Server) ListMessages(ctx context.Context, request *api.ListMessagesRequest) (*api.ListMessagesResponse, error) {
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

	response := &api.ListMessagesResponse{
		Messages:      apiMessages,
		NextPageToken: nextToken,
	}
	return response, nil
}

func translateMessages(messages []*model.Message) []*api.Message {
	var apiMessages []*api.Message
	for _, message := range messages {
		apiMessages = append(apiMessages, message.ToProto())
	}
	return apiMessages
}
