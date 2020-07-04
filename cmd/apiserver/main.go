package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/create-go-app/fiber-go-template/pkg/apiserver"
	"github.com/gofiber/fiber"
)

func main() {
	// Parse config path from environment variable
	configPath := apiserver.GetEnv("CONFIG_PATH", "configs/apiserver.yml")

	// Create new config
	config, err := apiserver.NewConfig(configPath)
	apiserver.ErrChecker(err)

	// Create channels
	quit := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Catch OS signals
	signal.Notify(quit, os.Interrupt)

	// Create new server
	server := apiserver.NewServer(config).Start()

	// Create gracefull shutdown goroutine
	go func(server *fiber.App, quit <-chan os.Signal, done chan<- bool) {
		<-quit

		fmt.Println("API server is shutting down...")
		if err := server.Shutdown(); err != nil {
			fmt.Println("API server shutdown failed!")
			return
		}

		close(done)
	}(server, quit, done)

	// Start server
	apiserver.ErrChecker(
		server.Listen(config.Server.Host + ":" + config.Server.Port),
	)

	<-done
}
