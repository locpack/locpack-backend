package entities

import (
	"time"

	"github.com/google/uuid"
)

type PlacelistPlace struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	PlacelistID uuid.UUID
	PlaceID     uuid.UUID

	Placelist Placelist `gorm:"foreignkey:PlacelistID;references:ID"`
	Place     Place     `gorm:"foreignkey:PlaceID;references:ID"`
}
