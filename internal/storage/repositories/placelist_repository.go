package repositories

import (
	"placelists/internal/storage/database"
	"placelists/internal/storage/entities"
	"time"

	"github.com/google/uuid"
)

type placelistRepositoryImpl struct {
	db *database.DB
}

func NewPlacelistRepository(db *database.DB) *placelistRepositoryImpl {
	return &placelistRepositoryImpl{db}
}

func (r *placelistRepositoryImpl) GetByID(id uuid.UUID) (*entities.Placelist, error) {
	var p *entities.Placelist
	result := r.db.First(&p, "id = ? AND deleted IS NULL", id)
	return p, result.Error
}

func (r *placelistRepositoryImpl) Create(p *entities.Placelist) error {
	result := r.db.Create(&p)
	return result.Error
}

func (r *placelistRepositoryImpl) Update(p *entities.Placelist) error {
	result := r.db.Save(&p)
	return result.Error
}

func (r *placelistRepositoryImpl) Delete(p *entities.Placelist) error {
	result := r.db.First(&p)
	if result.Error != nil {
		return result.Error
	}

	p.DeletedAt = time.Now()

	result = r.db.Save(&p)
	return result.Error
}
