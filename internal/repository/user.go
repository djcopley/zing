package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/djcopley/zing/internal/model"
	"github.com/redis/go-redis/v9"
)

type RedisUserRepository struct {
	r *redis.Client
}

func (r RedisUserRepository) CreateUser(ctx context.Context, username, password string) error {
	if strings.TrimSpace(username) == "" {
		return fmt.Errorf("username is required")
	}
	key := userKey(username)
	// Use HSetNX to ensure we don't overwrite an existing user
	created, err := r.r.HSetNX(ctx, key, "password", password).Result()
	if err != nil {
		return err
	}
	if !created {
		return fmt.Errorf("user already exists")
	}
	return nil
}

func (r RedisUserRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	key := userKey(username)
	pwd, err := r.r.HGet(ctx, key, "password").Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, err
	}
	return &model.User{Username: username, Password: pwd}, nil
}

func NewRedisUserRepository(r *redis.Client) *RedisUserRepository {
	return &RedisUserRepository{
		r: r,
	}
}

func userKey(username string) string { return "user:" + username }
