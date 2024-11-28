package models

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	GuestID      uuid.UUID `gorm:"type:uuid;not null" json:"guest_id"`
	RoomID       uuid.UUID `gorm:"type:uuid;not null" json:"room_id"`
	CheckInDate  time.Time `gorm:"type:date;not null" json:"check_in_date"`
	CheckOutDate time.Time `gorm:"type:date;not null" json:"check_out_date"`
	TotalAmount  float64   `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	Status       string    `gorm:"not null;default:'available'" json:"status"`

	Guest Guest `gorm:"foreignKey:GuestID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // FK para hóspede
	Room  Room  `gorm:"foreignKey:RoomID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`  // FK para quarto

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
