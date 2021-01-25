// Package mongo contains MongoDB repository implementation.
package mongo

import (
	"context"
	"log"
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

// Insert adds a new book to the books collection.
func (r *BookRepository) Insert(ctx context.Context, book *model.Book) (*model.Book, error) {
	ratingBytes, err := book.Rating.MarshalBinary()
	if err != nil {
		return nil, types.ErrorInternalServerError
	}

	priceBytes, err := book.Price.MarshalBinary()
	if err != nil {
		return nil, types.ErrorInternalServerError
	}

	fieldsToInsert := bson.M{"id": book.ID, "name": book.Name, "dateOfIssue": book.DateOfIssue, "author": book.Author,
		"description": book.Description, "rating": ratingBytes, "price": priceBytes, "inStock": book.InStock}
	if _, err := r.db.Collection("books").InsertOne(ctx, fieldsToInsert); err != nil {
		switch {
		case strings.Contains(err.Error(), "E11000"):
			return nil, types.ErrorDuplicateValue
		default:
			return nil, types.ErrorInternalServerError
		}
	}

	return r.Get(ctx, book.ID)
}

// Get receives a book from the books collection by bookID.
func (r *BookRepository) Get(ctx context.Context, bookID uuid.UUID) (*model.Book, error) {
	filter := bson.D{{Key: "id", Value: bookID}}
	receivedBook := model.Book{}
	if err := r.db.Collection("books").FindOne(ctx, filter).Decode(&receivedBook); err != nil {
		switch {
		case err == mongo.ErrNoDocuments:
			return nil, types.ErrorNotFound
		default:
			log.Println(err.Error())
			return nil, types.ErrorInternalServerError
		}
	}

	return &receivedBook, nil
}

// Update updates a book from the books collection by book ID.
func (r *BookRepository) Update(ctx context.Context, bookID uuid.UUID, book *model.Book) (*model.Book, error) {
	filter := bson.D{{Key: "id", Value: bookID.String()}}
	fieldsToUpdate := bson.M{"$set": bson.M{"name": book.Name, "dateOfIssue": book.DateOfIssue, "author": book.Author,
		"description": book.Description, "rating": book.Rating, "price": book.Price, "inStock": book.InStock}}
	updatedBook := model.Book{}
	if err := r.db.Collection("books").FindOneAndUpdate(ctx, filter, fieldsToUpdate).Decode(&updatedBook); err != nil {
		switch {
		case err == mongo.ErrNoDocuments:
			return nil, types.ErrorNotFound
		case strings.Contains(err.Error(), "E11000"):
			return nil, types.ErrorDuplicateValue
		default:
			return nil, types.ErrorInternalServerError
		}
	}

	return r.Get(ctx, bookID)
}

// Delete deletes a book from the books collection by book ID.
func (r *BookRepository) Delete(ctx context.Context, bookID uuid.UUID) (*model.Book, error) {
	filter := bson.D{{Key: "id", Value: bookID.String()}}
	deletedBook := model.Book{}
	if err := r.db.Collection("books").FindOneAndDelete(ctx, filter).Decode(&deletedBook); err != nil {
		switch {
		case err == mongo.ErrNoDocuments:
			return nil, types.ErrorNotFound
		default:
			return nil, types.ErrorNotFound
		}
	}

	return &deletedBook, nil
}
