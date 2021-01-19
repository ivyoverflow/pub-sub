// Package mongo contains MongoDB repository implementation.
package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ivyoverflow/pub-sub/book/internal/model"
)

// BookRepository implements all MongoDB repository methods for BookRepository.
type BookRepository struct {
	db *DB
}

// NewBookRepository returns a new configured BookRepository object.
func NewBookRepository(db *DB) *BookRepository {
	return &BookRepository{db}
}

// Insert adds a new book to the books collection.
func (repo *BookRepository) Insert(ctx context.Context, book *model.Book) (*model.Book, error) {
	if _, err := repo.db.Collection("books").InsertOne(ctx, book); err != nil {
		return nil, err
	}

	time.Sleep(10 * time.Second)

	return repo.Get(ctx, book.ID)
}

// Get receives a book from the books collection by bookID.
func (repo *BookRepository) Get(ctx context.Context, bookID string) (*model.Book, error) {
	filter := bson.D{{Key: "id", Value: bookID}}
	receivedBook := model.Book{}
	if err := repo.db.Collection("books").FindOne(ctx, filter).Decode(&receivedBook); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &receivedBook, nil
}

// Update updates a book from the books collection by book ID.
func (repo *BookRepository) Update(ctx context.Context, bookID string, book *model.Book) (*model.Book, error) {
	filter := bson.D{{Key: "id", Value: bookID}}
	fieldsToUpdate := bson.M{"$set": bson.M{"name": book.Name, "dateOfIssue": book.DateOfIssue, "author": book.Author,
		"description": book.Description, "rating": book.Rating, "price": book.Price, "inStock": book.InStock}}
	updatedBook := model.Book{}
	if err := repo.db.Collection("books").FindOneAndUpdate(ctx, filter, fieldsToUpdate).Decode(&updatedBook); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &updatedBook, nil
}

// Delete deletes a book from the books collection by book ID.
func (repo *BookRepository) Delete(ctx context.Context, bookID string) (*model.Book, error) {
	filter := bson.D{{Key: "id", Value: bookID}}
	deletedBook := model.Book{}
	if err := repo.db.Collection("books").FindOneAndDelete(ctx, filter).Decode(&deletedBook); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &deletedBook, nil
}
