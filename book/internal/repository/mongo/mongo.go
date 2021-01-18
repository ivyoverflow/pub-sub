// Package mongo contains MongoDB repository implementation.
package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
)

// DB ...
type DB struct {
	*mongo.Database
}

// New ...
func New(ctx context.Context, cfg *config.Config) (*DB, error) {
	clt, err := mongo.NewClient(options.Client().ApplyURI(cfg.Mongo.GetMongoConnectionURI()))
	if err != nil {
		return nil, err
	}

	if err = clt.Connect(ctx); err != nil {
		return nil, err
	}

	if err = clt.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	db := clt.Database("bookdb")

	return &DB{db}, nil
}
