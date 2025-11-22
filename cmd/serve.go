package cmd

import (
	"fmt"
	"net"

	repository2 "github.com/djcopley/zing/internal/repository"
	"github.com/djcopley/zing/internal/server"
	service2 "github.com/djcopley/zing/internal/service"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc/reflection"
)

var (
	serverAddr = "localhost"
	serverPort = 5132
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the zing server",
	RunE:  runServe,
}

func runServe(cmd *cobra.Command, args []string) error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	defer func() { _ = logger.Sync() }()

	addr := fmt.Sprintf("%s:%d", serverAddr, serverPort)
	logger.Info("starting zing server", zap.String("addr", addr))

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Error("failed to listen", zap.String("addr", addr), zap.Error(err))
		return fmt.Errorf("failed to listen: %v", err)
	}

	userRepo := repository2.NewTestInMemoryUserRepository()
	sessionRepo := repository2.NewInMemorySessionRepository()
	messageRepo := repository2.NewInMemoryMessageRepository()

	authService := service2.NewAuthenticationService(userRepo, sessionRepo)
	messageService := service2.NewMessageService(messageRepo)

	s := server.NewServer(logger, authService, messageService)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		logger.Error("failed to serve", zap.Error(err))
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&serverAddr, "addr", "a", serverAddr, "Server address to bind to")
	serveCmd.Flags().IntVarP(&serverPort, "port", "p", serverPort, "Server port to bind to")
}
