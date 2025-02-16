package main

import (
	_ "placelists/docs/swagger"
	"placelists/internal/server/controllers"
	"placelists/internal/service/domain"
	"placelists/internal/storage/repositories"
	"placelists/pkg/api"
	"placelists/pkg/database"
)

// @title Placelists
// @version 1.1
// @host localhost:8082
// @BasePath /api

func main() {
	//db := database.New("host=localhost user=postgres password=postgres dbname=postgres port=5432")
	db := database.New("test.db")
	r := repositories.NewRepository(db)
	s := domain.NewService(r)
	c := controllers.NewController(s)
	api := api.New(c)
	api.Run("localhost:8082")
}
