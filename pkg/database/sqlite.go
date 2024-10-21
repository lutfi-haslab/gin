package database

import (
	"clean-api/internal/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormDB(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the schemas
	err = db.AutoMigrate(
		&domain.User{},
		&domain.Product{},
		&domain.Order{},
		&domain.OrderItem{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
