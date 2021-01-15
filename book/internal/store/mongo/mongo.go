// Package mongo contains MongoDB repository implementation.
package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
)

// Client represents a MongoDB client.
type Client struct {
	*mongo.Client
}

// Dial connects to the MongoDB database and returns a new Client connection or an error.
func Dial(cfg *config.Config) (*Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", cfg.Mongo.User, cfg.Mongo.Password, cfg.Mongo.Host, cfg.Mongo.Port, cfg.Mongo.Name)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Connect(context.Background()); err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, err
	}

	return &Client{client}, nil
}
