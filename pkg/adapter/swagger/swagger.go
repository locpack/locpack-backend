package swagger

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"locpack-backend/pkg/adapter"
)

func GetHandler() adapter.APIHandler {
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
