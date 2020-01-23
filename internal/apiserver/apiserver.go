package apiserver

import (
	"time"

	"github.com/gofiber/fiber"
	"go.uber.org/zap"
)

// APIServer struct
type APIServer struct {
	config *Config
	logger *zap.Logger
}

// New method for init new server instance
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: Logger(config),
	}
}

// Start method for start new server
func (s *APIServer) Start() error {
	// Init new app
	app := fiber.New()

	// App config
	host := s.config.Server.Host + ":" + s.config.Server.Port
	app.Engine.ReadTimeout = time.Duration(s.config.Server.Timeout.Read) * time.Second
	app.Engine.WriteTimeout = time.Duration(s.config.Server.Timeout.Write) * time.Second
	app.Engine.IdleTimeout = time.Duration(s.config.Server.Timeout.Idle) * time.Second

	// Show Fiber logo on console for debug mode
	if s.config.Logging.Level != "debug" {
		app.Banner = false
	}

	app.Use(func(c *fiber.Ctx) {
		// Log each request
		s.logger.Info(
			"fetch URL",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
		)

		// Go to next middleware
		c.Next()
	})

	// App routes
	app.Get("/", IndexHandler)

	// Start server
	app.Listen(host)

	return nil
}
