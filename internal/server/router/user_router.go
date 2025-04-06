package router

import (
	"placelists-back/internal/server"
	"placelists-back/pkg/adapter"
)

func NewUserRouter(api adapter.API, c server.UserController) {
	api.GET("/api/v1/users/my", c.GetUserMy)
	api.GET("/api/v1/users/:id", c.GetUserByID)
	api.PUT("/api/v1/users/:id", c.PutUserByID)
}
