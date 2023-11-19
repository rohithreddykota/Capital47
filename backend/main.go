// main.go

package main

import (
	"Capital47/pkg/api"
	"fmt"
	"log"
	"net/http"

	"Capital47/pkg/config"
	"Capital47/pkg/nessie"
	"github.com/gorilla/mux"
)

func main() {
	// Load configuration from the config file
	loadedConfig, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %s\n", err)
	}

	// Print the loaded configuration
	config.PrintConfig(loadedConfig)

	// Initialize the Nessie API client
	nessieAPI := nessie.NewNessieAPI(loadedConfig.Nessie.BaseURL, loadedConfig.Nessie.APIKey)

	// Set the initialized NessieAPI
	nessie.SetNessieAPI(nessieAPI)

	// Create a new instance of the server
	server := NewServer(loadedConfig.Server.Port, nessieAPI)

	// Start the server
	log.Printf("Server is starting on port %d...\n", loadedConfig.Server.Port)
	err = server.Start()
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

// NewServer creates a new instance of the server.
func NewServer(port int, nessieAPI *nessie.NessieAPI) *Server {
	router := api.NewRouter(nessieAPI)
	return &Server{
		Router: router,
		Port:   port,
	}
}

// Server represents the HTTP server.
type Server struct {
	Router *mux.Router
	Port   int
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.Port)
	return http.ListenAndServe(addr, s.Router)
}