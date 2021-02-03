package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ivyoverflow/pub-sub/book/internal/lib/types"
)

// RunMigration creates unique indexes.
func RunMigration(ctx context.Context, db *mongo.Database) error {
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}

	if _, err := db.Collection("books").Indexes().CreateMany(ctx, indexes); err != nil {
		log.Println(err.Error())

		return types.ErrorMigrate
	}

	return nil
}
