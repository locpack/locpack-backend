package database

import (
	"placelists/internal/storage/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func New(dsn string) *DB {
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
