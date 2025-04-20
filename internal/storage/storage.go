package storage

import (
	"locpack-backend/internal/storage/entity"
)

type PlaceRepository interface {
	GetByPublicID(placeID string) (entity.Place, error)
	GetByPublicIDFull(placeID string) (entity.Place, error)
	GetByNameOrAddress(query string) ([]entity.Place, error)
	GetByNameOrAddressFull(query string) ([]entity.Place, error)
	Create(p entity.Place) error
	Update(p entity.Place) error
}

type PackRepository interface {
	GetByPublicIDFull(id string) (entity.Pack, error)
	GetByNameOrAuthorFull(query string) ([]entity.Pack, error)
	Create(p entity.Pack) error
	Update(p entity.Pack) error
}

type UserRepository interface {
	GetByPublicID(id string) (entity.User, error)
	GetByPublicIDFull(id string) (entity.User, error)
	Create(u entity.User) error
	Update(u entity.User) error
}
