package server

import (
	"context"
	"strings"

	"github.com/djcopley/zing/model"
	"github.com/djcopley/zing/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type userContextKey struct{}

func newAuthInterceptor(authService *service.AuthenticationService) grpc.UnaryServerInterceptor {
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
		token := strings.Split(bearer, " ")[1]

		user, err := authService.ValidateToken(token)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		ctx = context.WithValue(ctx, userContextKey{}, user)

		return handler(ctx, req)
	}
}

func shouldSkipAuth(fullMethod string) bool {
	switch fullMethod {
	case "/zing.Zing/Login",
		"/zing.Zing/Register":
		return true
	default:
		return false
	}
}

func getUserFromContext(ctx context.Context) *model.User {
	user := ctx.Value(userContextKey{}).(*model.User)
	return user
}
