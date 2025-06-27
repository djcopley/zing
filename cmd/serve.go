package cmd

import (
	"fmt"
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/repository"
	"github.com/djcopley/zing/server"
	"github.com/djcopley/zing/service"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	rootCmd.AddCommand(serveCommand)
}

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "Serve the zing server",
	Run: func(cmd *cobra.Command, args []string) {
		addr := fmt.Sprintf("%s:%d", host, port)
		log.Printf("starting zing server @ %s\n", addr)

		lis, err := net.Listen("tcp", fmt.Sprintf("%s", addr))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		userRepo := repository.NewTestInMemoryUserRepository()
		sessionRepo := repository.NewInMemorySessionRepository()
		messageRepo := repository.NewInMemoryMessageRepository()

		authService := service.NewAuthenticationService(userRepo, sessionRepo)
		messageService := service.NewMessageService(messageRepo)

		s := grpc.NewServer()
		server := server.NewServer(authService, messageService)
		api.RegisterZingServer(s, server)

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}
