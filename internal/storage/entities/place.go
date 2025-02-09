package entities

import (
	"time"

	"github.com/google/uuid"
)

type Place struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time `gorm:"index;default:null"`
	PublicID  string    `gorm:"unique;not null"`

	Name    string `gorm:"not null"`
	Address string `gorm:"not null"`

	Placelists *[]PlacelistPlace
	Users      *[]UserPlace
}
