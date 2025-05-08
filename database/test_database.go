package database

import (
	"book-api/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func ConnectTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to in-memory test database")
	}

	db.AutoMigrate(&models.Book{})
	DB = db // set global DB for handlers/services to use
	return db
}
