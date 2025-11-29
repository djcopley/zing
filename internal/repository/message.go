package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/djcopley/zing/internal/model"
	"github.com/redis/go-redis/v9"
)

type RedisMessageRepository struct {
	r *redis.Client
}

func (r *RedisMessageRepository) Create(ctx context.Context, message *model.Message) error {
	if message == nil {
		return nil
	}
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}
	key := messageKey(message.Metadata.To.Username)
	return r.r.RPush(ctx, key, b).Err()
}

func (r *RedisMessageRepository) Read(ctx context.Context, userId string) ([]*model.Message, error) {
	key := messageKey(userId)
	// Lua script to atomically get all list items and delete the key
	// Returns an array of strings (bulk replies)
	lua := `
local k = KEYS[1]
if redis.call("EXISTS", k) == 0 then
    return {}
end
local vals = redis.call("LRANGE", k, 0, -1)
redis.call("DEL", k)
return vals
`
	res, err := r.r.Eval(ctx, lua, []string{key}).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	var msgs []*model.Message
	if arr, ok := res.([]interface{}); ok {
		for _, v := range arr {
			if vb, ok := v.(string); ok {
				var m model.Message
				if err := json.Unmarshal([]byte(vb), &m); err == nil {
					msgs = append(msgs, &m)
				}
			} else if vb2, ok := v.([]byte); ok {
				var m model.Message
				if err := json.Unmarshal(vb2, &m); err == nil {
					msgs = append(msgs, &m)
				}
			}
		}
	}
	if msgs == nil {
		msgs = []*model.Message{}
	}
	return msgs, nil
}

func (r *RedisMessageRepository) Clear(ctx context.Context, userId string) error {
	key := messageKey(userId)
	return r.r.Del(ctx, key).Err()
}

func NewRedisMessageRepository(r *redis.Client) *RedisMessageRepository {
	return &RedisMessageRepository{
		r: r,
	}
}

func NewInMemoryMessageRepository() *InMemoryMessageRepository {
	return &InMemoryMessageRepository{
		messages: make(map[string][]*model.Message),
	}
}

type InMemoryMessageRepository struct {
	messages map[string][]*model.Message
}

func (m *InMemoryMessageRepository) Create(_ context.Context, message *model.Message) error {
	messages := m.messages[message.Metadata.To.Username]
	messages = append(messages, message)
	m.messages[message.Metadata.To.Username] = messages
	return nil
}

func (m *InMemoryMessageRepository) Read(_ context.Context, userId string) ([]*model.Message, error) {
	msgs, ok := m.messages[userId]
	if !ok {
		// If the user has no messages yet, return an empty slice instead of an error
		return []*model.Message{}, nil
	}
	return msgs, nil
}

func (m *InMemoryMessageRepository) Clear(_ context.Context, userId string) error {
	// Set to an empty slice to preserve key existence semantics
	m.messages[userId] = []*model.Message{}
	return nil
}

func messageKey(username string) string { return "messages:" + username }
