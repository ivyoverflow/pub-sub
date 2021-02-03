// Package mongo contains MongoDB repository implementation.
package mongo

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ivyoverflow/pub-sub/book/internal/lib/types"
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

// Collection returns a MongoDB collection.
func (r *BookRepository) Collection() *mongo.Collection {
	return r.db.Collection("books")
}

// Insert adds a new book to the books collection.
func (r *BookRepository) Insert(ctx context.Context, book *model.Book) (*model.Book, error) {
	_, err := r.Collection().InsertOne(ctx, book)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "E11000"):
			return nil, types.ErrorDuplicateValue
		default:
			return nil, err
		}
	}

	return r.Get(ctx, book.ID)
}

// Get receives a book from the books collection by bookID.
func (r *BookRepository) Get(ctx context.Context, bookID uuid.UUID) (*model.Book, error) {
	filter := bson.D{{Key: "id", Value: bookID}}
	receivedBook := model.Book{}
	err := r.Collection().FindOne(ctx, filter).Decode(&receivedBook)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, types.ErrorNotFound
		default:
			return nil, err
		}
	}

	return &receivedBook, nil
}

// Update updates a book from the books collection by book ID.
func (r *BookRepository) Update(ctx context.Context, bookID uuid.UUID, book *model.Book) (*model.Book, error) {
	filter := bson.D{{Key: "id", Value: bookID}}
	fieldsToUpdate := bson.M{"$set": bson.M{"name": book.Name, "dateOfIssue": book.DateOfIssue, "author": book.Author,
		"description": book.Description, "rating": book.Rating, "price": book.Price, "inStock": book.InStock}}
	updatedBook := model.Book{}
	err := r.Collection().FindOneAndUpdate(ctx, filter, fieldsToUpdate).Decode(&updatedBook)
	if err != nil {
		switch {
		case err == mongo.ErrNoDocuments:
			return nil, types.ErrorNotFound
		case strings.Contains(err.Error(), "E11000"):
			return nil, types.ErrorDuplicateValue
		default:
			return nil, err
		}
	}

	return r.Get(ctx, updatedBook.ID)
}

// Delete deletes a book from the books collection by book ID.
func (r *BookRepository) Delete(ctx context.Context, bookID uuid.UUID) (*model.Book, error) {
	filter := bson.D{{Key: "id", Value: bookID}}
	deletedBook := model.Book{}
	err := r.Collection().FindOneAndDelete(ctx, filter).Decode(&deletedBook)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, types.ErrorNotFound
		default:
			return nil, err
		}
	}

	return &deletedBook, nil
}
