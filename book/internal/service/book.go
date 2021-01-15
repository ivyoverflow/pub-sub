// Package service contains all service logic.
package service

import (
	"github.com/ivyoverflow/pub-sub/book/internal/model"
	"github.com/ivyoverflow/pub-sub/book/internal/store"
)

// Book implements all service methods for book.
type Book struct {
	str *store.Store
}

// NewBook returns a new configured Book object.
func NewBook(str *store.Store) *Book {
	return &Book{str}
}

// Add calls Add repository method.
func (service *Book) Add(book *model.Book) (*model.Book, error) {
	book.ID = GenerateUniqueID()

	return service.str.Book.Add(book)
}

// Get calls Get repository method.
func (service *Book) Get(bookID string) (*model.Book, error) {
	return service.str.Book.Get(bookID)
}

// Update calls Update repository method.
func (service *Book) Update(bookID string, book *model.Book) (*model.Book, error) {
	return service.str.Book.Update(bookID, book)
}

// Delete calls Delete repository method.
func (service *Book) Delete(bookID string) (*model.Book, error) {
	return service.str.Book.Delete(bookID)
}
