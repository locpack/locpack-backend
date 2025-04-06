package database

import (
	"placelists-back/internal/storage/entity"
	"placelists-back/pkg/adapter"
	"placelists-back/pkg/cfg"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg *cfg.Database) (adapter.Database, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN))
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Place{},
		&entity.Placelist{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
