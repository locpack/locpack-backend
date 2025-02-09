package domain

import (
	"placelists/internal/service/models"
	"placelists/internal/storage"
)

type userService struct {
	repository storage.UserRepository
}

func NewUserService(repository storage.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) GetByPublicID(publicID string) (*models.User, error) {
	u, err := s.repository.GetByPublicID(publicID)
	if err != nil {
		return nil, err
	}

	foundUser := &models.User{ID: u.PublicID, Username: u.Username}

	return foundUser, nil
}

func (s *userService) UpdateByPublicID(publicID string, uu *models.UserUpdate) error {
	u, err := s.repository.GetByPublicID(publicID)
	if err != nil {
		return err
	}

	u.Username = uu.Username
	u.PublicID = uu.Username

	err = s.repository.Update(u)

	return err
}
