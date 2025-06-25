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

func (as *AuthenticationService) Authenticate(username string, password string) error {
	user, err := as.userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}
	token := generateSessionToken(user)
	err = as.sessionRepo.Create(user.Username, token)
	if err != nil {
		return err
	}
	return nil
}

func generateSessionToken(user *model.User) string {
	return fmt.Sprintf("%s-%s", user.Username, user.Password)
}
