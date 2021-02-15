package mongo_test

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/ivyoverflow/pub-sub/book/internal/storage"
	"github.com/ivyoverflow/pub-sub/book/internal/storage/mongo"
)

func clearDB(db *mongo.DB) error {
	if _, err := db.Collection("books").DeleteMany(context.Background(), bson.M{}); err != nil {
		return err
	}

	return nil
}

func TestMongoBookRepository(t *testing.T) {
	ctx := context.Background()
	db, err := mongo.New(ctx)
	if err != nil {
		t.Errorf("Mongo connection throws an error: %v", err)
	}

	if err := clearDB(db); err != nil {
		t.Errorf("ClearDB function throws an error: %v", err)
	}

	repo := mongo.NewBookRepository(db)
	suite := storage.NewSuite(repo)
	suite.Run(t)
}
