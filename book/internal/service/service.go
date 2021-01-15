// Package service contains all service logic.
package service

import "github.com/ivyoverflow/pub-sub/book/internal/model"

// BookService describes all service methods for book.
type BookService interface {
	Add(book *model.Book) (*model.Book, error)
	Get(bookID string) (*model.Book, error)
	Update(bookID string, book *model.Book) (*model.Book, error)
	Delete(bookID string) (*model.Book, error)
}
