package main

import (
	apiV1 "placelists/internal/app/api/v1"
	"placelists/internal/core/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	placeService := services.NewPlaceService()
	placelistService := services.NewPlacelistService()
	userService := services.NewUserService()

	placeController := apiV1.NewPlaceController(placeService)
	placelistController := apiV1.NewPlacelistController(placelistService)
	userController := apiV1.NewUserController(userService)

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			places := v1.Group("/places")
			{
				places.GET("/", placeController.GetPlacesByQuery)
				places.POST("/", placeController.PostPlace)
				places.GET("/:id", placeController.GetPlaceByID)
				places.PUT("/:id", placeController.PutPlaceByID)
			}
			placelists := v1.Group("/placelists")
			{
				placelists.GET("/placelists", placelistController.GetPlacelistsByQuery)
				placelists.POST("/placelists", placelistController.PostPlacelist)
				placelists.GET("/placelists/followed", placelistController.GetPlacelistsFollowed)
				placelists.GET("/placelists/created", placelistController.GetPlacelistsCreated)
				placelists.GET("/placelists/:id", placelistController.GetPlacelistByID)
				placelists.PUT("/placelists/:id", placelistController.PutPlacelistByID)
				placelists.DELETE("/placelists/:id", placelistController.DeletePlacelistByID)
				placelists.GET("/placelists/:id/places", placelistController.GetPlacelistPlacesByID)
				placelists.PUT("/placelists/:id/places", placelistController.PutPlacelistPlacesByID)
			}
			users := v1.Group("/users")
			{
				users.GET("/my", userController.GetUserMy)
				users.GET("/:username", userController.GetUserByUsername)
				users.PUT("/:username", userController.PutUserByUsername)
			}
		}
	}

	r.Run("8082")
}
