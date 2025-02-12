package storage

import (
	"placelists/internal/storage/entities"
)

type Repository struct {
	Place     PlaceRepository
	Placelist PlacelistRepository
	User      UserRepository
}

type PlaceRepository interface {
	GetByPublicID(placeID string) (*entities.Place, error)
	GetByPublicIDFull(placeID string) (*entities.Place, error)
	GetByNameOrAddress(query string) (*[]entities.Place, error)
	GetByNameOrAddressFull(query string) (*[]entities.Place, error)
	Create(p *entities.Place) error
	Update(p *entities.Place) error
}

type PlacelistRepository interface {
	GetByPublicIDFull(id string) (*entities.Placelist, error)
	GetByNameOrAuthorFull(query string) (*[]entities.Placelist, error)
	Create(p *entities.Placelist) error
	Update(p *entities.Placelist) error
}

type UserRepository interface {
	GetByPublicID(id string) (*entities.User, error)
	GetByPublicIDFull(id string) (*entities.User, error)
	Create(u *entities.User) error
	Update(u *entities.User) error
}
