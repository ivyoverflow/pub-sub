package postgres_test

import (
	"context"
	"testing"

	_ "github.com/lib/pq"

	"github.com/ivyoverflow/pub-sub/api/internal/storage"
	"github.com/ivyoverflow/pub-sub/api/internal/storage/postgres"
)

func clearDB(db *postgres.DB) error {
	return db.QueryRow("DELETE FROM books").Err()
}

func TestPostgresBookRepository(t *testing.T) {
	ctx := context.Background()
	db, err := postgres.New(ctx)
	if err != nil {
		t.Errorf("Postgres connection throws an error: %v", err)
	}

	if err := clearDB(db); err != nil {
		t.Errorf("ClearDB function throws an error: %v", err)
	}

	repo := postgres.NewBookRepository(db)
	suite := storage.NewSuite(repo)
	suite.Run(t)
}
