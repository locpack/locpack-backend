package entity

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

	Visitors []User `gorm:"many2many:user_visited_places"`
	Packs    []Pack `gorm:"many2many:place_packs"`
}

type Pack struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	PublicID string `gorm:"unique;not null"`
	Name     string `gorm:"not null"`

	AuthorID uuid.UUID `gorm:"type:uuid;not null"`
	Author   User      `gorm:"foreignKey:AuthorID"`

	Places        []Place `gorm:"many2many:pack_places"`
	FollowedUsers []User  `gorm:"many2many:pack_followed_users"`
}

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	PublicID string `gorm:"unique;not null"`
	Username string `gorm:"unique;not null"`

	FollowedPacks []Pack  `gorm:"many2many:user_followed_packs"`
	CreatedPacks  []Pack  `gorm:"foreignKey:AuthorID"`
	VisitedPlaces []Place `gorm:"many2many:user_visited_places"`
	CreatedPlaces []Place `gorm:"foreignKey:AuthorID"`
}
