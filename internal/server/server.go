package server

import (
	"github.com/gin-gonic/gin"
	"homeTestAlfaBet/internal/api"
	persistence "homeTestAlfaBet/internal/db"
	"homeTestAlfaBet/internal/repository"
	"homeTestAlfaBet/internal/service"
)

// Server represents the HTTP server
type Server struct {
	router   *gin.Engine
	eventAPI api.EventAPI // Use the EventAPI interface
}

// Ensure Server implements ServerInterface
var _ ServerInterface = (*Server)(nil)

// NewServer creates a new HTTP server with all dependencies set up
func NewServer() *Server {
	db := persistence.SetupDatabase()
	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventHandler := api.NewEventHandler(eventService)

	server := &Server{
		router:   gin.Default(),
		eventAPI: eventHandler, // Assigning the handler to eventAPI
	}

	// Set up routes using the api.Context interface
	server.setupRoutes()

	return server
}

func (s *Server) setupRoutes() {
	// Schedule a new event
	s.router.POST("/events", func(c *gin.Context) {
		s.eventAPI.ScheduleEvent(api.NewGinContextAdapter(c))
	})

	// Retrieve a list of all scheduled events
	s.router.GET("/events", func(c *gin.Context) {
		s.eventAPI.GetAllEvents(api.NewGinContextAdapter(c))
	})

	// Retrieve details of a specific event
	s.router.GET("/events/:id", func(c *gin.Context) {
		s.eventAPI.GetEvent(api.NewGinContextAdapter(c))
	})

	// Update details of a specific event
	s.router.PUT("/events/:id", func(c *gin.Context) {
		s.eventAPI.UpdateEvent(api.NewGinContextAdapter(c))
	})

	// Delete a specific event
	s.router.DELETE("/events/:id", func(c *gin.Context) {
		s.eventAPI.DeleteEvent(api.NewGinContextAdapter(c))
	})

	// Additional routes can be added here following the same pattern
}

// Run starts the HTTP server
func (s *Server) Run(port string) error {
	return s.router.Run(port)
}
