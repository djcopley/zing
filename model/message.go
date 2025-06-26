package model

import (
	"github.com/djcopley/zing/api"
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Content  string          `json:"content"`
	Metadata MessageMetadata `json:"metadata"`
}

type MessageMetadata struct {
	Id        uuid.UUID `json:"id"`
	To        User      `json:"to"`
	From      User      `json:"from"`
	Timestamp time.Time `json:"timestamp"`
}

func (m *Message) ToProto() *api.Message {
	return &api.Message{}
}
