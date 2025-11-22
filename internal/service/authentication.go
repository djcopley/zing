package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/djcopley/zing/internal/model"
	repository2 "github.com/djcopley/zing/internal/repository"
)

func NewAuthenticationService(userRepo repository2.UserRepositoryInterface, sessionRepo repository2.SessionRepositoryInterface) *AuthenticationService {
	return &AuthenticationService{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

type AuthenticationService struct {
	userRepo    repository2.UserRepositoryInterface
	sessionRepo repository2.SessionRepositoryInterface
}

func (as *AuthenticationService) Login(username string, password string) (string, error) {
	user, err := as.userRepo.GetUserByUsername(username)
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
	err = as.sessionRepo.Create(token, user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (as *AuthenticationService) Logout(token string) error {
	err := as.sessionRepo.Delete(token)
	if err != nil {
		return err
	}
	return nil
}

func (as *AuthenticationService) ValidateToken(token string) (*model.User, error) {
	user, err := as.sessionRepo.Read(token)
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
