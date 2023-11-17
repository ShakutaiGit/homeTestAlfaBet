package mocks

import (
	"homeTestAlfaBet/internal/domain"
	"time"
)

type EventRepositoryMock struct {
	MockCreate                    func(event *domain.Event) error
	MockUpdate                    func(event *domain.Event) error
	MockDelete                    func(id uint) error
	MockGetEvents                 func(location, sortBy string) ([]domain.Event, error)
	MockGetEventsHappeningBetween func(startTime, endTime time.Time) ([]domain.Event, error)
	MockGetAll                    func() ([]domain.Event, error)
	MockGetByID                   func(id uint) (domain.Event, error)
}

func (m *EventRepositoryMock) Create(event *domain.Event) error {
	return m.MockCreate(event)
}

func (m *EventRepositoryMock) Update(event *domain.Event) error {
	return m.MockUpdate(event)
}

func (m *EventRepositoryMock) Delete(id uint) error {
	return m.MockDelete(id)
}

func (m *EventRepositoryMock) GetEvents(location, sortBy string) ([]domain.Event, error) {
	return m.MockGetEvents(location, sortBy)
}

func (m *EventRepositoryMock) GetEventsHappeningBetween(startTime, endTime time.Time) ([]domain.Event, error) {
	return m.MockGetEventsHappeningBetween(startTime, endTime)
}

func (m *EventRepositoryMock) GetAll() ([]domain.Event, error) {
	return m.MockGetAll()
}

func (m *EventRepositoryMock) GetByID(id uint) (domain.Event, error) {
	return m.MockGetByID(id)
}
