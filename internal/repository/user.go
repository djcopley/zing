package repository

import (
    "context"
    "errors"
    "fmt"
    "strings"

    "github.com/djcopley/zing/internal/model"
    "github.com/redis/go-redis/v9"
    "golang.org/x/crypto/bcrypt"
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

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*model.User),
	}
}

func NewTestInMemoryUserRepository() *InMemoryUserRepository {
    imur := &InMemoryUserRepository{
        users: make(map[string]*model.User),
    }
    // Seed with users whose passwords are bcrypt-hashed versions of "pass"
    hash1, _ := bcrypt.GenerateFromPassword([]byte("pass"), bcrypt.DefaultCost)
    hash2, _ := bcrypt.GenerateFromPassword([]byte("pass"), bcrypt.DefaultCost)
    imur.users["user1"] = &model.User{Username: "user1", Password: string(hash1)}
    imur.users["user2"] = &model.User{Username: "user2", Password: string(hash2)}
    return imur
}

type InMemoryUserRepository struct {
	users map[string]*model.User
}

func (r *InMemoryUserRepository) CreateUser(ctx context.Context, username, password string) error {
	user := &model.User{
		Username: username,
		Password: password,
	}
	r.users[username] = user
	return nil
}

func (r *InMemoryUserRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	user, ok := r.users[username]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func userKey(username string) string { return "user:" + username }
