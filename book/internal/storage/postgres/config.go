// Package postgres contains PostgreSQL repository implementation.
package postgres

import (
	"fmt"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// Config contains fields that will be used to configure PostgreSQL connection.
type Config struct {
	Host     string `envconfig:"PGHOST" default:"localhost"`
	Port     string `envconfig:"PGPORT" default:"5432"`
	User     string `envconfig:"PGUSER" default:"postgres"`
	Name     string `envconfig:"PGNAME" default:"postgres"`
	Password string `envconfig:"PGPASSWORD" default:"qwerty"`
	SSLMode  string `envconfig:"PGSSLMODE" default:"disable"`
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

// GetConnectionURI returns the formatted Postgres URI.
func (cfg *Config) GetConnectionURI() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode)
}
