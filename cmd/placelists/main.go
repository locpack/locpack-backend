package main

import (
	"fmt"
	"placelists/internal/service/domain"
	"placelists/internal/storage/database"
	"placelists/internal/storage/entities"
	"placelists/internal/storage/repositories"
)

func main() {
	//db := database.New("host=localhost user=postgres password=postgres dbname=postgres port=5432")
	db := database.New("test.db")
	r := repositories.NewRepository(db)
	domain.NewService(r)

	var p *[]entities.Placelist
	result := db.Joins("Users").
		// Joins("JOIN users ON users.id = user_placelists.user_id").
		Where("name = ? OR username = ?", "name", "username").
		Find(&p)
	fmt.Println(result.Error)
	// user := &entities.User{
	// 	ID:       uuid.New(),
	// 	Username: "second",
	// 	PublicID: "second",
	// }
	// r.User.Create(user)

	// place := &models.PlaceCreate{
	// 	Name:    "megaplace",
	// 	Address: "nooo",
	// 	Visited: false,
	// }
	// s.Place.Create("second", place)

	// s.User.UpdateByUsername("master", &models.UserUpdate{Username: "second"})

	// r.Create(&entities.User{Username: "master", ID: uuid.New()})
	// fmt.Println(u.Username)

	// var repository *storage.Repository = postgresql.NewRepository(db)

	// r := gin.New()

	// placeService := services.NewPlaceService()
	// placelistService := services.NewPlacelistService()
	// userService := services.NewUserService()

	// placeController := v1.NewPlaceController(placeService)
	// placelistController := v1.NewPlacelistController(placelistService)
	// userController := v1.NewUserController(userService)

	// apiV1 := r.Group("/api/v1")

	// placeController.RegisterRoutes(apiV1)
	// placelistController.RegisterRoutes(apiV1)
	// userController.RegisterRoutes(apiV1)

	// // var placeRepository repositories.PlaceRepository = repositories.NewPlaceRepository()

	// r.Run("8082")
}
