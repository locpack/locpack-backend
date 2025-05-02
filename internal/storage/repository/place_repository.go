package repository

import (
	"locpack-backend/internal/storage"
	"locpack-backend/internal/storage/entity"
	"locpack-backend/pkg/adapter"
)

type placeRepoImpl struct {
	db adapter.Database
}

func NewPlaceRepository(db adapter.Database) storage.PlaceRepository {
	return &placeRepoImpl{db}
}

func (r *placeRepoImpl) GetByPublicID(id string) (entity.Place, error) {
	var p entity.Place
	result := r.db.First(&p, "public_id = ?", id)
	return p, result.Error
}

func (r *placeRepoImpl) GetByPublicIDFull(id string) (entity.Place, error) {
	var p entity.Place
	result := r.db.Preload("Visitors").First(&p, "public_id = ?", id)
	return p, result.Error
}

func (r *placeRepoImpl) GetByNameOrAddress(query string) ([]entity.Place, error) {
	var p []entity.Place
	result := r.db.Find(&p, "lower(name) LIKE lower(?) OR lower(address) LIKE lower(?)", "%"+query+"%", "%"+query+"%")
	return p, result.Error
}

func (r *placeRepoImpl) GetByNameOrAddressFull(query string) ([]entity.Place, error) {
	var p []entity.Place
	result := r.db.Preload("Visitors").Find(&p, "lower(name) LIKE lower(?) OR lower(address) LIKE lower(?)", "%"+query+"%", "%"+query+"%")
	return p, result.Error
}

func (r *placeRepoImpl) Create(p entity.Place) error {
	createErr := r.db.Create(&p).Error
	return createErr
}

func (r *placeRepoImpl) Update(p entity.Place) error {
	result := r.db.Save(&p)
	return result.Error
}
