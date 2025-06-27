package repository

import (
	"errors"
	"github.com/djcopley/zing/model"
)

type MessageRepositoryInterface interface {
	Create(message *model.Message) error
	Read(userId string) ([]*model.Message, error)
}

var _ MessageRepositoryInterface = &InMemoryMessageRepository{}

func NewInMemoryMessageRepository() *InMemoryMessageRepository {
	return &InMemoryMessageRepository{
		messages: make(map[string][]*model.Message),
	}
}

type InMemoryMessageRepository struct {
	messages map[string][]*model.Message
}

func (m *InMemoryMessageRepository) Create(message *model.Message) error {
	messages := m.messages[message.Metadata.To.Username]
	messages = append(messages, message)
	m.messages[message.Metadata.To.Username] = messages
	return nil
}

func (m *InMemoryMessageRepository) Read(userId string) ([]*model.Message, error) {
	msgs, ok := m.messages[userId]
	if !ok {
		return nil, errors.New("user not found")
	}
	return msgs, nil
}
