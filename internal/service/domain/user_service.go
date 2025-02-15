package domain

import (
	"placelists/internal/service"
	"placelists/internal/service/models"
	"placelists/internal/storage"
)

type userService struct {
	repository storage.UserRepository
}

func NewUserService(repository storage.UserRepository) service.UserService {
	return &userService{repository}
}

func (s *userService) GetByID(id string) (models.User, error) {
	u, err := s.repository.GetByPublicID(id)
	if err != nil {
		return models.User{}, err
	}

	foundUser := models.User{
		ID:       u.PublicID,
		Username: u.Username,
	}

	return foundUser, nil
}

func (s *userService) UpdateByID(id string, uu models.UserUpdate) error {
	u, err := s.repository.GetByPublicID(id)
	if err != nil {
		return err
	}

	u.Username = uu.Username
	u.PublicID = uu.Username

	err = s.repository.Update(u)

	return err
}
