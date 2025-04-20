package fake

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"locpack-backend/internal/storage/entity"
	"locpack-backend/pkg/adapter"
)

func NewDatabase() (adapter.Database, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"))
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
