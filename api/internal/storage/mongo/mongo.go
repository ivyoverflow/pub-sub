// Package mongo contains MongoDB repository implementation.
package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/ivyoverflow/pub-sub/api/internal/lib/types"
)

// DB represents a MongoDB database.
type DB struct {
	*mongo.Database
}

// New connects to the MongoDB database and returns a new mongo.Database object or an error.
func New(ctx context.Context) (*DB, error) {
	cfg := NewConfig()
	clt, err := mongo.NewClient(options.Client().ApplyURI(cfg.GetConnectionURI()))
	if err != nil {
		return nil, types.ErrorMongoConnectionRefused
	}

	if err := clt.Connect(ctx); err != nil {
		return nil, types.ErrorMongoConnectionRefused
	}

	if err := clt.Ping(ctx, readpref.Primary()); err != nil {
		return nil, types.ErrorMongoConnectionRefused
	}

	db := clt.Database(cfg.Name)

	if err := RunMigration(ctx, db); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
