package postgres_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/book/internal/model"
	"github.com/ivyoverflow/pub-sub/book/internal/repository/postgres"
	"github.com/ivyoverflow/pub-sub/book/internal/service"
)

func TestBookRepository_Insert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Mock initialization throws an error: %v", err)
	}

	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := postgres.NewBookRepository(&postgres.DB{sqlxDB})

	testCases := []struct {
		name     string
		input    model.Book
		mock     func(book *model.Book)
		expected *model.Book
	}{
		{
			name: "OK",
			input: model.Book{
				ID:          "974a3de439ed5c6e",
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `
				Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming
				language makes working with concurrency tractable and even easy. If you’re a developer familiar with Go,
				this practical book demonstrates best practices and patterns to help you incorporate concurrency into your systems.
				Author Katherine Cox-Buday takes you step-by-step through the process.
				You’ll understand how Go chooses to model concurrency, what issues arise from this model,
				and how you can compose primitives within this model to solve problems.
				Learn the skills and tooling you need to confidently write and implement concurrent systems of any size.`,
				Rating:  71.00,
				Price:   36.90,
				InStock: true,
			},
			mock: func(book *model.Book) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "name", "date_of_issue", "author", "description", "rating", "price", "in_stock"}).
					AddRow(book.ID, book.Name, book.DateOfIssue, book.Author, book.Description, book.Rating, book.Price, book.InStock)
				mock.ExpectQuery("INSERT INTO books").WithArgs(book.ID, book.Name, book.DateOfIssue, book.Author,
					book.Description, book.Rating, book.Price, book.InStock).WillReturnRows(rows)

				mock.ExpectCommit()
			},
			expected: &model.Book{
				ID:          "974a3de439ed5c6e",
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `
				Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming
				language makes working with concurrency tractable and even easy. If you’re a developer familiar with Go,
				this practical book demonstrates best practices and patterns to help you incorporate concurrency into your systems.
				Author Katherine Cox-Buday takes you step-by-step through the process.
				You’ll understand how Go chooses to model concurrency, what issues arise from this model,
				and how you can compose primitives within this model to solve problems.
				Learn the skills and tooling you need to confidently write and implement concurrent systems of any size.`,
				Rating:  71.00,
				Price:   36.90,
				InStock: true,
			},
		},
		{
			name: "Empty fields",
			input: model.Book{
				ID:          service.GenerateUniqueID(),
				Name:        "",
				DateOfIssue: "",
				Author:      "Katherine Cox-Buday",
				Description: `
				Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming
				language makes working with concurrency tractable and even easy. If you’re a developer familiar with Go,
				this practical book demonstrates best practices and patterns to help you incorporate concurrency into your systems.
				Author Katherine Cox-Buday takes you step-by-step through the process.
				You’ll understand how Go chooses to model concurrency, what issues arise from this model,
				and how you can compose primitives within this model to solve problems.
				Learn the skills and tooling you need to confidently write and implement concurrent systems of any size.`,
				Rating:  71.00,
				Price:   36.90,
				InStock: true,
			},
			mock: func(book *model.Book) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "name", "date_of_issue", "author", "description", "rating", "price", "in_stock"}).
					AddRow(book.ID, book.Name, book.DateOfIssue, book.Author, book.Description, book.Rating, book.Price, book.InStock).
					RowError(0, errors.New("insert error"))
				mock.ExpectQuery("INSERT INTO books").WithArgs(book.ID, book.Name, book.DateOfIssue, book.Author,
					book.Description, book.Rating, book.Price, book.InStock).WillReturnRows(rows)

				mock.ExpectRollback()
			},
			expected: nil,
		},
	}

	for _, testCase := range testCases {
		testCase.mock(&testCase.input)

		receivedBook, err := repo.Insert(context.Background(), &testCase.input)
		if err != nil {
			t.Errorf("Insert method throws an error: %v", err)
		}

		assert.Equal(t, testCase.expected, receivedBook)
	}
}

func TestBookRepository_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Mock initialization throws an error: %v", err)
	}

	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := postgres.NewBookRepository(&postgres.DB{sqlxDB})

	testCases := []struct {
		name     string
		input    model.Book
		mock     func(book *model.Book)
		expected *model.Book
	}{
		{
			name: "OK",
			input: model.Book{
				ID:          "974a3de439ed5c6e",
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `
				Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming
				language makes working with concurrency tractable and even easy. If you’re a developer familiar with Go,
				this practical book demonstrates best practices and patterns to help you incorporate concurrency into your systems.
				Author Katherine Cox-Buday takes you step-by-step through the process.
				You’ll understand how Go chooses to model concurrency, what issues arise from this model,
				and how you can compose primitives within this model to solve problems.
				Learn the skills and tooling you need to confidently write and implement concurrent systems of any size.`,
				Rating:  71.00,
				Price:   36.90,
				InStock: true,
			},
			mock: func(book *model.Book) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "name", "date_of_issue", "author", "description", "rating", "price", "in_stock"}).
					AddRow(book.ID, book.Name, book.DateOfIssue, book.Author, book.Description, book.Rating, book.Price, book.InStock)
				mock.ExpectQuery("SELECT (.+) FROM books WHERE (.+)").WithArgs(book.ID).WillReturnRows(rows)

				mock.ExpectCommit()
			},
			expected: &model.Book{
				ID:          "974a3de439ed5c6e",
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `
				Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming
				language makes working with concurrency tractable and even easy. If you’re a developer familiar with Go,
				this practical book demonstrates best practices and patterns to help you incorporate concurrency into your systems.
				Author Katherine Cox-Buday takes you step-by-step through the process.
				You’ll understand how Go chooses to model concurrency, what issues arise from this model,
				and how you can compose primitives within this model to solve problems.
				Learn the skills and tooling you need to confidently write and implement concurrent systems of any size.`,
				Rating:  71.00,
				Price:   36.90,
				InStock: true,
			},
		},
		{
			name: "Not found",
			input: model.Book{
				ID: "dawjdi12i3jdhwaj",
			},
			mock: func(book *model.Book) {
				mock.ExpectBegin()

				mock.ExpectQuery("SELECT (.+) FROM books WHERE (.+)").WithArgs(book.ID)

				mock.ExpectRollback()
			},
			expected: nil,
		},
	}

	for _, testCase := range testCases {
		testCase.mock(&testCase.input)

		receivedBook, err := repo.Get(context.Background(), testCase.input.ID)
		if err != nil {
			t.Errorf("Get method throws an error: %v", err)
		}

		assert.Equal(t, testCase.expected, receivedBook)
	}
}

func TestBookRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Mock initialization throws an error: %v", err)
	}

	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := postgres.NewBookRepository(&postgres.DB{sqlxDB})

	testCases := []struct {
		name     string
		input    model.Book
		mock     func(book *model.Book)
		expected *model.Book
	}{
		{
			name: "OK",
			input: model.Book{
				ID:          "974a3de439ed5c6e",
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming
				language makes working with concurrency tractable and even easy. If you’re a developer familiar with Go,
				this practical book demonstrates best practices and patterns to help you incorporate concurrency into your systems.
				Author Katherine Cox-Buday takes you step-by-step through the process.
				You’ll understand how Go chooses to model concurrency, what issues arise from this model,
				and how you can compose primitives within this model to solve problems.
				Learn the skills and tooling you need to confidently write and implement concurrent systems of any size.`,
				Rating:  71.00,
				Price:   36.90,
				InStock: true,
			},
			mock: func(book *model.Book) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "name", "date_of_issue", "author", "description", "rating", "price", "in_stock"}).
					AddRow(book.ID, book.Name, book.DateOfIssue, book.Author, book.Description, book.Rating, book.Price, book.InStock)

				mock.ExpectQuery("UPDATE books SET (.+), (.+), (.+), (.+), (.+), (.+), (.+) WHERE (.+)").
					WithArgs(book.Name, book.DateOfIssue, book.Author, book.Description,
						book.Rating, book.Price, book.InStock, book.ID).WillReturnRows(rows)

				mock.ExpectCommit()
			},
			expected: &model.Book{
				ID:          "974a3de439ed5c6e",
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming
				language makes working with concurrency tractable and even easy. If you’re a developer familiar with Go,
				this practical book demonstrates best practices and patterns to help you incorporate concurrency into your systems.
				Author Katherine Cox-Buday takes you step-by-step through the process.
				You’ll understand how Go chooses to model concurrency, what issues arise from this model,
				and how you can compose primitives within this model to solve problems.
				Learn the skills and tooling you need to confidently write and implement concurrent systems of any size.`,
				Rating:  71.00,
				Price:   36.90,
				InStock: true,
			},
		},
		{
			name: "Not found",
			input: model.Book{
				ID: "dawjdi12i3jdhwaj",
			},
			mock: func(book *model.Book) {
				mock.ExpectBegin()

				mock.ExpectQuery("UPDATE books SET (.+), (.+), (.+), (.+), (.+), (.+), (.+) WHERE (.+)").
					WithArgs(book.Name, book.DateOfIssue, book.Author, book.Description,
						book.Rating, book.Price, book.InStock, book.ID)

				mock.ExpectRollback()
			},
			expected: nil,
		},
	}

	for _, testCase := range testCases {
		testCase.mock(&testCase.input)

		updatedBook, err := repo.Update(context.Background(), testCase.input.ID, &testCase.input)
		if err != nil {
			t.Errorf("Update method throws an error: %v", err)
		}

		assert.Equal(t, testCase.expected, updatedBook)
	}
}

func TestBookRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Mock initialization throws an error: %v", err)
	}

	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := postgres.NewBookRepository(&postgres.DB{sqlxDB})

	testCases := []struct {
		name     string
		input    model.Book
		mock     func(book *model.Book)
		expected *model.Book
	}{
		{
			name: "OK",
			input: model.Book{
				ID:          "974a3de439ed5c6e",
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming
				language makes working with concurrency tractable and even easy. If you’re a developer familiar with Go,
				this practical book demonstrates best practices and patterns to help you incorporate concurrency into your systems.
				Author Katherine Cox-Buday takes you step-by-step through the process.
				You’ll understand how Go chooses to model concurrency, what issues arise from this model,
				and how you can compose primitives within this model to solve problems.
				Learn the skills and tooling you need to confidently write and implement concurrent systems of any size.`,
				Rating:  71.00,
				Price:   36.90,
				InStock: true,
			},
			mock: func(book *model.Book) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "name", "date_of_issue", "author", "description", "rating", "price", "in_stock"}).
					AddRow(book.ID, book.Name, book.DateOfIssue, book.Author, book.Description, book.Rating, book.Price, book.InStock)

				mock.ExpectQuery("DELETE FROM books WHERE (.+)").WithArgs(book.ID).WillReturnRows(rows)

				mock.ExpectCommit()
			},
			expected: &model.Book{
				ID:          "974a3de439ed5c6e",
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming
				language makes working with concurrency tractable and even easy. If you’re a developer familiar with Go,
				this practical book demonstrates best practices and patterns to help you incorporate concurrency into your systems.
				Author Katherine Cox-Buday takes you step-by-step through the process.
				You’ll understand how Go chooses to model concurrency, what issues arise from this model,
				and how you can compose primitives within this model to solve problems.
				Learn the skills and tooling you need to confidently write and implement concurrent systems of any size.`,
				Rating:  71.00,
				Price:   36.90,
				InStock: true,
			},
		},
		{
			name: "Not found",
			input: model.Book{
				ID: "dawjdi12i3jdhwaj",
			},
			mock: func(book *model.Book) {
				mock.ExpectBegin()

				mock.ExpectQuery("DELETE FROM books WHERE (.+)").WithArgs(book.ID)

				mock.ExpectRollback()
			},
			expected: nil,
		},
	}

	for _, testCase := range testCases {
		testCase.mock(&testCase.input)

		deletedBook, err := repo.Delete(context.Background(), testCase.input.ID)
		if err != nil {
			t.Errorf("Delete method throws an error: %v", err)
		}

		assert.Equal(t, testCase.expected, deletedBook)
	}
}
