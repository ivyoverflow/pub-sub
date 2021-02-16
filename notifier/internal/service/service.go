// Package service cotains a Notifier inteface and its service implementation.
package service

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Notifier describes all service methods that will be used for notifications.
type Notifier interface {
	Publish(ctx context.Context, book string, message interface{}) error
	Subscribe(ctx context.Context, book string) <-chan *redis.Message
}
