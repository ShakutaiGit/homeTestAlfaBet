package repository

import (
	"gorm.io/gorm"
	"homeTestAlfaBet/internal/domain"
)

type EventRepository interface {
	Create(event *domain.Event) error
	GetAll() ([]domain.Event, error)
	GetByID(id uint) (domain.Event, error)
	Update(event *domain.Event) error
	Delete(id uint) error
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
