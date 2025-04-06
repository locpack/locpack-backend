package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	_ "placelists-back/docs/swagger"
	"placelists-back/internal/cfg"
	"placelists-back/internal/server/controller"
	"placelists-back/internal/server/router"
	"placelists-back/internal/service/domain"
	"placelists-back/internal/storage/repository"
	"placelists-back/pkg/adapter/api"
	"placelists-back/pkg/adapter/database"
)

// @title Placelists API
// @version 1.0
// @description API for managing places, placelists and users
// @contact.name Aleksey
// @contact.email a.e.sokolkov@gmail.com
// @host localhost:8080

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

	placeRepository := repository.NewPlaceRepository(db)
	placelistRepository := repository.NewPlacelistRepository(db)
	userRepository := repository.NewUserRepository(db)

	placeService := domain.NewPlaceService(placeRepository, userRepository)
	placelistService := domain.NewPlacelistService(placelistRepository, placeRepository, userRepository)
	userService := domain.NewUserService(userRepository)

	placeController := controller.NewPlaceController(placeService)
	placelistController := controller.NewPlacelistController(placelistService)
	userController := controller.NewUserController(userService)

	api := api.New(&config.API)

	router.NewPlaceRouter(api, placeController)
	router.NewPlacelistRouter(api, placelistController)
	router.NewUserRouter(api, userController)
	router.NewSwaggerRouter(api)

	err = api.Run(config.API.Address)
	if err != nil {
		panic(err)
	}
}
