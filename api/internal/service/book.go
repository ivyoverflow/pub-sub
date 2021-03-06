// Package service contains all service logic.
package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/ivyoverflow/pub-sub/api/internal/lib/types"
	"github.com/ivyoverflow/pub-sub/api/internal/model"
	"github.com/ivyoverflow/pub-sub/api/internal/storage"
)

// BookController implements all service methods for book.
type BookController struct {
	repo storage.Booker
	gen  Generator
}

// NewBookController returns a new configured BookController object.
func NewBookController(repo storage.Booker, gen Generator) *BookController {
	return &BookController{repo, gen}
}

// Insert calls Insert repository method.
func (s *BookController) Insert(ctx context.Context, book *model.Book) (*model.Book, error) {
	if err := Validate(book); err != nil {
		return nil, types.ErrorValidation
	}

	book.ID = s.gen.GenerateUUID()

	return s.repo.Insert(ctx, book)
}

// Get calls Get repository method.
func (s *BookController) Get(ctx context.Context, bookID uuid.UUID) (*model.Book, error) {
	return s.repo.Get(ctx, bookID)
}

// Update calls Update repository method.
func (s *BookController) Update(ctx context.Context, bookID uuid.UUID, book *model.Book) (*model.Book, error) {
	if err := Validate(book); err != nil {
		return nil, types.ErrorValidation
	}

	return s.repo.Update(ctx, bookID, book)
}

// Delete calls Delete repository method.
func (s *BookController) Delete(ctx context.Context, bookID uuid.UUID) (*model.Book, error) {
	return s.repo.Delete(ctx, bookID)
}
