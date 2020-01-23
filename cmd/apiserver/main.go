package main

import (
	"flag"
	"log"

	"github.com/create-go-app/fiber-go-template/internal/apiserver"
)

var (
	// Path to config file
	configPath string
)

func init() {
	// Looking for flags
	flag.StringVar(&configPath, "config-path", "configs/apiserver.yml", "path to config file")
}

func main() {
	// Parse flags
	flag.Parse()

	// Init app config
	config := apiserver.NewConfig(configPath)

	// Define new server
	server := apiserver.New(config)

	// Start new server
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
