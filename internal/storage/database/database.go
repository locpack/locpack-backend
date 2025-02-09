package database

import (
	"placelists/internal/entities"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

type Provider = int

const (
	PostgreSQLProvider Provider = iota
	SQLiteProvider
)

func New(p Provider, dsn string) *DB {
	providers := map[Provider]func(dsn string) gorm.Dialector{
		PostgreSQLProvider: postgres.Open,
		SQLiteProvider:     sqlite.Open,
	}
	db, _ := gorm.Open(providers[p](dsn), &gorm.Config{})

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
