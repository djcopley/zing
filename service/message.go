package service

import (
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/repository"
	"log"
)

type MessageService struct {
	messageRepo repository.MessageRepositoryInterface
}

func (m *MessageService) GetMessages(username string) <-chan *api.Message {
	ch := make(chan *api.Message)
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
