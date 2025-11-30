package server

import (
	"context"
	"strings"
	"time"

	"github.com/djcopley/zing/internal/model"
	"github.com/djcopley/zing/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type userContextKey struct{}

func shouldSkipAuth(fullMethod string) bool {
	switch fullMethod {
	case "/zing.Zing/Login",
		"/zing.Zing/Register",
		"/grpc.health.v1.Health/Check",
		"/grpc.health.v1.Health/List",
		"/grpc.health.v1.Health/Watch":
		return true
	default:
		return false
	}
}

func getUserFromContext(ctx context.Context) *model.User {
	user := ctx.Value(userContextKey{}).(*model.User)
	return user
}

func NewAuthInterceptor(authService *service.AuthenticationService) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Skip auth for public endpoints
		if shouldSkipAuth(info.FullMethod) {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		authHeaders := md.Get("authorization")
		if len(authHeaders) == 0 {
			return nil, status.Error(codes.Unauthenticated, "authorization header required")
		}

		bearer := authHeaders[0]
		parts := strings.Split(bearer, " ")
		if len(parts) < 2 {
			return nil, status.Error(codes.Unauthenticated, "authorization header format must be: Bearer <token>")
		}
		token := parts[1]

		user, err := authService.ValidateToken(context.TODO(), token)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "your session has expired. please log in again")
		}

		ctx = context.WithValue(ctx, userContextKey{}, user)

		return handler(ctx, req)
	}
}

func NewLoggingInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		dur := time.Since(start)

		p, _ := peer.FromContext(ctx)
		peerAddr := ""
		if p != nil {
			peerAddr = p.Addr.String()
		}

		if err != nil {
			st, _ := status.FromError(err)
			logger.Warn("grpc.request",
				zap.String("grpc.method", info.FullMethod),
				zap.Duration("grpc.duration", dur),
				zap.String("peer", peerAddr),
				zap.String("grpc.code", st.Code().String()),
				zap.Error(st.Err()),
			)
			return resp, err
		}

		logger.Info("grpc.request",
			zap.String("grpc.method", info.FullMethod),
			zap.Duration("grpc.duration", dur),
			zap.String("peer", peerAddr),
			zap.String("grpc.code", codes.OK.String()),
		)
		return resp, nil
	}
}
