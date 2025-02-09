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

func (r *placeRepositoryImpl) GetByPublicIDWithUser(publicID string, userID uuid.UUID) (*entities.Place, error) {
	var p *entities.Place
	result := r.db.
		Preload("Users", "user_id = ?", userID).
		First(&p, "public_id = ?", publicID)
	return p, result.Error
}

func (r *placeRepositoryImpl) GetByNameOrAddressWithUser(query string, userID uuid.UUID) (*[]entities.Place, error) {
	var p *[]entities.Place
	result := r.db.
		Preload("Users", "user_id = ?", userID).
		Find(&p, "lower(name) LIKE lower(?) OR lower(address) LIKE lower(?)", "%"+query+"%", "%"+query+"%")
	return p, result.Error
}

func (r *placeRepositoryImpl) Create(p *entities.Place) error {
	result := r.db.Create(&p)
	return result.Error
}

func (r *placeRepositoryImpl) Update(p *entities.Place) error {
	return r.db.Transaction(func(tx *database.DB) error {
		placeSaveResult := tx.Save(&p)
		if placeSaveResult.Error != nil {
			return placeSaveResult.Error
		}

		if len(p.Users) > 0 {
			userPlaceSaveResult := tx.Save(&p.Users[0])
			if userPlaceSaveResult.Error != nil {
				return userPlaceSaveResult.Error
			}
		}

		return nil
	})
}
