package entities

import "github.com/google/uuid"

type PlacelistPlace struct {
	Entity

	PlacelistID uuid.UUID
	PlaceID     uuid.UUID

	Placelist *Placelist `gorm:"foreignkey:PlacelistID;references:ID"`
	Place     *Place     `gorm:"foreignkey:PlaceID;references:ID"`
}
