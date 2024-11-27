package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
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

	VerifyToken(token string) (*models.User, error)
	VerifyTokenAdmin(token string) (*models.User, error)
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

func (s *authService) Login(ctx context.Context, email, password string) (string, bool, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", false, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", false, ErrInvalidCredentials
	}

	isAdmin := user.Role == "admin"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", false, err
	}

	return tokenString, isAdmin, nil
}

func (a *authService) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	user, err := a.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *authService) UpdateUser(ctx context.Context, userID uuid.UUID, name string, password string) error {
	panic("unimplemented")
}

func (a *authService) VerifyToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(a.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	userID, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return nil, ErrInvalidToken
	}

	user, err := a.userRepo.GetByID(context.Background(), userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

var (
	ErrUnauthorized = errors.New("user is not authorized as admin")
)

func (s *authService) VerifyTokenAdmin(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	userID, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return nil, ErrInvalidToken
	}

	user, err := s.userRepo.GetByID(context.Background(), userID)
	if err != nil {
		return nil, err
	}

	if user.Role != "admin" {
		return nil, ErrUnauthorized
	}

	return user, nil
}

// hellper functions
func WithUserContext(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, UserContextKey, user)
}

func UserFromContext(ctx context.Context) (*models.User, bool) {
	user, ok := ctx.Value(UserContextKey).(*models.User)
	return user, ok
}
