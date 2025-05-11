package repository

import (
	"locpack-backend/internal/storage"
	"locpack-backend/internal/storage/entity"
	"locpack-backend/pkg/adapter"
)

type userRepoImpl struct {
	db adapter.Database
}

func NewUserRepository(db adapter.Database) storage.UserRepository {
	return &userRepoImpl{db}
}

func (r *userRepoImpl) GetByPublicID(id string) (entity.User, error) {
	var u entity.User
	result := r.db.First(&u, "lower(public_id) = lower(?)", id)
	return u, result.Error
}

func (r *userRepoImpl) GetByPublicIDFull(id string) (entity.User, error) {
	var u entity.User
	result := r.db.
		Preload("FollowedPacks").
		Preload("FollowedPacks.Author").
		Preload("CreatedPacks").
		Preload("CreatedPacks.Author").
		First(&u, "lower(public_id) = lower(?)", id)
	return u, result.Error
}

func (r *userRepoImpl) Create(u entity.User) error {
	result := r.db.Create(&u)
	return result.Error
}
