package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	_ "locpack-backend/docs/swagger"
	"locpack-backend/internal/cfg"
	"locpack-backend/internal/server/controller"
	"locpack-backend/internal/server/router"
	"locpack-backend/internal/service/domain"
	"locpack-backend/internal/storage/repository"
	"locpack-backend/pkg/adapter/api"
	"locpack-backend/pkg/adapter/auth"
	"locpack-backend/pkg/adapter/database"
)

// @title Locpack API
// @version 1.0
// @description API for managing places, packs and users
// @contact.name Aleksey
// @contact.email a.e.sokolkov@gmail.com
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	var config cfg.Cfg
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		panic(err)
	}

	db, err := database.New(&config.Database)
	if err != nil {
		panic(err)
	}

	auth := auth.New(&config.Auth)

	placeRepository := repository.NewPlaceRepository(db)
	packRepository := repository.NewPackRepository(db)
	userRepository := repository.NewUserRepository(db)

	placeService := domain.NewPlaceService(placeRepository, userRepository)
	packService := domain.NewPackService(packRepository, placeRepository, userRepository)
	userService := domain.NewUserService(userRepository)
	authService := domain.NewAuthService(auth, userRepository)

	placeController := controller.NewPlaceController(placeService)
	packController := controller.NewPackController(packService)
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService)

	server := api.New(&config.API)

	router.New(server, auth, packController, placeController, userController, authController)

	err = server.Run(config.API.Address)
	if err != nil {
		panic(err)
	}
}
