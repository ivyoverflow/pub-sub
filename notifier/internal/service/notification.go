// Package service cotains a Notifier inteface and its service implementation.
package service

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/ivyoverflow/pub-sub/notifier/internal/repository"
)

// Notification implements all service methods for Notifier.
type Notification struct {
	repo repository.Notifier
}

// NewNotification returns a new configured Notification object.
func NewNotification(repo repository.Notifier) *Notification {
	return &Notification{repo}
}

// Publish calls a Publish repository method.
func (s *Notification) Publish(ctx context.Context, book string, message interface{}) error {
	return s.repo.Publish(ctx, book, message)
}

// Subscribe calls a Subscribe repository method.
func (s *Notification) Subscribe(ctx context.Context, book string) <-chan *redis.Message {
	return s.repo.Subscribe(ctx, book)
}
