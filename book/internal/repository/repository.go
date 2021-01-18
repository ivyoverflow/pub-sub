// Package repository contains repository interfaces and implementations.
package repository

import (
	"context"

	"github.com/ivyoverflow/pub-sub/book/internal/model"
)

// BookI describes all repository methods for book.
type BookI interface {
	Insert(ctx context.Context, book *model.Book) (*model.Book, error)
	Get(ctx context.Context, bookID string) (*model.Book, error)
	Update(ctx context.Context, bookID string, book *model.Book) (*model.Book, error)
	Delete(ctx context.Context, bookID string) (*model.Book, error)
}
