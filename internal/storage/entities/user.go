package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	PublicID string `gorm:"unique;not null"`
	Username string `gorm:"unique;not null"`

	FollwedPlacelists []Placelist `gorm:"many2many:user_followed_placelists"`
	CreatedPlacelists []Placelist `gorm:"foreignKey:AuthorID"`
	VisitedPlaces     []Place     `gorm:"many2many:user_visited_places"`
	CreatedPlaces     []Place     `gorm:"foreignKey:AuthorID"`
}
