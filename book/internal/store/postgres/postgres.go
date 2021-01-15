// Package postgres contains PostgreSQL repository implementation.
package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
)

// DB represents a PostgreSQL connection.
type DB struct {
	*sqlx.DB
}

// Dial connects to the PostgreSQL database and returns a new Client connection or an error.
func Dial(cfg *config.Config) (*DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Name, cfg.Postgres.Password, cfg.Postgres.SSLMode))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
