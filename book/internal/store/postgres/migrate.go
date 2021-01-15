// Package postgres contains PostgreSQL repository implementation.
package postgres

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Should be imprted for run migrations.
	_ "github.com/golang-migrate/migrate/v4/source/file"       // Should be imprted for run migrations.

	"github.com/ivyoverflow/pub-sub/book/internal/config"
)

// RunMigration automatically starts PostgreSQL migrations.
func RunMigration(cfg *config.Config) error {
	if cfg.Postgres.MigartionsPath == "" {
		return nil
	}

	if cfg.Postgres.Host == "" || cfg.Postgres.Name == "" || cfg.Postgres.Port == "" ||
		cfg.Postgres.User == "" || cfg.Postgres.Password == "" {
		return errors.New("postgreSQL URL is incorrect")
	}

	mrt, err := migrate.New(
		cfg.Postgres.MigartionsPath,
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host,
			cfg.Postgres.Port, cfg.Postgres.Name, cfg.Postgres.SSLMode))
	if err != nil {
		return err
	}

	if err = mrt.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
