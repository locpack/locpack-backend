package entities

import (
	"time"

	"github.com/google/uuid"
)

type PublicID = string

type PlacelistStatus = string

const (
	Created  PlacelistStatus = "CREATED"
	Followed PlacelistStatus = "FOLLOWED"
)

type Entity struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:(generate_random_uuid())"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time `gorm:"index"`
}

type PublicEntity struct {
	Entity
	PublicID PublicID `gorm:"unique;not null"`
}
