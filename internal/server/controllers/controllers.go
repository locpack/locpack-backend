package controllers

import (
	"placelists-back/internal/server"
	"placelists-back/internal/service"
)

func NewController(service *service.Service) *server.Controller {
	return &server.Controller{
		Place:     NewPlaceController(service.Place),
		Placelist: NewPlacelistController(service.Placelist),
		User:      NewUserController(service.User),
	}
}
