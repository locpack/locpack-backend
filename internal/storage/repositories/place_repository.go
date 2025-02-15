package repositories

import (
	"placelists/internal/storage/database"
	"placelists/internal/storage/entities"
)

type placeRepositoryImpl struct {
	db *database.DB
}

func NewPlaceRepository(db *database.DB) *placeRepositoryImpl {
	return &placeRepositoryImpl{db}
}

func (r *placeRepositoryImpl) GetByPublicID(id string) (entities.Place, error) {
	var p entities.Place
	result := r.db.First(p, "public_id = ?", id)
	return p, result.Error
}

func (r *placeRepositoryImpl) GetByPublicIDFull(id string) (entities.Place, error) {
	var p entities.Place
	result := r.db.Preload("Visitors").First(p, "public_id = ?", id)
	return p, result.Error
}

func (r *placeRepositoryImpl) GetByNameOrAddress(query string) ([]entities.Place, error) {
	var p []entities.Place
	result := r.db.Find(p, "lower(name) LIKE lower(?) OR lower(address) LIKE lower(?)", "%"+query+"%", "%"+query+"%")
	return p, result.Error
}

func (r *placeRepositoryImpl) GetByNameOrAddressFull(query string) ([]entities.Place, error) {
	var p []entities.Place
	result := r.db.Preload("Visitors").Find(p, "lower(name) LIKE lower(?) OR lower(address) LIKE lower(?)", "%"+query+"%", "%"+query+"%")
	return p, result.Error
}

func (r *placeRepositoryImpl) Create(p entities.Place) error {
	createErr := r.db.Create(p).Error
	return createErr
}

func (r *placeRepositoryImpl) Update(p entities.Place) error {
	result := r.db.Save(p)
	return result.Error
}
