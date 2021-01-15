// Package store contains repository interfaces and implementations.
package store

import (
	"log"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/store/postgres"
)

// Store is a struct that contains all repository implementations.
type Store struct {
	Book BookRepository
}

// New returns a new configured Store object.
func New(cfg *config.Config) (*Store, error) {
	pg, err := postgres.Dial(cfg)
	if err != nil {
		return nil, err
	}

	if pg != nil {
		if err := postgres.RunMigration(cfg); err != nil {
			log.Println("Welcome!")
			return nil, err
		}
	}

	return &Store{
		Book: postgres.NewBookRepository(pg),
	}, nil
}
