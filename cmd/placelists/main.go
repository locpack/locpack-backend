package main

import (
	v1 "placelists/api/v1"
	"placelists/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	placeService := services.NewPlaceService()
	placelistService := services.NewPlacelistService()
	userService := services.NewUserService()

	placeController := v1.NewPlaceController(placeService)
	placelistController := v1.NewPlacelistController(placelistService)
	userController := v1.NewUserController(userService)

	apiV1 := r.Group("/api/v1")

	placeController.RegisterRoutes(apiV1)
	placelistController.RegisterRoutes(apiV1)
	userController.RegisterRoutes(apiV1)

	r.Run("8082")
}
