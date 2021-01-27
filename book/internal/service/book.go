// Package service contains all service logic.
package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/ivyoverflow/pub-sub/book/internal/lib/types"
	"github.com/ivyoverflow/pub-sub/book/internal/model"
	"github.com/ivyoverflow/pub-sub/book/internal/repository"
)

// Book implements all service methods for book.
type Book struct {
	repo repository.BookI
	gen  IDGeneratorI
}

// NewBook returns a new configured Book object.
func NewBook(repo repository.BookI, gen IDGeneratorI) *Book {
	return &Book{repo, gen}
}

// Insert calls Insert repository method.
func (s *Book) Insert(ctx context.Context, book *model.Book) (*model.Book, error) {
	if err := Validate(book); err != nil {
		return nil, types.ErrorBadRequest
	}

	book.ID = s.gen.Generate()

	return s.repo.Insert(ctx, book)
}

// Get calls Get repository method.
func (s *Book) Get(ctx context.Context, bookID uuid.UUID) (*model.Book, error) {
	return s.repo.Get(ctx, bookID)
}

// Update calls Update repository method.
func (s *Book) Update(ctx context.Context, bookID uuid.UUID, book *model.Book) (*model.Book, error) {
	if err := Validate(book); err != nil {
		return nil, types.ErrorBadRequest
	}

	return s.repo.Update(ctx, bookID, book)
}

// Delete calls Delete repository method.
func (s *Book) Delete(ctx context.Context, bookID uuid.UUID) (*model.Book, error) {
	return s.repo.Delete(ctx, bookID)
}
