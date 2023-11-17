package service

import (
	"homeTestAlfaBet/internal/domain"
	"homeTestAlfaBet/internal/mocks"
	"testing"
)

// Test for ScheduleEvent method
func TestEventService_ScheduleEvent(t *testing.T) {
	mockRepo := new(mocks.EventRepositoryMock)
	mockRepo.MockCreate = func(event *domain.Event) error {
		return nil // Assume successful event scheduling
	}

	svc := NewEventService(mockRepo)

	err := svc.ScheduleEvent(&domain.Event{Name: "Scheduled Event"})
	if err != nil {
		t.Errorf("ScheduleEvent() error = %v, wantErr false", err)
	}
}

// Test for GetEvents method
func TestEventService_GetEvents(t *testing.T) {
	mockRepo := new(mocks.EventRepositoryMock)
	mockRepo.MockGetEvents = func(location, sortBy string) ([]domain.Event, error) {
		return []domain.Event{{Name: "Event 1"}}, nil
	}

	svc := NewEventService(mockRepo)

	events, err := svc.GetEvents("", "") // Test with empty filters and sort options
	if err != nil {
		t.Errorf("GetEvents() error = %v, wantErr false", err)
	}
	if len(events) != 1 {
		t.Errorf("GetEvents() got = %v, want %v", len(events), 1)
	}
}

// Additional tests for GetEventByID, UpdateEvent, and DeleteEvent can be written similarly...
