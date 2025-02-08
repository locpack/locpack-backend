package repositories

import (
	"placelists/internal/entities"
	"placelists/internal/storage/database"
)

type userRepositoryImpl struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db}
}

func (r *userRepositoryImpl) GetByUsername(username string) (*entities.User, error) {
	var u *entities.User
	result := r.db.First(&u, "username = ? AND deleted_at IS NULL", username)
	return u, result.Error
}

func (r *userRepositoryImpl) Create(u *entities.User) error {
	result := r.db.Create(&u)
	return result.Error
}

func (r *userRepositoryImpl) Update(u *entities.User) error {
	result := r.db.Save(&u)
	return result.Error
}
