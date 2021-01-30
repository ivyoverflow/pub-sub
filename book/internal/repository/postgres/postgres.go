// Package postgres contains PostgreSQL repository implementation.
package postgres

import (
	"github.com/jmoiron/sqlx"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/types"
)

// DB represents a PostgreSQL database.
type DB struct {
	*sqlx.DB
}

// New connects to the PostgreSQL database and returns a new sqlx.DB object or an error.
func New(cfg *config.PostgresConfig) (*DB, error) {
	db, err := sqlx.Open("postgres", cfg.GetPostgresConnectionURI())
	if err != nil {
		return nil, types.ErrorPostgresConnectionRefused
	}

	if err := db.Ping(); err != nil {
		return nil, types.ErrorPostgresConnectionRefused
	}

	if err := RunMigration(cfg); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
