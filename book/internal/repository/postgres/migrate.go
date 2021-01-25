// Package postgres contains PostgreSQL repository implementation.
package postgres

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Should be imported for run migrations.
	_ "github.com/golang-migrate/migrate/v4/source/file"       // Should be imported for run migrations.

	"github.com/ivyoverflow/pub-sub/book/internal/config"
)

// runMigration automatically starts PostgreSQL migrations.
func runMigration(cfg *config.Config) error {
	if cfg.Postgres.MigartionsPath == "" {
		return nil
	}

	if cfg.Postgres.Host == "" || cfg.Postgres.Name == "" || cfg.Postgres.Port == "" ||
		cfg.Postgres.User == "" || cfg.Postgres.Password == "" {
		return errors.New("postgreSQL URL is incorrect")
	}

	mrt, err := migrate.New(cfg.Postgres.MigartionsPath, cfg.Postgres.GetPostgresConnectionURI())
	if err != nil {
		return err
	}

	if err = mrt.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
