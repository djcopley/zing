package service

import (
	"github.com/djcopley/zing/model"
	"github.com/djcopley/zing/repository"
	"log"
)

func NewMessageService(messageRepo repository.MessageRepositoryInterface) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
	}
}

type MessageService struct {
	messageRepo repository.MessageRepositoryInterface
}

func (m *MessageService) GetMessages(username string) <-chan *model.Message {
	ch := make(chan *model.Message)
	go func() {
		defer close(ch)
		messages, err := m.messageRepo.Read(username)
		if err != nil {
			log.Printf("unable to read messages for user %s: %v", username, err)
		}
		for _, msg := range messages {
			ch <- msg
		}
	}()
	return ch
}

func (m *MessageService) CreateMessage(message *model.Message) error {
	return m.messageRepo.Create(message)
}
