package api

import (
	"github.com/gin-gonic/gin"
	"homeTestAlfaBet/internal/domain"
	"homeTestAlfaBet/internal/service"
	"net/http"
	"strconv"
)

// EventHandler represents the HTTP handler for events
type EventHandler struct {
	service service.EventService
}

// NewEventHandler creates a new instance of EventHandler
func NewEventHandler(s service.EventService) *EventHandler {
	return &EventHandler{service: s}
}

func (h *EventHandler) ScheduleEvent(ctx Context) {
	var event domain.Event
	if err := ctx.BindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.ScheduleEvent(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (h *EventHandler) GetAllEvents(ctx Context) {
	location := ctx.Query("location")
	sortBy := ctx.Query("sortBy")

	events, err := h.service.GetEvents(location, sortBy)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func (h *EventHandler) GetEvent(ctx Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	event, err := h.service.GetEventByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (h *EventHandler) UpdateEvent(ctx Context) {
	var event domain.Event
	if err := ctx.BindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateEvent(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (h *EventHandler) DeleteEvent(ctx Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.service.DeleteEvent(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
