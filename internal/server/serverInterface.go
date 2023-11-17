package server

// ServerInterface defines the interface for a server
type ServerInterface interface {
	Run(port string) error
}
