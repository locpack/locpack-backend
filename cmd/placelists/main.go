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

	apiV1 := r.Group("/api/v1")

	placeController.RegisterRoutes(apiV1)
	placelistController.RegisterRoutes(apiV1)
	userController.RegisterRoutes(apiV1)

	r.Run("8082")
}
