package service

import (
	"fmt"

	"github.com/djcopley/zing/internal/model"
	"github.com/djcopley/zing/internal/repository"
)

func NewMessageService(messageRepo repository.MessageRepositoryInterface) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
	}
}

type MessageService struct {
	messageRepo repository.MessageRepositoryInterface
}

func (m *MessageService) GetMessages(username string) ([]*model.Message, error) {
	messages, err := m.messageRepo.Read(username)
	if err != nil {
		return nil, fmt.Errorf("unable to read messages for user %s: %v", username, err)
	}
	return messages, nil
}

func (m *MessageService) CreateMessage(message *model.Message) error {
	return m.messageRepo.Create(message)
}

func (m *MessageService) ClearMessages(username string) error {
	if err := m.messageRepo.Clear(username); err != nil {
		return fmt.Errorf("unable to clear messages for user %s: %v", username, err)
	}
	return nil
}
