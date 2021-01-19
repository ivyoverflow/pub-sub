package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/model"
	"github.com/ivyoverflow/pub-sub/book/internal/repository/postgres"
)

var (
	cfg = config.New()
	ctx = context.Background()
)

func clearDB(db *postgres.DB) error {
	return db.QueryRow("DELETE FROM books").Err()
}

func TestBookRepository_Insert(t *testing.T) {
	testCases := []struct {
		name     string
		input    model.Book
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
			name: "Duplicate index",
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
			expected: nil,
		},
		{
			name: "Duplicate book name",
			input: model.Book{
				ID:          "213dawdf31fka",
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
			expected: nil,
		},
	}

	db, err := postgres.New(cfg)
	if err != nil {
		t.Errorf("Postgres connection throws an error: %v", err)
	}

	if err := clearDB(db); err != nil {
		t.Errorf("ClearDB function throws an error: %v", err)
	}

	repo := postgres.NewBookRepository(db)
	for _, testCase := range testCases {
		receivedBook, err := repo.Insert(ctx, &testCase.input)
		if err != nil {
			t.Errorf("Insert method throws an error: %v", err)
		}

		assert.Equal(t, testCase.expected, receivedBook)
	}
}

func TestBookRepository_Get(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected *model.Book
	}{
		{
			name:  "OK",
			input: "974a3de439ed5c6e",
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
			name:     "Book not found",
			input:    "213dawdf31fka",
			expected: nil,
		},
	}

	db, err := postgres.New(cfg)
	if err != nil {
		t.Errorf("Postgres connection throws an error: %v", err)
	}

	repo := postgres.NewBookRepository(db)
	for _, testCase := range testCases {
		receivedBook, err := repo.Get(ctx, testCase.input)
		if err != nil {
			t.Errorf("Get method throws an error: %v", err)
		}

		assert.Equal(t, testCase.expected, receivedBook)
	}
}

func TestBookRepository_Update(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		toUpdate model.Book
		expected *model.Book
	}{
		{
			name:  "OK",
			input: "974a3de439ed5c6e",
			toUpdate: model.Book{
				Name:        "Concurrency in Go: TTD",
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
			expected: &model.Book{
				ID:          "974a3de439ed5c6e",
				Name:        "Concurrency in Go: TTD",
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
			name:     "Book not found",
			input:    "213dawdf31fka",
			expected: nil,
		},
	}

	db, err := postgres.New(cfg)
	if err != nil {
		t.Errorf("Postgres connection throws an error: %v", err)
	}

	repo := postgres.NewBookRepository(db)
	for _, testCase := range testCases {
		updatedBook, err := repo.Update(ctx, testCase.input, &testCase.toUpdate)
		if err != nil {
			t.Errorf("Update method throws an error: %v", err)
		}

		assert.Equal(t, testCase.expected, updatedBook)
	}
}

func TestBookRepository_Delete(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected *model.Book
	}{
		{
			name:  "OK",
			input: "974a3de439ed5c6e",
			expected: &model.Book{
				ID:          "974a3de439ed5c6e",
				Name:        "Concurrency in Go: TTD",
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
			name:     "Book not found",
			input:    "213dawdf31fka",
			expected: nil,
		},
	}

	db, err := postgres.New(cfg)
	if err != nil {
		t.Errorf("Postgres connection throws an error: %v", err)
	}

	repo := postgres.NewBookRepository(db)
	for _, testCase := range testCases {
		deletedBook, err := repo.Delete(ctx, testCase.input)
		if err != nil {
			t.Errorf("Delete method throws an error: %v", err)
		}

		assert.Equal(t, testCase.expected, deletedBook)
	}
}
