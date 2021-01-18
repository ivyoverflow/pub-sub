// Package postgres contains PostgreSQL repository implementation.
package postgres

import (
	"context"

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
func (repo *BookRepository) Insert(ctx context.Context, book *model.Book) (*model.Book, error) {
	tx, err := repo.pg.Begin()
	if err != nil {
		return nil, err
	}

	createdBook := model.Book{}
	query := `INSERT INTO books (id, name, date_of_issue, author, description, rating, price, in_stock)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *`
	row := tx.QueryRow(query, book.ID, book.Name, book.DateOfIssue, book.Author, book.Description, book.Rating, book.Price, book.InStock)
	if err = row.Scan(&createdBook.ID, &createdBook.Name, &createdBook.DateOfIssue, &createdBook.Author,
		&createdBook.Description, &createdBook.Rating, &createdBook.Price, &createdBook.InStock); err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}

		return nil, err
	}

	return &createdBook, tx.Commit()
}

// Get receives a book from the books table by bookID.
func (repo *BookRepository) Get(ctx context.Context, bookID string) (*model.Book, error) {
	tx, err := repo.pg.Begin()
	if err != nil {
		return nil, err
	}

	book := model.Book{}
	query := "SELECT * FROM books WHERE id = $1"
	row := tx.QueryRow(query, bookID)
	if err = row.Scan(&book.ID, &book.Name, &book.DateOfIssue, &book.Author, &book.Description,
		&book.Rating, &book.Price, &book.InStock); err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}

		return nil, err
	}

	return &book, tx.Commit()
}

// Update updates a book from the books table by book ID.
func (repo *BookRepository) Update(ctx context.Context, bookID string, book *model.Book) (*model.Book, error) {
	tx, err := repo.pg.Begin()
	if err != nil {
		return nil, err
	}

	updatedBook := model.Book{}
	query := `UPDATE books SET name = $1, date_of_issue = $2, author = $3, description = $4, rating = $5, price = $6, in_stock = $7
	WHERE id = $8 RETURNING *`
	row := tx.QueryRow(query, book.Name, book.DateOfIssue, book.Author, book.Description, book.Rating, book.Price, book.InStock, bookID)
	if err = row.Scan(&updatedBook.ID, &updatedBook.Name, &updatedBook.DateOfIssue, &updatedBook.Author, &updatedBook.Description,
		&updatedBook.Rating, &updatedBook.Price, &updatedBook.InStock); err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}

		return nil, err
	}

	return &updatedBook, tx.Commit()
}

// Delete deletes a book from the books table by book ID.
func (repo *BookRepository) Delete(ctx context.Context, bookID string) (*model.Book, error) {
	tx, err := repo.pg.Begin()
	if err != nil {
		return nil, err
	}

	deletedBook := model.Book{}
	query := "DELETE FROM books WHERE id = $1 RETURNING *"
	row := tx.QueryRow(query, bookID)
	if err = row.Scan(&deletedBook.ID, &deletedBook.Name, &deletedBook.DateOfIssue, &deletedBook.Author, &deletedBook.Description,
		&deletedBook.Rating, &deletedBook.Price, &deletedBook.InStock); err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}

		return nil, err
	}

	return &deletedBook, tx.Commit()
}
