package repositories

import (
	"placelists-back/internal/storage"
	"placelists-back/pkg/database"
)

func NewRepository(db *database.DB) *storage.Repository {
	return &storage.Repository{
		Place:     NewPlaceRepository(db),
		Placelist: NewPlacelistRepository(db),
		User:      NewUserRepository(db),
	}
}
