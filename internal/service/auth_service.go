package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ruanv123/acme-hotel-api/internal/models"
	"github.com/ruanv123/acme-hotel-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type contextKey string

const (
	UserContextKey contextKey = "user"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
)

type AuthService interface {
	Register(ctx context.Context, email, password, name string) (*models.User, error)
	Login(ctx context.Context, email, password string) (token string, isAdmin bool, err error)
	UpdateUser(ctx context.Context, userID uuid.UUID, name, password string) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
}

type authService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewAuthService(
	userRepo repository.UserRepository,
	jwtSecret string,
) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *authService) Register(ctx context.Context, email string, password string, name string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *authService) Login(ctx context.Context, email string, password string) (token string, isAdmin bool, err error) {
	panic("unimplemented")
}

func (a *authService) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	panic("unimplemented")
}

func (a *authService) UpdateUser(ctx context.Context, userID uuid.UUID, name string, password string) error {
	panic("unimplemented")
}
