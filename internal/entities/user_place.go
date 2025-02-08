package entities

import "github.com/google/uuid"

type UserPlace struct {
	Entity

	Visited bool `gorm:"not null"`

	UserID  uuid.UUID
	PlaceID uuid.UUID

	User  *User  `gorm:"foreignkey:UserID;references:ID"`
	Place *Place `gorm:"foreignkey:PlaceID;references:ID"`
}
