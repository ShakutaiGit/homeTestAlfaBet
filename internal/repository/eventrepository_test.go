package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"homeTestAlfaBet/internal/domain"
	"testing"
)

func TestGetAll(t *testing.T) {
	// Setup in-memory database for testing
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	err = db.AutoMigrate(&domain.Event{})
	if err != nil {
		return
	}

	// Setup repository
	repo := NewEventRepository(db)

	// Seed with test data
	db.Create(&domain.Event{Name: "Test Event", Location: "Test Location"})

	// Test GetAll
	events, err := repo.GetAll()
	if err != nil {
		t.Errorf("Failed to get all events: %v", err)
	}
	if len(events) != 1 {
		t.Errorf("Expected 1 event, got %d", len(events))
	}
}
