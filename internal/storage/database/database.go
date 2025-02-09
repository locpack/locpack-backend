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
		&entities.UserPlace{},
		&entities.PlacelistPlace{},
		&entities.UserPlacelist{},
	)

	return &DB{db}
}

func (db *DB) Begin() (*DB, error) {
	tx := db.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &DB{tx}, nil
}

func (db *DB) Commit() error {
	return db.DB.Commit().Error
}

func (db *DB) Rollback() error {
	return db.DB.Rollback().Error
}

func (db *DB) Transaction(fn func(tx *DB) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		return commitErr
	}

	return nil
}
