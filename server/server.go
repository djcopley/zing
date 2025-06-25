package server

import (
	"context"
	"fmt"
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

var _ api.ZingServer = &Server{}

type Server struct {
	authService *service.AuthenticationService
	api.UnimplementedZingServer
}

func (s *Server) Login(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	username := request.Username
	password := request.Password
	if err := s.authService.Authenticate(username, password); err != nil {
		return &api.LoginResponse{}, fmt.Errorf("incorrect username or password")
	}
	return &api.LoginResponse{Id: "", Token: ""}, nil
}

func (s *Server) Logout(ctx context.Context, request *api.LogoutRequest) (*api.LogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetMessages(request *api.GetMessagesRequest, g grpc.ServerStreamingServer[api.GetMessagesResponse]) error {
	log.Printf("received request %+v\n", request)
	userId := request.Id
	_ = userId
	err := g.Send(&api.GetMessagesResponse{
		Sender: &api.User{
			Id:   "1",
			Name: "pop-o",
		},
		Content: &api.Message{
			Id:      "2",
			Time:    timestamppb.Now(),
			Content: "this is a message",
		},
	})
	return err
}
