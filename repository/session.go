package repository

import (
	"fmt"

	"github.com/djcopley/zing/model"
)

type SessionRepositoryInterface interface {
	Create(token string, user *model.User) error
	Read(token string) (*model.User, error)
	Delete(token string) error
}

type InMemorySessionRepository struct {
	// username to session token
	sessions map[string]*model.User
}

var _ SessionRepositoryInterface = &InMemorySessionRepository{}

func NewInMemorySessionRepository() *InMemorySessionRepository {
	return &InMemorySessionRepository{
		sessions: make(map[string]*model.User),
	}
}

func (r *InMemorySessionRepository) Create(token string, user *model.User) error {
	if user == nil {
		return fmt.Errorf("user is nil")
	}
	r.sessions[token] = user
	return nil
}

func (r *InMemorySessionRepository) Read(token string) (*model.User, error) {
	if user, ok := r.sessions[token]; ok {
		return user, nil
	}
	return nil, fmt.Errorf("user not found")
}

func (r *InMemorySessionRepository) Delete(token string) error {
	if _, ok := r.sessions[token]; ok {
		delete(r.sessions, token)
	}
	return nil
}
