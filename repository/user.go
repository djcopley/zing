package repository

import (
	"fmt"
	"github.com/djcopley/zing/model"
)

type UserRepositoryInterface interface {
	CreateUser(username, password string) error
	GetUserByUsername(username string) (*model.User, error)
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*model.User),
	}
}

func NewTestInMemoryUserRepository() *InMemoryUserRepository {
	imur := &InMemoryUserRepository{
		users: make(map[string]*model.User),
	}
	imur.users["user1"] = &model.User{
		Username: "user1",
		Password: "pass",
	}
	imur.users["user2"] = &model.User{
		Username: "user2",
		Password: "pass",
	}
	return imur
}

var _ UserRepositoryInterface = &InMemoryUserRepository{}

type InMemoryUserRepository struct {
	users map[string]*model.User
}

func (r *InMemoryUserRepository) CreateUser(username, password string) error {
	user := &model.User{
		Username: username,
		Password: password,
	}
	r.users[username] = user
	return nil
}

func (r *InMemoryUserRepository) GetUserByUsername(userId string) (*model.User, error) {
	user, ok := r.users[userId]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}
