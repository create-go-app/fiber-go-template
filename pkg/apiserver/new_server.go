package apiserver

import (
	"os"

	"github.com/gofiber/fiber/v2"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
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
		// Simple logger
		logger.New(
			logger.Config{
				Format:     "${time} [${status}] ${method} ${path} (${latency})\n",
				TimeFormat: "Mon, 2 Jan 2006 15:04:05 MST",
				Output:     os.Stdout,
			},
		),
	)

	// Add static files, if prefix and path was defined in config
	if s.config.Static.Prefix != "" && s.config.Static.Path != "" {
		app.Static(s.config.Static.Prefix, s.config.Static.Path)
	}

	// Register API routes
	Routes(app)

	return app
}
