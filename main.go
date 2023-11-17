package main

import (
	"homeTestAlfaBet/internal/scheduler"
	"homeTestAlfaBet/internal/server"
)

func main() {
	// Initialize and start the scheduler
	scheduler.SetupEventReminders()

	// Initialize the server with all dependencies
	srv := server.NewServer()

	// Start the server
	if err := srv.Run(":8080"); err != nil {
		panic(err)
	}
}
