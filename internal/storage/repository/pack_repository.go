package repository

import (
	"locpack-backend/internal/storage"
	"locpack-backend/internal/storage/entity"
	"locpack-backend/pkg/adapter"
)

type packRepositoryImpl struct {
	db adapter.Database
}

func NewPackRepository(db adapter.Database) storage.PackRepository {
	return &packRepositoryImpl{db}
}

func (r *packRepositoryImpl) GetByPublicIDFull(id string) (entity.Pack, error) {
	var p entity.Pack
	result := r.db.Preload("FollowedUsers").First(p, "public_id = ?", id)
	return p, result.Error
}

func (r *packRepositoryImpl) GetByNameOrAuthorFull(query string) ([]entity.Pack, error) {
	var p []entity.Pack
	result := r.db.Preload("FollowedUsers").Find(p, "lower(name) LIKE lower(?) OR lower(author_id) LIKE lower(?)", "%"+query+"%", "%"+query+"%")
	return p, result.Error
}

func (r *packRepositoryImpl) Create(p entity.Pack) error {
	createErr := r.db.Create(p).Error
	return createErr
}

func (r *packRepositoryImpl) Update(p entity.Pack) error {
	result := r.db.Save(p)
	return result.Error
}
