package router

import (
	"locpack-backend/internal/server"
	"locpack-backend/pkg/adapter"
)

func NewPackRouter(api adapter.API, c server.PackController) {
	api.GET("/api/v1/packs", c.GetPacksByQuery)
	api.POST("/api/v1/packs", c.PostPack)
	api.GET("/api/v1/packs/followed", c.GetPacksFollowed)
	api.GET("/api/v1/packs/created", c.GetPacksCreated)
	api.GET("/api/v1/packs/:id", c.GetPackByID)
	api.PUT("/api/v1/packs/:id", c.PutPackByID)
}
