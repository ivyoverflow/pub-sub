// Package postgres contains PostgreSQL repository implementation.
package postgres

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Should be imported for run migrations.
	_ "github.com/golang-migrate/migrate/v4/source/file"       // Should be imported for run migrations.

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/types"
)

// RunMigration automatically starts PostgreSQL migrations.
func RunMigration(cfg *config.PostgresConfig) error {
	if cfg.MigartionsPath == "" {
		return types.ErrorMigrate
	}

	if cfg.Host == "" || cfg.Name == "" || cfg.Port == "" ||
		cfg.User == "" || cfg.Password == "" {
		return types.ErrorMigrate
	}

	mrt, err := migrate.New(cfg.MigartionsPath, cfg.GetPostgresConnectionURI())
	if err != nil {
		return types.ErrorMigrate
	}

	if err = mrt.Up(); err != nil && err != migrate.ErrNoChange {
		return types.ErrorMigrate
	}

	return nil
}
