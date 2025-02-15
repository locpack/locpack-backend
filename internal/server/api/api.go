package api

import (
	"placelists/internal/server"
	"placelists/internal/server/dtos"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, statusCode int, data any, errors []dtos.Error) {
	ctx.JSON(statusCode, dtos.ResponseWrapper[any]{
		Data:   data,
		Meta:   dtos.Meta{Success: len(errors) == 0},
		Errors: errors,
	})
}

func New(controller *server.Controller) *gin.Engine {
	g := gin.New()
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
	return g
}
