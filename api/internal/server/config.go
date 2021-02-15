// Package server implements server logic: routes initialization and server configuration.
package server

import (
	"fmt"
	"os"
	"sync"
)

// Config contains addr and port fields that will be used to configure the server.
type Config struct {
	Addr string `envconfig:"ADDR" default:"localhost"`
	Port string `envconfig:"PORT" default:"8080"`
}

// NewConfig returrns a new configured ServerConfig object.
func NewConfig() *Config {
	var config Config
	var once sync.Once
	once.Do(func() {
		config = Config{
			Addr: os.Getenv("ADDR"),
			Port: os.Getenv("PORT"),
		}
	})

	return &config
}

// GetConnectionURI returns the formatted connection URI.
func (cfg *Config) GetConnectionURI() string {
	return fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port)
}
