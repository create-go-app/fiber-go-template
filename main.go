package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/create-go-app/fiber-go-template/pkg/apiserver"
)

func main() {
	// Parse config path from environment variable.
	configPath := apiserver.GetEnv("CONFIG_PATH", "configs/apiserver.yml")

	// Create new config.
	config, err := apiserver.NewConfig(configPath)
	apiserver.ErrChecker(err)

	// Create new server.
	server := apiserver.NewServer(config).Start()

	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // catch OS signals
		<-sigint

		// We received an interrupt signal, shut down.
		if err := server.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("API server Shutdown: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Start API server.
	apiserver.ErrChecker(
		server.Listen(config.Server.Host + ":" + config.Server.Port),
	)

	<-idleConnsClosed
}
