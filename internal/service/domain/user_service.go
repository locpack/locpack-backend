package domain

import (
	"locpack-backend/internal/service"
	"locpack-backend/internal/service/model"
	"locpack-backend/internal/storage"

	"github.com/jinzhu/copier"
)

type userServiceImpl struct {
	repository storage.UserRepository
}

func NewUserService(repository storage.UserRepository) service.UserService {
	return &userServiceImpl{repository}
}

func (s *userServiceImpl) GetByID(id string) (model.User, error) {
	userEntity, err := s.repository.GetByPublicID(id)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{}
	err = copier.Copy(&userEntity, &user)
	if err != nil {
		return model.User{}, err
	}

	return user, err
}

func (s *userServiceImpl) UpdateByID(id string, uu model.UserUpdate) (model.User, error) {
	userEntity, err := s.repository.GetByPublicID(id)
	if err != nil {
		return model.User{}, err
	}

	userEntity.Username = uu.Username
	userEntity.PublicID = uu.Username

	err = s.repository.Update(userEntity)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{}
	err = copier.Copy(&userEntity, &user)
	if err != nil {
		return model.User{}, err
	}

	return user, err
}
