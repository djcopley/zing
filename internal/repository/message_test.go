package repository

import (
	"testing"
	"time"

	model2 "github.com/djcopley/zing/internal/model"
	"github.com/google/uuid"
)

func TestCreateMessage(t *testing.T) {
	repo := NewInMemoryMessageRepository()
	message := &model2.Message{
		Content: "This is a test message",
		Metadata: model2.MessageMetadata{
			Id: uuid.New(),
			To: model2.User{
				Username: "test_to",
				Password: "test",
			},
			From: model2.User{
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
