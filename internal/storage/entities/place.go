package entities

import (
	"time"

	"github.com/google/uuid"
)

type Place struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	PublicID string `gorm:"unique;not null"`
	Name     string `gorm:"not null"`
	Address  string `gorm:"not null"`

	AuthorID uuid.UUID `gorm:"type:uuid;not null"`
	Author   User      `gorm:"foreignKey:AuthorID"`

	Visitors   []User      `gorm:"many2many:user_visited_places"`
	Placelists []Placelist `gorm:"many2many:place_placelists"`
}
