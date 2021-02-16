// Package repository contains Notifier interface and its repository implementation.
package repository

import (
	"context"

	driver "github.com/go-redis/redis/v8"

	"github.com/ivyoverflow/pub-sub/notifier/internal/redis"
)

// Notification implements all Redis repository methods for Notifier.
type Notification struct {
	db *redis.DB
}

// NewNotification returns a new configured Notification object.
func NewNotification(db *redis.DB) *Notification {
	return &Notification{db}
}

// Publish makes Publish query to the Redis database.
func (r *Notification) Publish(ctx context.Context, book string, message interface{}) error {
	return r.db.Publish(ctx, book, message).Err()
}

// Subscribe makes a subscription request to the Redis database and returns a channel
// that will be used to receive notifications.
func (r *Notification) Subscribe(ctx context.Context, book string) <-chan *driver.Message {
	ps := r.db.Subscribe(ctx, book)

	return ps.Channel()
}
