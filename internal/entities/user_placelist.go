package entities

import "github.com/google/uuid"

type UserPlacelist struct {
	Entity

	Status PlacelistStatus `gorm:"not null"`

	UserID      uuid.UUID
	PlacelistID uuid.UUID

	User      *User      `gorm:"foreignkey:UserID;references:ID"`
	Placelist *Placelist `gorm:"foreignkey:PlacelistID;references:ID"`
}
