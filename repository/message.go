package repository

import (
	"errors"
	"github.com/djcopley/zing/api"
)

type MessageRepositoryInterface interface {
	Read(userId string) ([]*api.Message, error)
}

var _ MessageRepositoryInterface = &InMemoryMessageRepository{}

func NewInMemoryMessageRepository() *InMemoryMessageRepository {
	return &InMemoryMessageRepository{
		messages: make(map[string][]*api.Message),
	}
}

type InMemoryMessageRepository struct {
	messages map[string][]*api.Message
}

func (m *InMemoryMessageRepository) Read(userId string) ([]*api.Message, error) {
	msgs, ok := m.messages[userId]
	if !ok {
		return nil, errors.New("user not found")
	}
	return msgs, nil
}
