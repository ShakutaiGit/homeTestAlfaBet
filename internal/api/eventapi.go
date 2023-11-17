package api

// EventAPI defines the interface for event HTTP handlers
type EventAPI interface {
	ScheduleEvent(ctx Context)
	GetAllEvents(ctx Context)
	GetEvent(ctx Context)
	UpdateEvent(ctx Context)
	DeleteEvent(ctx Context)
}
