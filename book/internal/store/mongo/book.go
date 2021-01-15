// Package mongo contains MongoDB repository implementation.
package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ivyoverflow/pub-sub/book/internal/model"
)

// BookRepository implements all MongoDB repository methods for BookRepository.
type BookRepository struct {
	clt *Client
}

// NewBookRepository returns a new configured BookRepository object.
func NewBookRepository(clt *Client) *BookRepository {
	return &BookRepository{clt}
}

// Add adds a new book to the books collection.
func (repository *BookRepository) Add(book *model.Book) (*model.Book, error) {
	bookDB := repository.clt.Database("bookdb")
	bookCollection := bookDB.Collection("books")
	if _, err := bookCollection.InsertOne(context.Background(), book); err != nil {
		return nil, err
	}

	return nil, nil
}

// Get receives a book from the books collection by bookID.
func (repository *BookRepository) Get(bookID int) (*model.Book, error) {
	bookDB := repository.clt.Database("bookdb")
	bookCollection := bookDB.Collection("books")
	book := model.Book{}
	filter := bson.D{primitive.E{Key: "id", Value: bookID}}
	if err := bookCollection.FindOne(context.Background(), filter).Decode(&book); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		return nil, err
	}

	return nil, nil
}

// Update updates a book from the books collection by book ID.
func (repository *BookRepository) Update(bookID int, book *model.Book) (*model.Book, error) {
	bookDB := repository.clt.Database("bookdb")
	bookCollection := bookDB.Collection("books")
	filter := bson.D{primitive.E{Key: "id", Value: bookID}}
	if _, err := bookCollection.UpdateOne(context.Background(), filter, book); err != nil {
		return nil, err
	}

	return nil, nil
}

// Delete deletes a book from the books table by book ID.
func (repository *BookRepository) Delete(bookID int) (*model.Book, error) {
	bookDB := repository.clt.Database("bookdb")
	bookCollection := bookDB.Collection("books")
	filter := bson.D{primitive.E{Key: "id", Value: bookID}}
	if _, err := bookCollection.DeleteOne(context.Background(), filter); err != nil {
		return nil, err
	}

	return nil, nil
}
