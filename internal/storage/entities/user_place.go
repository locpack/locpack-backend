package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserPlace struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time `gorm:"index;default:null"`

	Visited bool `gorm:"not null"`

	UserID  uuid.UUID
	PlaceID uuid.UUID

	User  *User  `gorm:"foreignkey:UserID;references:ID"`
	Place *Place `gorm:"foreignkey:PlaceID;references:ID"`
}
