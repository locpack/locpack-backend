package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"locpack-backend/pkg/adapter"
	"locpack-backend/pkg/cfg"
)

func New(cfg *cfg.API) adapter.API {
	gin.SetMode(cfg.Mode)
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("*")
	router.Use(cors.New(config))
	return router
}
