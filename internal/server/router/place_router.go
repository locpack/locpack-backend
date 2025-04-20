package router

import (
	"locpack-backend/internal/server"
	"locpack-backend/pkg/adapter"
)

func NewPlaceRouter(api adapter.API, c server.PlaceController) {
	api.GET("/api/v1/places", c.GetPlacesByQuery)
	api.POST("/api/v1/places", c.PostPlace)
	api.GET("/api/v1/places/:id", c.GetPlaceByID)
	api.PUT("/api/v1/places/:id", c.PutPlaceByID)
}
