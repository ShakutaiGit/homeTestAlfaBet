package domain

import (
	"time"
)

// Event represents the event structure
type Event struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:255"`
	Description string `gorm:"size:1000"`
	Location    string `gorm:"size:255"`
	StartTime   time.Time
	EndTime     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
