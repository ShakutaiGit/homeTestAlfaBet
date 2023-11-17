package service

import (
	"homeTestAlfaBet/internal/domain"
	"homeTestAlfaBet/internal/repository"
)

// EventService defines the interface for event service
type EventService interface {
	ScheduleEvent(event *domain.Event) error
	GetEvents(location, sortBy string) ([]domain.Event, error)
	GetEventByID(id uint) (domain.Event, error)
	UpdateEvent(event *domain.Event) error
	DeleteEvent(id uint) error
}

type eventService struct {
	repo repository.EventRepository
}

// NewEventService creates a new instance of EventService
func NewEventService(repo repository.EventRepository) EventService {
	return &eventService{repo: repo}
}

func (s *eventService) ScheduleEvent(event *domain.Event) error {
	return s.repo.Create(event)
}

func (s *eventService) GetEvents(location, sortBy string) ([]domain.Event, error) {
	return s.repo.GetEvents(location, sortBy)
}

func (s *eventService) GetEventByID(id uint) (domain.Event, error) {
	return s.repo.GetByID(id)
}

func (s *eventService) UpdateEvent(event *domain.Event) error {
	return s.repo.Update(event)
}

func (s *eventService) DeleteEvent(id uint) error {
	return s.repo.Delete(id)
}
