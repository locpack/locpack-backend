package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"locpack-backend/pkg/types"
)

type Database = *gorm.DB

type API = *gin.Engine
type APIHandler = gin.HandlerFunc
type APIContext = *gin.Context

type Auth interface {
	Register(username string, email string, password string) (uuid.UUID, error)
	Login(username string, password string) (types.AccessToken, error)
	Refresh(value string) (types.AccessToken, error)
	DecodeToken(accessToken string) (types.Token, error)
}
