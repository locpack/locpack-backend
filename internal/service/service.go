package service

import "placelists/internal/service/models"

type Service struct {
	Place     PlaceService
	Placelist PlacelistService
	User      UserService
}

type PlaceService interface {
	GetByPublicID(placeID string, userID string) (*models.Place, error)
	GetByNameOrAddress(query string, userID string) (*[]models.Place, error)
	Create(userID string, pc *models.PlaceCreate) error
	UpdateByPublicID(placeID string, userID string, pu *models.PlaceUpdate) error
}

type PlacelistService interface {
	GetByPublicID(placelistID string, userID string) (*models.Placelist, error)
	GetByNameOrAuthor(query string, userID string) (*[]models.Placelist, error)
	GetFollowedByUserID(userID string) (*[]models.Placelist, error)
	GetCreatedByUserID(userID string) (*[]models.Placelist, error)
	GetPlacesByPublicID(placelistID string, userID string) (*[]models.Place, error)
	Create(userID string, pc *models.PlacelistCreate) error
	UpdateByPublicID(placelistID string, userID string, pu *models.PlacelistUpdate) error
}

type UserService interface {
	GetByPublicID(publicID string) (*models.User, error)
	UpdateByPublicID(publicID string, uu *models.UserUpdate) error
}
