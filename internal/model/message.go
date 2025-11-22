package model

import (
	"time"

	api2 "github.com/djcopley/zing/internal/api"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MessageMetadata struct {
	Id        uuid.UUID `json:"id"`
	To        User      `json:"to"`
	From      User      `json:"from"`
	Timestamp time.Time `json:"timestamp"`
}

func (m *MessageMetadata) ToProto() *api2.MessageMetadata {
	return &api2.MessageMetadata{
		Id:        m.Id.String(),
		To:        &api2.User{Username: m.To.Username},
		From:      &api2.User{Username: m.From.Username},
		Timestamp: timestamppb.New(m.Timestamp),
	}
}

type Message struct {
	Metadata MessageMetadata `json:"metadata"`
	Content  string          `json:"content"`
}

func (m *Message) ToProto() *api2.Message {
	return &api2.Message{
		Metadata: m.Metadata.ToProto(),
		Content:  m.Content,
	}
}
