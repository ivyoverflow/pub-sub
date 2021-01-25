package mongo

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb" // Should be imported for run migrations.
	_ "github.com/golang-migrate/migrate/v4/source/file"      // Should be imported for run migrations.

	"github.com/ivyoverflow/pub-sub/book/internal/config"
)

func runMigration(cfg *config.MongoConfig) error {
	if cfg.MigartionsPath == "" {
		return nil
	}

	if cfg.Host == "" || cfg.Name == "" || cfg.Port == "" ||
		cfg.User == "" || cfg.Password == "" {
		return errors.New("mongoDB URL is incorrect")
	}

	mrt, err := migrate.New(cfg.MigartionsPath, cfg.GetMongoConnectionURI())
	if err != nil {
		return err
	}

	if err = mrt.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
