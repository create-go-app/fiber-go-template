package apiserver

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

// APIServer struct
type APIServer struct {
	config *Config
}

// NewServer method for init new server instance
func NewServer(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

// Start method for start new server
func (s *APIServer) Start() *fiber.App {
	// Initialize a new app
	app := fiber.New()

	// Register middlewares
	app.Use(
		cors.New(), // Add CORS to each route
		middleware.Logger("${time} [${status}] ${method} ${path} (${latency})\n"), // Simple logger
	)

	// Add static files, if prefix and path was defined in config
	if s.config.Static.Prefix != "" && s.config.Static.Path != "" {
		app.Static(s.config.Static.Prefix, s.config.Static.Path)
	}

	// Register API routes
	Routes(app)

	return app
}
