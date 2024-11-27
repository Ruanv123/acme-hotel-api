package models

import "github.com/google/uuid"

type Reservation struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
}
