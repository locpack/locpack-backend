package storage

import (
	"placelists/internal/storage/entities"

	"github.com/google/uuid"
)

type Repository struct {
	Place     PlaceRepository
	Placelist PlacelistRepository
	User      UserRepository
}

type PlaceRepository interface {
	GetByPublicIDWithUser(publicID string, userID uuid.UUID) (*entities.Place, error)
	GetByNameOrAddressWithUser(query string, userID uuid.UUID) (*[]entities.Place, error)
	Create(p *entities.Place) error
	Update(p *entities.Place) error
}

type PlacelistRepository interface {
	GetByPublicID(publicID string) (*entities.Placelist, error)
	GetByNameOrAuthorWithUser(query string, userID uuid.UUID) (*[]entities.Placelist, error)
	// GetFollowedByUsername(username string) (*[]entities.Placelist, error)
	// GetCreatedByUsername(username string) (*[]entities.Placelist, error)
	Create(p *entities.Placelist) error
	Update(p *entities.Placelist) error
	// GetPlacelistPlacesByID
	// UpdatePlacelistPlacesByID
}

type UserRepository interface {
	GetByPublicID(publicID string) (*entities.User, error)
	Create(p *entities.User) error
	Update(p *entities.User) error
}
