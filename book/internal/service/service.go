// Package service contains all service logic.
package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/ivyoverflow/pub-sub/book/internal/model"
)

// BookI describes all service methods for book.
type BookI interface {
	Insert(ctx context.Context, book *model.Book) (*model.Book, error)
	Get(ctx context.Context, bookID uuid.UUID) (*model.Book, error)
	Update(ctx context.Context, bookID uuid.UUID, book *model.Book) (*model.Book, error)
	Delete(ctx context.Context, bookID uuid.UUID) (*model.Book, error)
}

// IDGeneratorI ...
type IDGeneratorI interface {
	Generate() uuid.UUID
}
