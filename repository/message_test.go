package repository

import (
	"github.com/djcopley/zing/model"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestCreateMessage(t *testing.T) {
	repo := NewInMemoryMessageRepository()
	message := &model.Message{
		Content: "This is a test message",
		Metadata: model.MessageMetadata{
			Id: uuid.New(),
			To: model.User{
				Username: "test_to",
				Password: "test",
			},
			From: model.User{
				Username: "test_from",
				Password: "test",
			},
			Timestamp: time.Now(),
		},
	}
	err := repo.Create(message)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if repo.messages["test_to"][0] != message {
		t.Fatalf("unexpected message: %v", repo.messages["test_to"])
	}
}
