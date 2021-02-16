// Package repository contains Notifier interface and its repository implementation.
package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Notifier describes all methods that will be used for notifications.
type Notifier interface {
	Publish(ctx context.Context, book string, message interface{}) error
	Subscribe(ctx context.Context, book string) <-chan *redis.Message
}
