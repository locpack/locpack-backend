package service

import "placelists/internal/service/models"

type Service struct {
	Place     PlaceService
	Placelist PlacelistService
	User      UserService
}

type PlaceService interface {
	GetByPublicIDWithUser(publicID string, userPublicID string) (*models.Place, error)
	GetByNameOrAddressWithUser(query string, userPublicID string) (*[]models.Place, error)
	Create(userPublicID string, pc *models.PlaceCreate) error
	UpdateByPublicIDWithUser(publicID string, userPublicID string, pu *models.PlaceUpdate) error
}

type PlacelistService interface {
	GetByNameOrUsernameWithUser(query string, userPublicID string) (*[]models.Placelist, error)
	Create(userPublicID string, pc *models.PlacelistCreate) error
}

type UserService interface {
	GetByPublicID(publicID string) (*models.User, error)
	UpdateByPublicID(publicID string, uu *models.UserUpdate) error
}
