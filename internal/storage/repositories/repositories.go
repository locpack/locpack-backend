package repositories

import (
	"placelists/internal/storage"
	"placelists/pkg/database"
)

func NewRepository(db *database.DB) *storage.Repository {
	return &storage.Repository{
		Place:     NewPlaceRepository(db),
		Placelist: NewPlacelistRepository(db),
		User:      NewUserRepository(db),
	}
}
