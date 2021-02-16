// Package redis contains Redis client initialization.
package redis

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/ivyoverflow/pub-sub/notifier/internal/config"
)

// DB represents a Redis database.
type DB struct {
	*redis.Client
}

// NewDB connects to the Redis database and returns a Redis client.
func NewDB(ctx context.Context, cfg *config.RedisConfig) (*DB, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := db.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
