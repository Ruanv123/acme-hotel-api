package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ruanv123/acme-hotel-api/internal/errors"
	"github.com/ruanv123/acme-hotel-api/internal/models"
	"gorm.io/gorm"
)

type GuestRepository interface {
	Create(ctx context.Context, guest *models.Guest) error
	ListAll(ctx context.Context) ([]models.Guest, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Guest, error)
	Update(ctx context.Context, guest *models.Guest) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type guestRepository struct {
	db *gorm.DB
}

func NewGuestRepository(db *gorm.DB) GuestRepository {
	return &guestRepository{db: db}
}

func (g *guestRepository) Create(ctx context.Context, guest *models.Guest) error {
	result := g.db.WithContext(ctx).Create(guest)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to create guest")
	}

	return nil
}

func (g *guestRepository) ListAll(ctx context.Context) ([]models.Guest, error) {
	var guests []models.Guest
	err := g.db.WithContext(ctx).Model(&models.Guest{}).Find(&guests).Error

	return guests, err
}

func (g *guestRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Guest, error) {
	var guest models.Guest
	result := g.db.WithContext(ctx).First(&guest, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, errors.Wrap(result.Error, "failed to get guest by ID")
	}

	return &guest, nil
}

func (g *guestRepository) Update(ctx context.Context, guest *models.Guest) error {
	result := g.db.WithContext(ctx).Model(guest).Updates(map[string]interface{}{
		"name":        guest.Name,
		"cpf":         guest.Cpf,
		"data_nasc":   guest.DataNasc,
		"telefone":    guest.Telefone,
		"email":       guest.Email,
		"observacoes": guest.Observacoes,
	})

	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to update guest")
	}

	if result.RowsAffected == 0 {
		return errors.ErrNotFound
	}

	return nil
}

func (g *guestRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := g.db.WithContext(ctx).Delete(&models.Guest{}, "id = ?", id)

	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to delete guest")
	}

	if result.RowsAffected == 0 {
		return errors.ErrNotFound
	}

	return nil
}
