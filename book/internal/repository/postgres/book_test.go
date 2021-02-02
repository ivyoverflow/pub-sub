package postgres_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/constants"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/types"
	"github.com/ivyoverflow/pub-sub/book/internal/model"
	"github.com/ivyoverflow/pub-sub/book/internal/repository/postgres"
)

func newPostgresTestConfig() *config.PostgresConfig {
	return &config.PostgresConfig{
		Host:           constants.PostgresHost,
		Port:           constants.PostgresPort,
		Name:           constants.PostgresName,
		User:           constants.PostgresUser,
		Password:       constants.PostgresPassword,
		MigartionsPath: constants.PostgresMigartionsPath,
		SSLMode:        constants.PostgresSSLMode,
	}
}

func clearDB(db *postgres.DB) error {
	return db.QueryRow("DELETE FROM books").Err()
}

func TestPostgresBookRepository_Insert(t *testing.T) {
	testCases := []struct {
		name          string
		input         model.Book
		expected      *model.Book
		expectedError error
	}{
		{
			name: "OK",
			input: model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
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
				Rating:  model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:   model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock: true,
			},
			expected: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
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
				Rating:  model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:   model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock: true,
			},
			expectedError: nil,
		},
		{
			name: "OK",
			input: model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
				Name:        "Introducing Go: Build Reliable, Scalable Programs",
				DateOfIssue: "2016",
				Author:      "Caleb Doxsey",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(45.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(129.24)},
				InStock:     true,
			},
			expected: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
				Name:        "Introducing Go: Build Reliable, Scalable Programs",
				DateOfIssue: "2016",
				Author:      "Caleb Doxsey",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(45.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(129.24)},
				InStock:     true,
			},
			expectedError: nil,
		},
		{
			name: "Duplicate value",
			input: model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
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
				Rating:  model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:   model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock: true,
			},
			expected:      nil,
			expectedError: types.ErrorDuplicateValue,
		},
	}

	ctx := context.Background()
	cfg := newPostgresTestConfig()
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
			assert.Equal(t, testCase.expectedError, err)
		}

		assert.Equal(t, testCase.expected, receivedBook)
	}
}

func TestPostgresBookRepository_Get(t *testing.T) {
	testCases := []struct {
		name          string
		input         uuid.UUID
		expected      *model.Book
		expectedError error
	}{
		{
			name:  "OK",
			input: uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
			expected: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
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
				Rating:  model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:   model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock: true,
			},
			expectedError: nil,
		},
		{
			name:          "Book not found",
			input:         uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120005"),
			expected:      nil,
			expectedError: types.ErrorNotFound,
		},
	}

	ctx := context.Background()
	cfg := newPostgresTestConfig()
	db, err := postgres.New(cfg)
	if err != nil {
		t.Errorf("Postgres connection throws an error: %v", err)
	}

	repo := postgres.NewBookRepository(db)
	for _, testCase := range testCases {
		receivedBook, err := repo.Get(ctx, testCase.input)
		if err != nil {
			assert.Equal(t, testCase.expectedError, err)
		}

		assert.Equal(t, testCase.expected, receivedBook)
	}
}

func TestPostgresBookRepository_Update(t *testing.T) {
	testCases := []struct {
		name          string
		input         uuid.UUID
		toUpdate      model.Book
		expected      *model.Book
		expectedError error
	}{
		{
			name:  "OK",
			input: uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
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
				Rating:  model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:   model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock: true,
			},
			expected: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
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
				Rating:  model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:   model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock: true,
			},
			expectedError: nil,
		},
		{
			name:  "Duplicate value",
			input: uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
			toUpdate: model.Book{
				Name:        "Concurrency in Go: TTD",
				DateOfIssue: "2016",
				Author:      "Caleb Doxsey",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(45.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(129.24)},
				InStock:     true,
			},
			expected:      nil,
			expectedError: types.ErrorDuplicateValue,
		},
		{
			name:  "Book not found",
			input: uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120005"),
			toUpdate: model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
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
				Rating:  model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:   model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock: true,
			},
			expected:      nil,
			expectedError: types.ErrorNotFound,
		},
	}

	ctx := context.Background()
	cfg := newPostgresTestConfig()
	db, err := postgres.New(cfg)
	if err != nil {
		t.Errorf("Postgres connection throws an error: %v", err)
	}

	repo := postgres.NewBookRepository(db)
	for _, testCase := range testCases {
		updatedBook, err := repo.Update(ctx, testCase.input, &testCase.toUpdate)
		if err != nil {
			assert.Equal(t, testCase.expectedError, err)
		}

		assert.Equal(t, testCase.expected, updatedBook)
	}
}

func TestPostgresBookRepository_Delete(t *testing.T) {
	testCases := []struct {
		name          string
		input         uuid.UUID
		expected      *model.Book
		expectedError error
	}{
		{
			name:  "OK",
			input: uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
			expected: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
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
				Rating:  model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:   model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock: true,
			},
			expectedError: nil,
		},
		{
			name:          "Book not found",
			input:         uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120005"),
			expected:      nil,
			expectedError: types.ErrorNotFound,
		},
	}

	ctx := context.Background()
	cfg := newPostgresTestConfig()
	db, err := postgres.New(cfg)
	if err != nil {
		t.Errorf("Postgres connection throws an error: %v", err)
	}

	repo := postgres.NewBookRepository(db)
	for _, testCase := range testCases {
		deletedBook, err := repo.Delete(ctx, testCase.input)
		if err != nil {
			assert.Equal(t, testCase.expectedError, err)
		}

		assert.Equal(t, testCase.expected, deletedBook)
	}
}
