package fakes

import (
	"placelists/internal/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func New() *DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	db.AutoMigrate(
		&entities.User{},
		&entities.Place{},
		&entities.Placelist{},
		&entities.UserPlace{},
		&entities.PlacelistPlace{},
		&entities.UserPlacelist{},
	)

	return &DB{db}
}
