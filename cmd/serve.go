package cmd

import (
	"fmt"
	"github.com/djcopley/zing/repository"
	"github.com/djcopley/zing/server"
	"github.com/djcopley/zing/service"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/reflection"
	"net"
)

var (
	serverAddr = "localhost"
	serverPort = 5132
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the zing server",
	RunE: func(cmd *cobra.Command, args []string) error {
		addr := fmt.Sprintf("%s:%d", serverAddr, serverPort)
		_, err := fmt.Fprintf(cmd.OutOrStdout(), "starting zing server @ %s\n", addr)
		if err != nil {
			return err
		}

		lis, err := net.Listen("tcp", addr)
		if err != nil {
			return fmt.Errorf("failed to listen: %v", err)
		}

		userRepo := repository.NewTestInMemoryUserRepository()
		sessionRepo := repository.NewInMemorySessionRepository()
		messageRepo := repository.NewInMemoryMessageRepository()

		authService := service.NewAuthenticationService(userRepo, sessionRepo)
		messageService := service.NewMessageService(messageRepo)

		server := server.NewServer(authService, messageService)
		reflection.Register(server)

		if err := server.Serve(lis); err != nil {
			return fmt.Errorf("failed to serve: %v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&serverAddr, "addr", "a", serverAddr, "Server address to bind to")
	serveCmd.Flags().IntVarP(&serverPort, "port", "p", serverPort, "Server port to bind to")
}
