package store

import (
	"fmt"

	"fastfit/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewConnection(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	return db, nil
}

func InitDB() {
	db, err := NewConnection("test.db")
	if err != nil {
		fmt.Printf("error opening sqlite:" + err.Error())
		return
	}

	DB = db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
