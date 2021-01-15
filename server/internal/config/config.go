// Package config contains the logic to initialize the application config.
package config

import (
	"os"
)

// Config contains Addr and Port fields that will be used to configure server.
type Config struct {
	Addr string
	Port string
}

// New returrns a new configured Config object.
func New() *Config {
	return &Config{
		Addr: os.Getenv("ADDR"),
		Port: os.Getenv("PORT"),
	}
}
