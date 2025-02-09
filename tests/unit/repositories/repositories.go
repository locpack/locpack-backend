package repositories_test

import (
	"placelists/internal/entities"
	"placelists/internal/storage/database"
)

func InitDB() *database.DB {
	return database.New("host=localhost user=postgres password=postgres dbname=postgres port=5432")
}

func DropDB(db *database.DB) {
	db.Migrator().DropTable(
		&entities.User{},
		&entities.Place{},
		&entities.Placelist{},
		&entities.UserPlace{},
		&entities.PlacelistPlace{},
		&entities.UserPlacelist{},
	)
}
