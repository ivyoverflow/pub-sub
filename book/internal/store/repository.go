// Package store contains repository interfaces and implementations.
package store

import (
	"github.com/ivyoverflow/pub-sub/book/internal/model"
)

// BookRepository describes all repository methods for book.
type BookRepository interface {
	Add(book *model.Book) (*model.Book, error)
	Get(bookID string) (*model.Book, error)
	Update(bookID string, book *model.Book) (*model.Book, error)
	Delete(bookID string) (*model.Book, error)
}
