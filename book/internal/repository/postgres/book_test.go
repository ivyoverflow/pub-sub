package postgres_test

import (
	"testing"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/constants"
	"github.com/ivyoverflow/pub-sub/book/internal/repository"
	"github.com/ivyoverflow/pub-sub/book/internal/repository/postgres"
)

func newPostgresTestConfig() *config.PostgresConfig {
	return &config.PostgresConfig{
		Host:           constants.PostgresHost,
		Port:           constants.PostgresPort,
		Name:           constants.PostgresName,
		User:           constants.PostgresUser,
		Password:       constants.PostgresPassword,
		MigartionsPath: constants.PostgresMigartionsPath,
		SSLMode:        constants.PostgresSSLMode,
	}
}

func clearDB(db *postgres.DB) error {
	return db.QueryRow("DELETE FROM books").Err()
}

func TestPostgresBookRepository(t *testing.T) {
	cfg := newPostgresTestConfig()
	db, err := postgres.New(cfg)
	if err != nil {
		t.Errorf("Postgres connection throws an error: %v", err)
	}

	if err := clearDB(db); err != nil {
		t.Errorf("ClearDB function throws an error: %v", err)
	}

	repo := postgres.NewBookRepository(db)
	suite := repository.NewSuite(repo)
	suite.Run(t)
}
