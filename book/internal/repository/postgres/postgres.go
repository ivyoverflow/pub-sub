// Package postgres contains PostgreSQL repository implementation.
package postgres

import (
	"github.com/jmoiron/sqlx"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
)

// DB represents a PostgreSQL database.
type DB struct {
	*sqlx.DB
}

// New connects to the PostgreSQL database and returns a new sqlx.DB object or an error.
func New(cfg *config.Config) (*DB, error) {
	db, err := sqlx.Open("postgres", cfg.Postgres.GetPostgresConnectionURI())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if db != nil {
		if err := runMigration(cfg); err != nil {
			return nil, err
		}
	}

	return &DB{db}, nil
}
