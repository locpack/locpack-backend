package domain

import (
	"placelists/internal/service/models"
	"placelists/internal/storage"
)

type userService struct {
	r storage.UserRepository
}

func NewUserService(r storage.UserRepository) *userService {
	return &userService{r: r}
}

func (s *userService) GetByUsername(username string) (*models.User, error) {
	u, err := s.r.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	foundUser := &models.User{ID: u.PublicID, Username: u.Username}

	return foundUser, nil
}

func (s *userService) UpdateByUsername(username string, uu *models.UserUpdate) error {
	u, err := s.r.GetByUsername(username)
	if err != nil {
		return err
	}

	u.Username = uu.Username
	u.PublicID = uu.Username

	err = s.r.Update(u)

	return err
}
