package service

import (
	"fmt"

	"github.com/djcopley/zing/model"
	"github.com/djcopley/zing/repository"
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
