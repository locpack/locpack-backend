package repositories

import (
	"placelists/internal/storage/database"
	"placelists/internal/storage/entities"

	"github.com/google/uuid"
)

type placelistRepositoryImpl struct {
	db *database.DB
}

func NewPlacelistRepository(db *database.DB) *placelistRepositoryImpl {
	return &placelistRepositoryImpl{db}
}

func (r *placelistRepositoryImpl) GetByPublicID(publicID string) (*entities.Placelist, error) {
	var p *entities.Placelist
	result := r.db.First(&p, "public_id = ?", publicID)
	return p, result.Error
}

func (r *placelistRepositoryImpl) GetByNameOrAuthorWithUser(query string, userID uuid.UUID) (*[]entities.Placelist, error) {
	var p *[]entities.Placelist
	result := r.db.
		Preload("Users", "user_id = ?", userID).
		Find(&p, "lower(name) LIKE lower(?)", "%"+query+"%")
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
