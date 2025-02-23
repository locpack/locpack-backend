package api

import (
	"placelists/internal/server"
	"placelists/internal/server/dtos"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New(controller *server.Controller) *gin.Engine {
	g := gin.Default()
	router := g.Group("/api")
	{
		v1Router := router.Group("/v1")
		{
			placeRouter := v1Router.Group("/places")
			{
				placeRouter.GET("/", controller.Place.GetPlacesByQuery)
				placeRouter.POST("/", controller.Place.PostPlace)
				placeRouter.GET("/:id", controller.Place.GetPlaceByID)
				placeRouter.PUT("/:id", controller.Place.PutPlaceByID)
			}
			placelistRouter := v1Router.Group("/placelists")
			{
				placelistRouter.GET("/", controller.Placelist.GetPlacelistsByQuery)
				placelistRouter.POST("/", controller.Placelist.PostPlacelist)
				placelistRouter.GET("/followed", controller.Placelist.GetPlacelistsFollowed)
				placelistRouter.GET("/created", controller.Placelist.GetPlacelistsCreated)
				placelistRouter.GET("/:id", controller.Placelist.GetPlacelistByID)
				placelistRouter.PUT("/:id", controller.Placelist.PutPlacelistByID)
			}
			userRouter := v1Router.Group("/users")
			{
				userRouter.GET("/my", controller.User.GetUserMy)
				userRouter.GET("/:id", controller.User.GetUserByID)
				userRouter.PUT("/:id", controller.User.PutUserByID)
			}
		}
	}
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return g
}

func SuccessResponse(ctx *gin.Context, statusCode int, data any) {
	ctx.JSON(statusCode, dtos.ResponseWrapper[any]{
		Data:   data,
		Meta:   dtos.Meta{Success: true},
		Errors: []dtos.Error{},
	})
}

func ErrorResponse(ctx *gin.Context, statusCode int, errors []dtos.Error) {
	ctx.JSON(statusCode, dtos.ResponseWrapper[any]{
		Data:   nil,
		Meta:   dtos.Meta{Success: false},
		Errors: errors,
	})
}
