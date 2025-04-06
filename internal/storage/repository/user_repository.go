package repository

import (
	"placelists-back/internal/storage"
	"placelists-back/internal/storage/entity"
	"placelists-back/pkg/adapter"
)

type userRepositoryImpl struct {
	db adapter.Database
}

func NewUserRepository(db adapter.Database) storage.UserRepository {
	return &userRepositoryImpl{db}
}

func (r *userRepositoryImpl) GetByPublicID(id string) (entity.User, error) {
	var u entity.User
	result := r.db.First(u, "public_id = ?", id)
	return u, result.Error
}

func (r *userRepositoryImpl) GetByPublicIDFull(id string) (entity.User, error) {
	var u entity.User
	result := r.db.Preload("FollowedPlacelists").Preload("CreatedPlacelists").First(u, "public_id = ?", id)
	return u, result.Error
}

func (r *userRepositoryImpl) Create(u entity.User) error {
	result := r.db.Create(u)
	return result.Error
}

func (r *userRepositoryImpl) Update(u entity.User) error {
	result := r.db.Save(u)
	return result.Error
}
