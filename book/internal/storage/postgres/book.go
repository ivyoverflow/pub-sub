// Package postgres contains PostgreSQL repository implementation.
package postgres

import (
	"context"
	"database/sql"
	"strings"

	"github.com/google/uuid"

	"github.com/ivyoverflow/pub-sub/book/internal/lib/types"
	"github.com/ivyoverflow/pub-sub/book/internal/model"
)

// BookRepository implements all PostgreSQL repository methods for BookRepository.
type BookRepository struct {
	pg *DB
}

// NewBookRepository returns a new configured BookRepository object.
func NewBookRepository(pg *DB) *BookRepository {
	return &BookRepository{pg}
}

// Insert adds a new book to the books table.
func (r *BookRepository) Insert(ctx context.Context, book *model.Book) (*model.Book, error) {
	insertedBook := model.Book{}
	query := `INSERT INTO books (id, name, date_of_issue, author, description, rating, price, in_stock)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *`
	row := r.pg.QueryRowContext(ctx, query, book.ID, book.Name, book.DateOfIssue, book.Author,
		book.Description, book.Rating, book.Price, book.InStock)
	if err := row.Scan(&insertedBook.ID, &insertedBook.Name, &insertedBook.DateOfIssue, &insertedBook.Author,
		&insertedBook.Description, &insertedBook.Rating, &insertedBook.Price, &insertedBook.InStock); err != nil {
		switch {
		case strings.Contains(err.Error(), "unique constraint"):
			return nil, types.ErrorDuplicateValue
		default:
			return nil, err
		}
	}

	return &insertedBook, nil
}

// Get receives a book from the books table by bookID.
func (r *BookRepository) Get(ctx context.Context, bookID uuid.UUID) (*model.Book, error) {
	book := model.Book{}
	query := "SELECT * FROM books WHERE id = $1"
	row := r.pg.QueryRowContext(ctx, query, bookID)
	if err := row.Scan(&book.ID, &book.Name, &book.DateOfIssue, &book.Author, &book.Description,
		&book.Rating, &book.Price, &book.InStock); err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, types.ErrorNotFound
		default:
			return nil, err
		}
	}

	return &book, nil
}

// Update updates a book from the books table by book ID.
func (r *BookRepository) Update(ctx context.Context, bookID uuid.UUID, book *model.Book) (*model.Book, error) {
	updatedBook := model.Book{}
	query := `UPDATE books SET name = $1, date_of_issue = $2, author = $3, description = $4, rating = $5, price = $6, in_stock = $7
	WHERE id = $8 RETURNING *`
	row := r.pg.QueryRowContext(ctx, query, book.Name, book.DateOfIssue, book.Author, book.Description, book.Rating, book.Price, book.InStock, bookID)
	if err := row.Scan(&updatedBook.ID, &updatedBook.Name, &updatedBook.DateOfIssue, &updatedBook.Author, &updatedBook.Description,
		&updatedBook.Rating, &updatedBook.Price, &updatedBook.InStock); err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, types.ErrorNotFound
		case strings.Contains(err.Error(), "unique constraint"):
			return nil, types.ErrorDuplicateValue
		default:
			return nil, err
		}
	}

	return &updatedBook, nil
}

// Delete deletes a book from the books table by book ID.
func (r *BookRepository) Delete(ctx context.Context, bookID uuid.UUID) (*model.Book, error) {
	deletedBook := model.Book{}
	query := "DELETE FROM books WHERE id = $1 RETURNING *"
	row := r.pg.QueryRowContext(ctx, query, bookID)
	if err := row.Scan(&deletedBook.ID, &deletedBook.Name, &deletedBook.DateOfIssue, &deletedBook.Author, &deletedBook.Description,
		&deletedBook.Rating, &deletedBook.Price, &deletedBook.InStock); err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, types.ErrorNotFound
		default:
			return nil, err
		}
	}

	return &deletedBook, nil
}
