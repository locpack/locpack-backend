package repository

import (
	"locpack-backend/internal/storage"
	"locpack-backend/internal/storage/entity"
	"locpack-backend/pkg/adapter"
)

type packRepoImpl struct {
	db adapter.Database
}

func NewPackRepository(db adapter.Database) storage.PackRepository {
	return &packRepoImpl{db}
}

func (r *packRepoImpl) GetByPublicIDFull(id string) (entity.Pack, error) {
	var p entity.Pack
	result := r.db.Preload("FollowedUsers").Preload("Author").First(&p, "public_id = ?", id)
	return p, result.Error
}

func (r *packRepoImpl) GetByNameOrAuthorFull(query string) ([]entity.Pack, error) {
	var p []entity.Pack
	result := r.db.
		Joins("JOIN users ON users.id = packs.author_id").
		Preload("FollowedUsers").
		Preload("Author").
		Where("LOWER(packs.name) LIKE LOWER(?) OR LOWER(packs.public_id) LIKE LOWER(?) OR LOWER(users.public_id) LIKE LOWER(?)", "%"+query+"%", "%"+query+"%", "%"+query+"%").
		Find(&p)
	return p, result.Error
}

func (r *packRepoImpl) Create(p entity.Pack) error {
	createErr := r.db.Create(&p).Error
	return createErr
}

func (r *packRepoImpl) Update(p entity.Pack) error {
	result := r.db.Save(&p)
	return result.Error
}
