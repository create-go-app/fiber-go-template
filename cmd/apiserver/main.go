package main

import (
	"log"

	"github.com/create-go-app/fiber-go-template/internal/apiserver"
)

func main() {
	// Generate our config based on the config supplied
	// by the user in the flags
	configPath, err := apiserver.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	// Create new config
	config, err := apiserver.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Create new server
	server := apiserver.NewServer(config)

	// Start server
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
