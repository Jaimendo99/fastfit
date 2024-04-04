package store

import (
	"fmt"

	"fastfit/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewConnection(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	return db, nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
