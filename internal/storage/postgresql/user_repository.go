package postgresql

import (
	"placelists/internal/entities"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepositoryImpl {
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
