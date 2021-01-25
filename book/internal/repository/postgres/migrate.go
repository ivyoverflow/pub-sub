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
func runMigration(cfg *config.PostgresConfig) error {
	if cfg.MigartionsPath == "" {
		return nil
	}

	if cfg.Host == "" || cfg.Name == "" || cfg.Port == "" ||
		cfg.User == "" || cfg.Password == "" {
		return errors.New("postgreSQL URL is incorrect")
	}

	mrt, err := migrate.New(cfg.MigartionsPath, cfg.GetPostgresConnectionURI())
	if err != nil {
		return err
	}

	if err = mrt.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
