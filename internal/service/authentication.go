package service

import (
    "context"
    "crypto/rand"
    "encoding/hex"
    "fmt"

    "github.com/djcopley/zing/internal/model"
    "golang.org/x/crypto/bcrypt"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, username, password string) error
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
}

type SessionRepositoryInterface interface {
	Create(ctx context.Context, token string, user *model.User) error
	Read(ctx context.Context, token string) (*model.User, error)
	Delete(ctx context.Context, token string) error
}

func NewAuthenticationService(userRepo UserRepositoryInterface, sessionRepo SessionRepositoryInterface) *AuthenticationService {
	return &AuthenticationService{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

type AuthenticationService struct {
    userRepo    UserRepositoryInterface
    sessionRepo SessionRepositoryInterface
}

func (as *AuthenticationService) Login(ctx context.Context, username string, password string) (string, error) {
    user, err := as.userRepo.GetUserByUsername(ctx, username)
    if err != nil {
        return "", err
    }
    // Compare provided password with stored bcrypt hash
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", fmt.Errorf("invalid username or password")
    }
    token, err := generateSessionToken()
    if err != nil {
        return "", err
    }
    err = as.sessionRepo.Create(ctx, token, user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (as *AuthenticationService) Logout(ctx context.Context, token string) error {
	err := as.sessionRepo.Delete(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

func (as *AuthenticationService) ValidateToken(ctx context.Context, token string) (*model.User, error) {
    user, err := as.sessionRepo.Read(ctx, token)
    if err != nil {
        return nil, err
    }
    return user, nil
}

// Register creates a new user with a bcrypt-hashed password and returns a session token.
func (as *AuthenticationService) Register(ctx context.Context, username string, password string) (string, error) {
    if username == "" || password == "" {
        return "", fmt.Errorf("username and password are required")
    }
    // Hash the password with bcrypt default cost
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", fmt.Errorf("failed to hash password: %w", err)
    }
    if err := as.userRepo.CreateUser(ctx, username, string(hashed)); err != nil {
        return "", err
    }
    // Auto-login: create session token
    token, err := generateSessionToken()
    if err != nil {
        return "", err
    }
    user := &model.User{Username: username, Password: string(hashed)}
    if err := as.sessionRepo.Create(ctx, token, user); err != nil {
        return "", err
    }
    return token, nil
}

func generateSessionToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
