package models

import "github.com/google/uuid"

type Payment struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
}
