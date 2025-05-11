package service

import (
	"locpack-backend/internal/service/model"
)

type PlaceService interface {
	GetByID(placeID string, userID string) (model.Place, error)
	GetByNameOrAddress(query string, userID string) ([]model.Place, error)
	Create(userID string, pc model.PlaceCreate) (model.Place, error)
	UpdateByID(placeID string, userID string, pu model.PlaceUpdate) (model.Place, error)
}

type PackService interface {
	GetByID(packID string, userID string) (model.Pack, error)
	GetByNameOrAuthor(query string, userID string) ([]model.Pack, error)
	GetFollowedByUserID(userID string) ([]model.Pack, error)
	GetCreatedByUserID(userID string) ([]model.Pack, error)
	Create(userID string, pc model.PackCreate) (model.Pack, error)
	UpdateByID(packID string, userID string, pu model.PackUpdate) (model.Pack, error)
}

type UserService interface {
	GetByID(id string) (model.User, error)
}

type AuthService interface {
	Register(register model.Register) (model.AccessToken, error)
	Login(login model.Login) (model.AccessToken, error)
	Refresh(refresh model.Refresh) (model.AccessToken, error)
}
