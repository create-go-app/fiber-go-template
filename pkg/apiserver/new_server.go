package apiserver

import (
	"os"

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
		// Add logger
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Next:       nil,
			Format:     "${time} [${status}] ${method} ${path} (${latency})\n",
			TimeFormat: "Mon, 2 Jan 2006 15:04:05 MST",
			Output:     os.Stdout,
		}),
	)

	// Static files
	if s.config.Static.Prefix != "" {
		app.Static(s.config.Static.Prefix, s.config.Static.Path)
	} else {
		app.Static("/", s.config.Static.Path)
	}

	// Register API routes
	Routes(app)

	return app
}
