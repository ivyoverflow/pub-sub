// Package config contains the logic to initialize the application config.
package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// ServerConfig contains fields that will be used to configure server.
type ServerConfig struct {
	Addr string `envconfig:"ADDR" default:"localhost"`
	Port string `envconfig:"PORT" default:"8081"`
}

// NewServer returrns a new configured ServerConfig object.
func NewServer() *ServerConfig {
	var config ServerConfig
	var once sync.Once
	once.Do(func() {
		if err := envconfig.Process("", &config); err != nil {
			return
		}
	})

	return &config
}

// RedisConfig contains fields that will be used to configure Redis client.
type RedisConfig struct {
	Addr     string `envconfig:"REDISADDR" default:"localhost:6379"`
	Password string `envconfig:"REDISPASSWORD" default:""`
	DB       int    `envconfig:"REDISDB" default:"0"`
}

// NewRedis returns a new configured RedisConfig object.
func NewRedis() *RedisConfig {
	var config RedisConfig
	var once sync.Once
	once.Do(func() {
		if err := envconfig.Process("", &config); err != nil {
			return
		}
	})

	return &config
}
