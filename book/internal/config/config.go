// Package config contains the logic to initialize the application config.
package config

import (
	"os"
	"sync"
)

// MongoConfig contains fields that will be used to configure MongoDB connection.
type MongoConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

// PostgresConfig contains fields that will be used to configure PostgreSQL connection.
type PostgresConfig struct {
	Host           string
	Port           string
	User           string
	Name           string
	Password       string
	MigartionsPath string
	SSLMode        string
}

// ServerConfig contains addr and port fields that will be used to configure the server.
type ServerConfig struct {
	Addr string
	Port string
}

// Config contains all configs.
type Config struct {
	Mongo    MongoConfig
	Postgres PostgresConfig
	Server   ServerConfig
}

// New returrns a new configured Config object.
func New() *Config {
	var config Config
	var once sync.Once
	once.Do(func() {
		config = Config{
			Mongo: MongoConfig{
				Host:     os.Getenv("MONGOHOST"),
				Port:     os.Getenv("MONGOPORT"),
				User:     os.Getenv("MONGOUSER"),
				Name:     os.Getenv("MONGONAME"),
				Password: os.Getenv("MONGOPASSWORD"),
			},
			Postgres: PostgresConfig{
				Host:           os.Getenv("PGHOST"),
				Port:           os.Getenv("PGPORT"),
				User:           os.Getenv("PGUSER"),
				Name:           os.Getenv("PGNAME"),
				Password:       os.Getenv("PGPASSWORD"),
				MigartionsPath: os.Getenv("PGMIGRATIONSPATH"),
				SSLMode:        os.Getenv("PGSSLMODE"),
			},
			Server: ServerConfig{
				Addr: os.Getenv("ADDR"),
				Port: os.Getenv("PORT"),
			},
		}
	})

	return &config
}
