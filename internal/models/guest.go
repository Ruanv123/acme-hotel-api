package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Guest struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Cpf         string    `gorm:"type:varchar(14);uniqueIndex;not null" json:"cpf"`
	DataNasc    time.Time `gorm:"not null" json:"data_nasc"`
	Telefone    string    `gorm:"type:varchar(15);not null" json:"telefone"`
	Email       string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Observacoes string    `gorm:"type:text;not null" json:"observacoes"`
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (g *Guest) BeforeCreate(tx *gorm.DB) error {
	if g.ID == uuid.Nil {
		g.ID = uuid.New()
	}

	now := time.Now()
	if g.CreatedAt.IsZero() {
		g.CreatedAt = now
	}
	if g.UpdatedAt.IsZero() {
		g.UpdatedAt = now
	}

	return nil
}

func (g *Guest) BeforeUpdate(tx *gorm.DB) error {
	g.UpdatedAt = time.Now()
	return nil
}

func (Guest) TableName() string {
	return "guests"
}
