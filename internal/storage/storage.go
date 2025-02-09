package storage

import (
	"placelists/internal/entities"

	"github.com/google/uuid"
)

type Repository struct {
	Place     PlaceRepository
	Placelist PlacelistRepository
	User      UserRepository
}

type PlaceRepository interface {
	GetByID(id uuid.UUID) (*entities.Place, error)
	GetByNameOrAddress(query string) (*[]entities.Place, error)
	Create(p *entities.Place) error
	Update(p *entities.Place) error
}

type PlacelistRepository interface {
	GetByID(id uuid.UUID) (*entities.Placelist, error)
	// GetByNameOrAuthor(query string) (*[]entities.Placelist, error)
	// GetFollowedByUsername(username string) (*[]entities.Placelist, error)
	// GetCreatedByUsername(username string) (*[]entities.Placelist, error)
	Create(p *entities.Placelist) error
	Update(p *entities.Placelist) error
	Delete(p *entities.Placelist) error
	// GetPlacelistPlacesByID
	// UpdatePlacelistPlacesByID
}

type UserRepository interface {
	GetByUsername(username string) (*entities.User, error)
	Create(p *entities.User) error
	Update(p *entities.User) error
}
