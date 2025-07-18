package model

import (
	"github.com/djcopley/zing/api"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type MessageMetadata struct {
	Id        uuid.UUID `json:"id"`
	To        User      `json:"to"`
	From      User      `json:"from"`
	Timestamp time.Time `json:"timestamp"`
}

func (m *MessageMetadata) ToProto() *api.MessageMetadata {
	return &api.MessageMetadata{
		Id:        m.Id.String(),
		To:        &api.User{Username: m.To.Username},
		From:      &api.User{Username: m.From.Username},
		Timestamp: timestamppb.New(m.Timestamp),
	}
}

type Message struct {
	Metadata MessageMetadata `json:"metadata"`
	Content  string          `json:"content"`
}

func (m *Message) ToProto() *api.Message {
	return &api.Message{
		Metadata: m.Metadata.ToProto(),
		Content:  m.Content,
	}
}
