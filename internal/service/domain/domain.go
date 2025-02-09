package domain

import (
	"placelists/internal/service"
	"placelists/internal/storage"
)

func NewService(r *storage.Repository) *service.Service {
	return &service.Service{
		// Place:     NewPlaceService(r.Place),
		// Placelist: NewPlacelistService(r.Placelist),
		User: NewUserService(r.User),
	}
}
