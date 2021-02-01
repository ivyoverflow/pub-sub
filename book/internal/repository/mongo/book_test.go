package mongo_test

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/constants"
	"github.com/ivyoverflow/pub-sub/book/internal/repository"
	"github.com/ivyoverflow/pub-sub/book/internal/repository/mongo"
)

func newMongoTestConfig() *config.MongoConfig {
	return &config.MongoConfig{
		Host:     constants.MongoHost,
		Port:     constants.MongoPort,
		Name:     constants.MongoName,
		User:     constants.MongoUser,
		Password: constants.MongoPassword,
	}
}

func clearDB(db *mongo.DB) error {
	if _, err := db.Collection("books").DeleteMany(context.Background(), bson.M{}); err != nil {
		return err
	}

	return nil
}

func TestMongoBookRepository(t *testing.T) {
	ctx := context.Background()
	cfg := newMongoTestConfig()
	db, err := mongo.New(ctx, cfg)
	if err != nil {
		t.Errorf("Mongo connection throws an error: %v", err)
	}

	if err := clearDB(db); err != nil {
		t.Errorf("ClearDB function throws an error: %v", err)
	}

	repo := mongo.NewBookRepository(db)
	suite := repository.NewSuite(repo)
	suite.Run(t)
}
