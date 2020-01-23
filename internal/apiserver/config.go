package apiserver

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config struct for app config
type Config struct {
	Server struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Timeout struct {
			Write int `yaml:"write"`
			Read  int `yaml:"read"`
			Idle  int `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
	Logging struct {
		Level string `yaml:"level"`
	} `yaml:"logging"`
	Static struct {
		Path string `yaml:"path"`
	} `yaml:"static"`
}

// NewConfig returns a new Config struct
func NewConfig(configPath string) *Config {
	// Define config structure
	c := &Config{}

	// Open config file
	f, err := os.Open(filepath.Clean(configPath))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(f)

	// Start YAML decoding from file
	if err := d.Decode(&c); err != nil {
		log.Fatal(err)
	}

	// Default return
	return c
}
