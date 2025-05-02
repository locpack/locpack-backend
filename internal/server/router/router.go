package router

import (
	"locpack-backend/internal/server"
	"locpack-backend/internal/server/middleware"
	"locpack-backend/pkg/adapter"
	"locpack-backend/pkg/adapter/swagger"
)

func New(
	api adapter.API,
	authAdapter adapter.Auth,
	packController server.PackController,
	placeController server.PlaceController,
	userController server.UserController,
	authController server.AuthController,
) {
	public := api.Group("")
	{
		public.GET("/api/v1/packs", packController.GetPacksByQuery)
		public.GET("/api/v1/packs/:id", packController.GetPackByID)
		public.GET("/api/v1/places", placeController.GetPlacesByQuery)
		public.GET("/api/v1/places/:id", placeController.GetPlaceByID)
		public.GET("/api/v1/users/:id", userController.GetUserByID)
		public.GET("/swagger/*any", swagger.GetHandler())
	}

	auth := api.Group("")
	auth.Use(middleware.AuthenticatedMiddleware(authAdapter))
	{
		auth.POST("/api/v1/packs", packController.PostPack)
		auth.GET("/api/v1/packs/followed", packController.GetPacksFollowed)
		auth.GET("/api/v1/packs/created", packController.GetPacksCreated)
		auth.PUT("/api/v1/packs/:id", packController.PutPackByID)
		auth.POST("/api/v1/places", placeController.PostPlace)
		auth.PUT("/api/v1/places/:id", placeController.PutPlaceByID)
		auth.GET("/api/v1/users/my", userController.GetUserMy)
		auth.POST("/api/v1/auth/refresh", authController.Refresh)
	}

	notAuth := api.Group("")
	notAuth.Use(middleware.NotAuthenticatedMiddleware())
	{
		notAuth.POST("/api/v1/auth/register", authController.Register)
		notAuth.POST("/api/v1/auth/login", authController.Login)
	}
}
