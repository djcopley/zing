package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/djcopley/zing/internal/model"
	"github.com/redis/go-redis/v9"
)

const (
	// sessionTTL defines how long a session is valid for.
	sessionTTL = time.Hour * 24 * 7 // 1 week
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
	// Set the session with a TTL so it expires automatically in Redis
	if err := r.r.Set(ctx, sessionKey(token), b, sessionTTL).Err(); err != nil {
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

func sessionKey(token string) string { return "session:" + token }
