package service

import (
	"fmt"
	"github.com/djcopley/zing/model"
	"github.com/djcopley/zing/repository"
)

type AuthenticationService struct {
	userRepo    repository.UserRepositoryInterface
	sessionRepo repository.SessionRepositoryInterface
}

func (as *AuthenticationService) Login(username string, password string) (string, error) {
	user, err := as.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	if password != user.Password {
		return "", fmt.Errorf("invalid username or password")
	}
	token := generateSessionToken(user)
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

func generateSessionToken(user *model.User) string {
	return fmt.Sprintf("%s-%s", user.Username, user.Password)
}
