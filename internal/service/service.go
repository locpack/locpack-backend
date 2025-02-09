package service

import "placelists/internal/service/models"

type Service struct {
	// Place     PlaceService
	// Placelist PlacelistService
	User UserService
}

// type PlaceService interface {
// }

// type PlacelistService interface {
// }

type UserService interface {
	GetByUsername(username string) (*models.User, error)
	UpdateByUsername(username string, uu *models.UserUpdate) error
}
