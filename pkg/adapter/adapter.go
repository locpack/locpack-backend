package adapter

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database = *gorm.DB

type API = *gin.Engine
type APIHandler = gin.HandlerFunc
type APIContext = *gin.Context
