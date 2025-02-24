package main

import (
	_ "placelists-back/docs/swagger"
	"placelists-back/internal/server/controllers"
	"placelists-back/internal/service/domain"
	"placelists-back/internal/storage/repositories"
	"placelists-back/pkg/api"
	"placelists-back/pkg/database"
)

// @title Placelists
// @version 1.1
// @host 0.0.0.0:8080
// @BasePath /api

func main() {
	//db := database.New("host=localhost user=postgres password=postgres dbname=postgres port=5432")
	db := database.New("test.db")
	r := repositories.NewRepository(db)
	s := domain.NewService(r)
	c := controllers.NewController(s)
	api := api.New(c)

	err := api.Run("0.0.0.0:8080")
	if err != nil {
		return
	}
}
