package entities

import (
	"time"

	"github.com/google/uuid"
)

type Placelist struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	PublicID string `gorm:"unique;not null"`
	Name     string `gorm:"not null"`

	Places []PlacelistPlace
	Users  []UserPlacelist
}
