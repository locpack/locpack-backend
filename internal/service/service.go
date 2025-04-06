package service

import "placelists-back/internal/service/model"

type PlaceService interface {
	GetByID(placeID string, userID string) (model.Place, error)
	GetByNameOrAddress(query string, userID string) ([]model.Place, error)
	Create(userID string, pc model.PlaceCreate) (model.Place, error)
	UpdateByID(placeID string, userID string, pu model.PlaceUpdate) (model.Place, error)
}

type PlacelistService interface {
	GetByID(placelistID string, userID string) (model.Placelist, error)
	GetByNameOrAuthor(query string, userID string) ([]model.Placelist, error)
	GetFollowedByUserID(userID string) ([]model.Placelist, error)
	GetCreatedByUserID(userID string) ([]model.Placelist, error)
	GetPlacesByID(placelistID string, userID string) ([]model.Place, error)
	Create(userID string, pc model.PlacelistCreate) (model.Placelist, error)
	UpdateByID(placelistID string, userID string, pu model.PlacelistUpdate) (model.Placelist, error)
}

type UserService interface {
	GetByID(id string) (model.User, error)
	UpdateByID(id string, uu model.UserUpdate) (model.User, error)
}
