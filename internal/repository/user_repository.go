package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ruanv123/acme-hotel-api/internal/errors"
	"github.com/ruanv123/acme-hotel-api/internal/models"
	"gorm.io/gorm"
)

// interface assinando os metodos
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	// ====
	GrantAccess(ctx context.Context, id uuid.UUID) error
	RevokeAccess(ctx context.Context, id uuid.UUID) error
	// ====
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
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

func (u *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	result := u.db.WithContext(ctx).First(&user, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, errors.Wrap(result.Error, "failed to get user by ID")
	}

	return &user, nil
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	result := u.db.WithContext(ctx).First(&user, "email = ?", email)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, errors.Wrap(result.Error, "failed to get user by email")
	}

	return &user, nil
}

func (u *userRepository) Update(ctx context.Context, user *models.User) error {
	result := u.db.WithContext(ctx).Model(user).Updates(map[string]interface{}{
		"email":         user.Email,
		"name":          user.Name,
		"password_hash": user.PasswordHash,
		"updated_at":    user.UpdatedAt,
	})

	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to update user")
	}

	if result.RowsAffected == 0 {
		return errors.ErrNotFound
	}

	return nil
}

func (u *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := u.db.WithContext(ctx).Delete(&models.User{}, "id = ?", id)

	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to delete user")
	}

	if result.RowsAffected == 0 {
		return errors.ErrNotFound
	}

	return nil
}

func (u *userRepository) GrantAccess(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

func (u *userRepository) RevokeAccess(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
