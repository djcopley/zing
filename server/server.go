package server

import (
	"context"
	"github.com/djcopley/zing/api"
	"google.golang.org/grpc"
)

var _ api.ZingServer = &Server{}

type Server struct {
	api.UnimplementedZingServer
}

func (s *Server) Login(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Logout(ctx context.Context, request *api.LogoutRequest) (*api.LogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetMessages(request *api.GetMessagesRequest, g grpc.ServerStreamingServer[api.GetMessagesResponse]) error {
	//TODO implement me
	panic("implement me")
}
