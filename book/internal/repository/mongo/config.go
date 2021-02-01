// Package mongo contains MongoDB repository implementation.
package mongo

import (
	"fmt"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// Config contains fields that will be used to configure MongoDB connection.
type Config struct {
	Host     string `envconfig:"MONGOHOST" default:"localhost"`
	Port     string `envconfig:"MONGOPORT" default:"27017"`
	User     string `envconfig:"MONGONAME" default:"admin"`
	Name     string `envconfig:"MONGOUSER" default:"admin"`
	Password string `envconfig:"MONGOPASSWORD" default:"qwerty"`
}

// NewConfig returrns a new configured Config object.
func NewConfig() *Config {
	var config Config
	var once sync.Once
	once.Do(func() {
		if err := envconfig.Process("", &config); err != nil {
			return
		}
	})

	return &config
}

// GetConnectionURI returns the formatted Mongo URI.
func (cfg *Config) GetConnectionURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
}
