package service

import (
	"context"
	"fmt"

	"github.com/djcopley/zing/internal/model"
)

type MessageRepositoryInterface interface {
	Create(ctx context.Context, message *model.Message) error
	Read(ctx context.Context, userId string) ([]*model.Message, error)
	Clear(ctx context.Context, userId string) error
}

func NewMessageService(messageRepo MessageRepositoryInterface) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
	}
}

type MessageService struct {
	messageRepo MessageRepositoryInterface
}

func (m *MessageService) GetMessages(username string) ([]*model.Message, error) {
	messages, err := m.messageRepo.Read(context.TODO(), username)
	if err != nil {
		return nil, fmt.Errorf("unable to read messages for user %s: %v", username, err)
	}
	return messages, nil
}

func (m *MessageService) CreateMessage(message *model.Message) error {
	return m.messageRepo.Create(context.TODO(), message)
}

func (m *MessageService) ClearMessages(username string) error {
	if err := m.messageRepo.Clear(context.TODO(), username); err != nil {
		return fmt.Errorf("unable to clear messages for user %s: %v", username, err)
	}
	return nil
}
