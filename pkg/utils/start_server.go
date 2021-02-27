package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutdown! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Define server address, like `host:port` (from config).
	serverAddress := os.Getenv("SERVER_HOST") + os.Getenv("SERVER_PORT")

	// Run server.
	if err := a.Listen(serverAddress); err != nil {
		log.Fatal(err)
	}

	<-idleConnsClosed
}
