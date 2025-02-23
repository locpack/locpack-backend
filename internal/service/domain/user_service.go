package domain

import (
	"placelists/internal/service"
	"placelists/internal/service/models"
	"placelists/internal/storage"

	"github.com/jinzhu/copier"
)

type userServiceImpl struct {
	repository storage.UserRepository
}

func NewUserService(repository storage.UserRepository) service.UserService {
	return &userServiceImpl{repository}
}

func (s *userServiceImpl) GetByID(id string) (models.User, error) {
	userEntity, err := s.repository.GetByPublicID(id)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{}
	err = copier.Copy(&userEntity, &user)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}

func (s *userServiceImpl) UpdateByID(id string, uu models.UserUpdate) (models.User, error) {
	userEntity, err := s.repository.GetByPublicID(id)
	if err != nil {
		return models.User{}, err
	}

	userEntity.Username = uu.Username
	userEntity.PublicID = uu.Username

	err = s.repository.Update(userEntity)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{}
	err = copier.Copy(&userEntity, &user)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}
