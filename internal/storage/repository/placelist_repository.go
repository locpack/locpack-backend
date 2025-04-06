package repository

import (
	"placelists-back/internal/storage"
	"placelists-back/internal/storage/entity"
	"placelists-back/pkg/adapter"
)

type placelistRepositoryImpl struct {
	db adapter.Database
}

func NewPlacelistRepository(db adapter.Database) storage.PlacelistRepository {
	return &placelistRepositoryImpl{db}
}

func (r *placelistRepositoryImpl) GetByPublicIDFull(id string) (entity.Placelist, error) {
	var p entity.Placelist
	result := r.db.Preload("FollowedUsers").First(p, "public_id = ?", id)
	return p, result.Error
}

func (r *placelistRepositoryImpl) GetByNameOrAuthorFull(query string) ([]entity.Placelist, error) {
	var p []entity.Placelist
	result := r.db.Preload("FollowedUsers").Find(p, "lower(name) LIKE lower(?) OR lower(author_id) LIKE lower(?)", "%"+query+"%", "%"+query+"%")
	return p, result.Error
}

func (r *placelistRepositoryImpl) Create(p entity.Placelist) error {
	createErr := r.db.Create(p).Error
	return createErr
}

func (r *placelistRepositoryImpl) Update(p entity.Placelist) error {
	result := r.db.Save(p)
	return result.Error
}
