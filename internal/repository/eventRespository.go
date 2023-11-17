package repository

import (
	"gorm.io/gorm"
	"homeTestAlfaBet/internal/domain"
	"time"
)

type EventRepository interface {
	Create(event *domain.Event) error
	GetAll() ([]domain.Event, error)
	GetByID(id uint) (domain.Event, error)
	Update(event *domain.Event) error
	Delete(id uint) error
	GetEvents(location, sortBy string) ([]domain.Event, error)
	GetEventsHappeningBetween(startTime, endTime time.Time) ([]domain.Event, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db: db}
}

func (repo *eventRepository) Create(event *domain.Event) error {
	return repo.db.Create(event).Error
}

func (repo *eventRepository) GetAll() ([]domain.Event, error) {
	var events []domain.Event
	err := repo.db.Find(&events).Error
	return events, err
}

func (repo *eventRepository) GetByID(id uint) (domain.Event, error) {
	var event domain.Event
	err := repo.db.First(&event, id).Error
	return event, err
}

func (repo *eventRepository) Update(event *domain.Event) error {
	return repo.db.Save(event).Error
}

func (repo *eventRepository) Delete(id uint) error {
	return repo.db.Delete(&domain.Event{}, id).Error
}

func (repo *eventRepository) GetEvents(location, sortBy string) ([]domain.Event, error) {
	query := repo.db.Model(&domain.Event{})

	// Apply location filter if provided
	if location != "" {
		query = query.Where("location = ?", location)
	}

	// Apply sorting based on the sortBy parameter
	switch sortBy {
	case "date":
		query = query.Order("start_time asc")
	case "popularity":
		query = query.Order("participants desc")
	case "creation":
		query = query.Order("created_at desc")
	}

	var events []domain.Event
	err := query.Find(&events).Error
	return events, err
}

func (repo *eventRepository) GetEventsHappeningBetween(startTime, endTime time.Time) ([]domain.Event, error) {
	var events []domain.Event
	err := repo.db.Where("start_time >= ? AND start_time <= ?", startTime, endTime).Find(&events).Error
	return events, err
}
