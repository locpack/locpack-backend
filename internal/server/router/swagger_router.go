package router

import (
	"locpack-backend/pkg/adapter/swagger"

	"locpack-backend/pkg/adapter"
)

func NewSwaggerRouter(api adapter.API) {
	api.GET("/swagger/*any", swagger.GetHandler())
}
