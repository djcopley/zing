package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/djcopley/zing/internal/model"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, username, password string) error
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
}

type SessionRepositoryInterface interface {
	Create(ctx context.Context, token string, user *model.User) error
	Read(ctx context.Context, token string) (*model.User, error)
	Delete(ctx context.Context, token string) error
}

func NewAuthenticationService(userRepo UserRepositoryInterface, sessionRepo SessionRepositoryInterface) *AuthenticationService {
	return &AuthenticationService{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

type AuthenticationService struct {
	userRepo    UserRepositoryInterface
	sessionRepo SessionRepositoryInterface
}

func (as *AuthenticationService) Login(ctx context.Context, username string, password string) (string, error) {
	user, err := as.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	if password != user.Password {
		return "", fmt.Errorf("invalid username or password")
	}
	token, err := generateSessionToken()
	if err != nil {
		return "", err
	}
	err = as.sessionRepo.Create(ctx, token, user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (as *AuthenticationService) Logout(ctx context.Context, token string) error {
	err := as.sessionRepo.Delete(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

func (as *AuthenticationService) ValidateToken(ctx context.Context, token string) (*model.User, error) {
	user, err := as.sessionRepo.Read(ctx, token)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func generateSessionToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
