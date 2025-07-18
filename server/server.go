package server

import (
	"context"
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/model"
	"github.com/djcopley/zing/service"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"time"
)

var _ api.ZingServer = &Server{}

func NewServer(authService *service.AuthenticationService, messageService *service.MessageService) *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(newAuthInterceptor(authService)),
	)
	zingServer := &Server{
		authService:    authService,
		messageService: messageService,
	}
	api.RegisterZingServer(grpcServer, zingServer)
	return grpcServer
}

type Server struct {
	authService    *service.AuthenticationService
	messageService *service.MessageService
	api.UnimplementedZingServer
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
	apiMessages := translateMessages(messages)
	response := &api.ListMessagesResponse{
		Messages:      apiMessages,
		NextPageToken: "", // todo
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
