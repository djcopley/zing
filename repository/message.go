package repository

import (
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
		// If the user has no messages yet, return an empty slice instead of an error
		return []*model.Message{}, nil
	}
	return msgs, nil
}
