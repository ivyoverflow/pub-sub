// Package mongo contains MongoDB repository implementation.
package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/ivyoverflow/pub-sub/book/internal/model"
)

// BookRepository ...
type BookRepository struct {
	db *DB
}

// NewBookRepository ...
func NewBookRepository(db *DB) *BookRepository {
	return &BookRepository{db}
}

// Insert ...
func (repo *BookRepository) Insert(ctx context.Context, book *model.Book) (*model.Book, error) {
	if _, err := repo.db.Collection("books").InsertOne(ctx, book); err != nil {
		return nil, err
	}

	return repo.Get(ctx, book.ID)
}

// Get ...
func (repo *BookRepository) Get(ctx context.Context, bookID string) (*model.Book, error) {
	filter := bson.D{{Key: "id", Value: bookID}}
	receivedBook := model.Book{}
	if err := repo.db.Collection("books").FindOne(ctx, filter).Decode(&receivedBook); err != nil {
		return nil, err
	}

	return &receivedBook, nil
}

// Update ...
func (repo *BookRepository) Update(ctx context.Context, bookID string, book *model.Book) (*model.Book, error) {
	filter := bson.D{{Key: "id", Value: bookID}}
	fieldsToUpdate := bson.M{"$set": bson.M{"name": book.Name, "dateOfIssue": book.DateOfIssue, "author": book.Author,
		"description": book.Description, "rating": book.Rating, "price": book.Price, "inStock": book.InStock}}
	updatedBook := model.Book{}
	if err := repo.db.Collection("books").FindOneAndUpdate(ctx, filter, fieldsToUpdate).Decode(&updatedBook); err != nil {
		return nil, err
	}

	return &updatedBook, nil
}

// Delete ...
func (repo *BookRepository) Delete(ctx context.Context, bookID string) (*model.Book, error) {
	filter := bson.D{{Key: "id", Value: bookID}}
	deletedBook := model.Book{}
	if err := repo.db.Collection("books").FindOneAndDelete(ctx, filter).Decode(&deletedBook); err != nil {
		return nil, err
	}

	return &deletedBook, nil
}
