package domain

import (
	"placelists/internal/service"
	"placelists/internal/storage"
)

func NewService(repository *storage.Repository) *service.Service {
	return &service.Service{
		Place: NewPlaceService(repository.Place, repository.User),
		// Placelist: NewPlacelistService(r.Placelist),
		User: NewUserService(repository.User),
	}
}
