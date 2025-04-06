package router

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"placelists-back/pkg/adapter"
)

func NewSwaggerRouter(api adapter.API) {
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
