package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ReservationID uuid.UUID `gorm:"not null" json:"reservation_id"`
	AmountPaid    float64   `gorm:"type:decimal(10,2);not null" json:"amount_paid"`
	PaymentDate   time.Time `gorm:"type:date;not null" json:"payment_date"`
	PaymentMethod string    `gorm:"not null" json:"payment_method"`
	PaymentStatus string    `gorm:"default:'Pendente';not null" json:"payment_status"`

	Reservation Reservation `gorm:"foreignKey:ReservationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // FK

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
