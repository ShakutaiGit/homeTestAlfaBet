package persistence

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"homeTestAlfaBet/internal/domain"
)

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("events.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&domain.Event{})
	return db
}
