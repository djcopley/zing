package server

import (
	"context"
	"fmt"
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/service"
	"google.golang.org/grpc"
)

var _ api.ZingServer = &Server{}

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
		return &api.LoginResponse{}, fmt.Errorf("incorrect username or password")
	}
	return &api.LoginResponse{Token: token}, nil
}

func (s *Server) Logout(ctx context.Context, request *api.LogoutRequest) (*api.LogoutResponse, error) {
	token := request.GetToken()
	err := s.authService.Logout(token)
	if err != nil {
		return &api.LogoutResponse{}, fmt.Errorf("something went wrong")
	}
	return &api.LogoutResponse{}, nil
}

func (s *Server) SendMessage(ctx context.Context, request *api.SendMessageRequest) (*api.SendMessageResponse, error) {
	message := request.GetMessage()
	to := request.GetTo()
	return nil, nil
}

func (s *Server) GetMessages(request *api.GetMessagesRequest, g grpc.ServerStreamingServer[api.GetMessagesResponse]) error {
	token := request.GetToken()
	user, err := s.authService.ValidateToken(token)
	if err != nil {
		return err
	}
	messageChan := s.messageService.GetMessages(user.Username)
	for _ = range messageChan {
		msg := &api.GetMessagesResponse{
			Metadata: &api.MessageMetadata{
				Id:        "",
				To:        nil,
				From:      nil,
				Timestamp: nil,
			},
			Message: nil,
		}
		err = g.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
