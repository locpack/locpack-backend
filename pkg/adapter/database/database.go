package database

import (
	"locpack-backend/internal/storage/entity"
	"locpack-backend/pkg/adapter"
	"locpack-backend/pkg/cfg"

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
		&entity.Pack{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
