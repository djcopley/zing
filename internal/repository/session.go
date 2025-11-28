package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/djcopley/zing/internal/model"
	"github.com/redis/go-redis/v9"
)

type RedisSessionRepository struct {
	r *redis.Client
}

func (r RedisSessionRepository) Create(ctx context.Context, token string, user *model.User) error {
	if user == nil {
		return fmt.Errorf("user is nil")
	}
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := r.r.Set(ctx, sessionKey(token), b, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (r RedisSessionRepository) Read(ctx context.Context, token string) (*model.User, error) {
	res, err := r.r.Get(ctx, sessionKey(token)).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, err
	}
	var user model.User
	if err := json.Unmarshal(res, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r RedisSessionRepository) Delete(ctx context.Context, token string) error {
	if err := r.r.Del(ctx, sessionKey(token)).Err(); err != nil {
		return err
	}
	return nil
}

func NewRedisSessionRepository(r *redis.Client) *RedisSessionRepository {
	return &RedisSessionRepository{
		r: r,
	}
}

type InMemorySessionRepository struct {
	// username to session token
	sessions map[string]*model.User
}

func NewInMemorySessionRepository() *InMemorySessionRepository {
	return &InMemorySessionRepository{
		sessions: make(map[string]*model.User),
	}
}

func (r *InMemorySessionRepository) Create(ctx context.Context, token string, user *model.User) error {
	if user == nil {
		return fmt.Errorf("user is nil")
	}
	r.sessions[token] = user
	return nil
}

func (r *InMemorySessionRepository) Read(ctx context.Context, token string) (*model.User, error) {
	if user, ok := r.sessions[token]; ok {
		return user, nil
	}
	return nil, fmt.Errorf("user not found")
}

func (r *InMemorySessionRepository) Delete(ctx context.Context, token string) error {
	if _, ok := r.sessions[token]; ok {
		delete(r.sessions, token)
	}
	return nil
}

func sessionKey(token string) string { return "session:" + token }
