package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func New() *DB {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	return &DB{db}
}
