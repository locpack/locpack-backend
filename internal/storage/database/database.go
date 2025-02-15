package database

import (
	"placelists/internal/storage/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func New(dsn string) *DB {
	db, _ := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	db.AutoMigrate(
		&entities.User{},
		&entities.Place{},
		&entities.Placelist{},
	)

	return &DB{db}
}
