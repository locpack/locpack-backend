package entities

import (
	"time"

	"github.com/google/uuid"
)

type PlacelistStatus = string

const (
	Created  PlacelistStatus = "CREATED"
	Followed PlacelistStatus = "FOLLOWED"
)

type UserPlacelist struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	Status PlacelistStatus `gorm:"not null"`

	UserID      uuid.UUID
	PlacelistID uuid.UUID

	User      User      `gorm:"foreignkey:UserID;references:ID"`
	Placelist Placelist `gorm:"foreignkey:PlacelistID;references:ID"`
}
