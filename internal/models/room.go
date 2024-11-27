package models

import "github.com/google/uuid"

type Room struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
}
