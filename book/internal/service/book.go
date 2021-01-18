// Package service contains all service logic.
package service

import (
	"context"

	"github.com/ivyoverflow/pub-sub/book/internal/model"
	"github.com/ivyoverflow/pub-sub/book/internal/repository"
)

// Book implements all service methods for book.
type Book struct {
	repo repository.BookI
}

// NewBook returns a new configured Book object.
func NewBook(repo repository.BookI) *Book {
	return &Book{repo}
}

// Insert calls Insert repository method.
func (svc *Book) Insert(ctx context.Context, book *model.Book) (*model.Book, error) {
	book.ID = GenerateUniqueID()

	return svc.repo.Insert(ctx, book)
}

// Get calls Get repository method.
func (svc *Book) Get(ctx context.Context, bookID string) (*model.Book, error) {
	return svc.repo.Get(ctx, bookID)
}

// Update calls Update repository method.
func (svc *Book) Update(ctx context.Context, bookID string, book *model.Book) (*model.Book, error) {
	return svc.repo.Update(ctx, bookID, book)
}

// Delete calls Delete repository method.
func (svc *Book) Delete(ctx context.Context, bookID string) (*model.Book, error) {
	return svc.repo.Delete(ctx, bookID)
}
