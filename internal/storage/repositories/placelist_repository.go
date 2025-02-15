package repositories

import (
	"placelists/internal/storage/database"
	"placelists/internal/storage/entities"
)

type placelistRepositoryImpl struct {
	db *database.DB
}

func NewPlacelistRepository(db *database.DB) *placelistRepositoryImpl {
	return &placelistRepositoryImpl{db}
}

func (r *placelistRepositoryImpl) GetByPublicIDFull(id string) (entities.Placelist, error) {
	var p entities.Placelist
	result := r.db.Preload("FollowedUsers").First(p, "public_id = ?", id)
	return p, result.Error
}

func (r *placelistRepositoryImpl) GetByNameOrAuthorFull(query string) ([]entities.Placelist, error) {
	var p []entities.Placelist
	result := r.db.Preload("FollowedUsers").Find(p, "lower(name) LIKE lower(?) OR lower(author_id) LIKE lower(?)", "%"+query+"%", "%"+query+"%")
	return p, result.Error
}

func (r *placelistRepositoryImpl) Create(p entities.Placelist) error {
	createErr := r.db.Create(p).Error
	return createErr
}

func (r *placelistRepositoryImpl) Update(p entities.Placelist) error {
	result := r.db.Save(p)
	return result.Error
}
