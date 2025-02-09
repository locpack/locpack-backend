package database

import (
	"placelists/internal/storage/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

type Tx struct {
	tx *gorm.DB
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

func (t *Tx) Commit() error {
	return t.tx.Commit().Error
}

func (t *Tx) Rollback() error {
	return t.tx.Rollback().Error
}

func (t *Tx) DB() *gorm.DB {
	return t.tx
}

func (db *DB) Transaction(fn func(tx *Tx) error) error {
	gormTx := db.DB.Begin()
	if gormTx.Error != nil {
		return gormTx.Error
	}

	tx := &Tx{tx: gormTx}

	err := fn(tx)
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
