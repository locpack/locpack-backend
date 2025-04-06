package router

import (
	"placelists-back/internal/server"
	"placelists-back/pkg/adapter"
)

func NewPlacelistRouter(api adapter.API, c server.PlacelistController) {
	api.GET("/api/v1/placelists", c.GetPlacelistsByQuery)
	api.POST("/api/v1/placelists", c.PostPlacelist)
	api.GET("/api/v1/placelists/followed", c.GetPlacelistsFollowed)
	api.GET("/api/v1/placelists/created", c.GetPlacelistsCreated)
	api.GET("/api/v1/placelists/:id", c.GetPlacelistByID)
	api.PUT("/api/v1/placelists/:id", c.PutPlacelistByID)
}
