package api

import (
	"github.com/gin-gonic/gin"
	"locpack-backend/pkg/adapter"
	"locpack-backend/pkg/cfg"
)

func New(cfg *cfg.API) adapter.API {
	gin.SetMode(cfg.Mode)
	return gin.Default()
}
