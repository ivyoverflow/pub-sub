// Package storage contains storage interfaces and implementations.
package storage

import (
	"context"

	"github.com/google/uuid"

	"github.com/ivyoverflow/pub-sub/book/internal/model"
)

// BookI describes all repository methods for book.
type BookI interface {
	Insert(ctx context.Context, book *model.Book) (*model.Book, error)
	Get(ctx context.Context, bookID uuid.UUID) (*model.Book, error)
	Update(ctx context.Context, bookID uuid.UUID, book *model.Book) (*model.Book, error)
	Delete(ctx context.Context, bookID uuid.UUID) (*model.Book, error)
}
