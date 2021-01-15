// Package service contains all service logic.
package service

import (
	"errors"

	"github.com/ivyoverflow/pub-sub/book/internal/store"
)

// Manager is a struct that contains all service implementations.
type Manager struct {
	Book BookService
}

// NewManager returns a new configured Manager object.
func NewManager(str *store.Store) (*Manager, error) {
	if str == nil {
		return nil, errors.New("no store provided")
	}

	return &Manager{
		Book: NewBook(str),
	}, nil
}
