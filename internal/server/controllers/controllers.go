package controllers

import (
	"placelists/internal/server"
	"placelists/internal/service"
)

func NewController(service *service.Service) *server.Controller {
	return &server.Controller{
		Place:     NewPlaceController(service.Place),
		Placelist: NewPlacelistController(service.Placelist),
		User:      NewUserController(service.User),
	}
}
