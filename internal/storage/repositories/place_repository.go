package repositories

import (
	"placelists/internal/storage/database"
	"placelists/internal/storage/entities"

	"github.com/google/uuid"
)

type placeRepositoryImpl struct {
	db *database.DB
}

func NewPlaceRepository(db *database.DB) *placeRepositoryImpl {
	return &placeRepositoryImpl{db}
}

func (r *placeRepositoryImpl) GetByID(id uuid.UUID) (*entities.Place, error) {
	var p *entities.Place
	result := r.db.First(&p, "id = ? AND deleted_at IS NULL", id)
	return p, result.Error
}

func (r *placeRepositoryImpl) GetByNameOrAddress(query string) (*[]entities.Place, error) {
	var p *[]entities.Place
	result := r.db.Find(&p, "(name = ? OR address = ?) AND deleted_at IS NULL", query, query)
	return p, result.Error
}

func (r *placeRepositoryImpl) Create(p *entities.Place) error {
	result := r.db.Create(&p)
	return result.Error
}

func (r *placeRepositoryImpl) Update(p *entities.Place) error {
	result := r.db.Save(&p)
	return result.Error
}
