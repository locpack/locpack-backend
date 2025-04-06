package api

import (
	"github.com/gin-gonic/gin"
	"placelists-back/pkg/adapter"
	"placelists-back/pkg/cfg"
)

func New(cfg *cfg.API) adapter.API {
	gin.SetMode(cfg.Mode)
	return gin.Default()
}
