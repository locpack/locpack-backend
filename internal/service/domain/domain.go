package domain

import (
	"placelists-back/internal/service"
	"placelists-back/internal/storage"
)

func NewService(repository *storage.Repository) *service.Service {
	return &service.Service{
		Place:     NewPlaceService(repository.Place, repository.User),
		Placelist: NewPlacelistService(repository.Placelist, repository.Place, repository.User),
		User:      NewUserService(repository.User),
	}
}
