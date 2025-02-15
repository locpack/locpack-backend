package main

import (
	"placelists/internal/server/api"
	"placelists/internal/server/controllers"
	"placelists/internal/service/domain"
	"placelists/internal/storage/repositories"
	"placelists/pkg/database"
)

func main() {
	//db := database.New("host=localhost user=postgres password=postgres dbname=postgres port=5432")
	db := database.New("test.db")
	r := repositories.NewRepository(db)
	s := domain.NewService(r)
	c := controllers.NewController(s)
	api := api.New(c)
	api.Run("8082")
}
