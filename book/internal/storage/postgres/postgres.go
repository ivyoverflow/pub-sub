// Package postgres contains PostgreSQL repository implementation.
package postgres

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/ivyoverflow/pub-sub/book/internal/lib/types"
)

// DB represents a PostgreSQL database.
type DB struct {
	*sqlx.DB
}

// New connects to the PostgreSQL database and returns a new sqlx.DB object or an error.
func New(ctx context.Context) (*DB, error) {
	cfg := NewConfig()
	db, err := sqlx.Open("postgres", cfg.GetConnectionURI())
	if err != nil {
		log.Println(err.Error())

		return nil, types.ErrorPostgresConnectionRefused
	}

	if err := db.PingContext(ctx); err != nil {
		log.Println(err.Error())

		return nil, types.ErrorPostgresConnectionRefused
	}

	return &DB{db}, nil
}
