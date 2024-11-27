package repository

import (
	"context"

	"github.com/ruanv123/acme-hotel-api/internal/errors"
	"github.com/ruanv123/acme-hotel-api/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Create(ctx context.Context, user *models.User) error {
	result := u.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to create user")
	}
	return nil
}
