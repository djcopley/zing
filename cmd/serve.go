package cmd

import (
	"fmt"
	"net"

	"github.com/djcopley/zing/internal/api"
	"github.com/djcopley/zing/internal/repository"
	"github.com/djcopley/zing/internal/server"
	"github.com/djcopley/zing/internal/service"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

var (
	serverAddr = "localhost"
	serverPort = 50051
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

	r := redis.NewClient(&redis.Options{
		Addr:     config.GetRedisAddr(),
		Username: config.GetRedisUsername(),
		Password: config.GetRedisPassword(),
		DB:       config.GetRedisDB(),
	})

	userRepo := repository.NewRedisUserRepository(r)
	sessionRepo := repository.NewRedisSessionRepository(r)
	messageRepo := repository.NewRedisMessageRepository(r)

	authService := service.NewAuthenticationService(userRepo, sessionRepo)
	messageService := service.NewMessageService(messageRepo)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			server.NewLoggingInterceptor(logger),
			server.NewAuthInterceptor(authService),
		),
	)
	reflection.Register(grpcServer)

	zingService := server.NewServer(authService, messageService)
	api.RegisterZingServer(grpcServer, zingService)

	healthService := health.NewServer()
	healthService.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(grpcServer, healthService)

	if err := grpcServer.Serve(lis); err != nil {
		logger.Error("failed to serve", zap.Error(err))
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&serverAddr, "addr", "a", serverAddr, "Server address to bind to")
	serveCmd.Flags().IntVarP(&serverPort, "port", "", serverPort, "Server port to bind to")
}
