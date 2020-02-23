package apiserver

import (
	"github.com/gofiber/fiber"
	"go.uber.org/zap"
)

// APIServer struct
type APIServer struct {
	config *Config
	logger *zap.Logger
}

// NewServer method for init new server instance
func NewServer(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: Logger(config),
	}
}

// Start method for start new server
func (s *APIServer) Start() error {
	// Init new app
	app := fiber.New()

	// App host config
	host := s.config.Server.Host + ":" + s.config.Server.Port

	// Static files
	if s.config.Static.Prefix != "" {
		app.Static(s.config.Static.Prefix, s.config.Static.Path)
	} else {
		app.Static(s.config.Static.Path)
	}

	// Middlewares
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
	if err := app.Listen(host); err != nil {
		s.logger.Info(
			"error",
			zap.Error(err),
		)

		return err
	}

	return nil
}
