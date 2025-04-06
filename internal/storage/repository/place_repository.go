package repository

import (
	"placelists-back/internal/storage"
	"placelists-back/internal/storage/entity"
	"placelists-back/pkg/adapter"
)

type placeRepositoryImpl struct {
	db adapter.Database
}

func NewPlaceRepository(db adapter.Database) storage.PlaceRepository {
	return &placeRepositoryImpl{db}
}

func (r *placeRepositoryImpl) GetByPublicID(id string) (entity.Place, error) {
	var p entity.Place
	result := r.db.First(p, "public_id = ?", id)
	return p, result.Error
}

func (r *placeRepositoryImpl) GetByPublicIDFull(id string) (entity.Place, error) {
	var p entity.Place
	result := r.db.Preload("Visitors").First(p, "public_id = ?", id)
	return p, result.Error
}

func (r *placeRepositoryImpl) GetByNameOrAddress(query string) ([]entity.Place, error) {
	var p []entity.Place
	result := r.db.Find(p, "lower(name) LIKE lower(?) OR lower(address) LIKE lower(?)", "%"+query+"%", "%"+query+"%")
	return p, result.Error
}

func (r *placeRepositoryImpl) GetByNameOrAddressFull(query string) ([]entity.Place, error) {
	var p []entity.Place
	result := r.db.Preload("Visitors").Find(p, "lower(name) LIKE lower(?) OR lower(address) LIKE lower(?)", "%"+query+"%", "%"+query+"%")
	return p, result.Error
}

func (r *placeRepositoryImpl) Create(p entity.Place) error {
	createErr := r.db.Create(p).Error
	return createErr
}

func (r *placeRepositoryImpl) Update(p entity.Place) error {
	result := r.db.Save(p)
	return result.Error
}
