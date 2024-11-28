package models

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Number    int       `gorm:"not null" json:"number"`
	Type      string    `gorm:"type:varchar(50);not null" json:"type"`
	Capacity  int       `gorm:"not null" json:"capacity"`
	DailyRate float64   `gorm:"not null" json:"daily_rate"`
	Status    string    `gorm:"not null;default:'available'" json:"status"`

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
