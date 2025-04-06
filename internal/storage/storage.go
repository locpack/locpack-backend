package storage

import (
	"placelists-back/internal/storage/entity"
)

type PlaceRepository interface {
	GetByPublicID(placeID string) (entity.Place, error)
	GetByPublicIDFull(placeID string) (entity.Place, error)
	GetByNameOrAddress(query string) ([]entity.Place, error)
	GetByNameOrAddressFull(query string) ([]entity.Place, error)
	Create(p entity.Place) error
	Update(p entity.Place) error
}

type PlacelistRepository interface {
	GetByPublicIDFull(id string) (entity.Placelist, error)
	GetByNameOrAuthorFull(query string) ([]entity.Placelist, error)
	Create(p entity.Placelist) error
	Update(p entity.Placelist) error
}

type UserRepository interface {
	GetByPublicID(id string) (entity.User, error)
	GetByPublicIDFull(id string) (entity.User, error)
	Create(u entity.User) error
	Update(u entity.User) error
}
