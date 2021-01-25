package mongo_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/constants"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/types"
	"github.com/ivyoverflow/pub-sub/book/internal/model"
	"github.com/ivyoverflow/pub-sub/book/internal/repository/mongo"
)

var (
	ctx = context.Background()
	cfg = newMongoTestConfig()
)

func newMongoTestConfig() *config.MongoConfig {
	return &config.MongoConfig{
		Host:     constants.MongoHost,
		Port:     constants.MongoPort,
		Name:     constants.MongoName,
		User:     constants.MongoUser,
		Password: constants.MongoPassword,
	}
}

func clearDB(db *mongo.DB) error {
	if _, err := db.Collection("books").DeleteMany(ctx, bson.M{}); err != nil {
		return err
	}

	return nil
}

func TestMongoBookRepository_Insert(t *testing.T) {
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
				Rating:  decimal.NewFromFloat(71.00),
				Price:   decimal.NewFromFloat(36.90),
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
				Rating:  decimal.NewFromFloat(71.00),
				Price:   decimal.NewFromFloat(36.90),
				InStock: true,
			},
			expectedError: nil,
		},
		{
			name: "OK",
			input: model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
				Name:        "Testing in Go",
				DateOfIssue: "2020",
				Author:      "Unknown",
				Description: `...`,
				Rating:      decimal.NewFromFloat(71.00),
				Price:       decimal.NewFromFloat(36.90),
				InStock:     true,
			},
			expected: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
				Name:        "Testing in Go",
				DateOfIssue: "2020",
				Author:      "Unknown",
				Description: `...`,
				Rating:      decimal.NewFromFloat(71.00),
				Price:       decimal.NewFromFloat(36.90),
				InStock:     true,
			},
			expectedError: nil,
		},
		{
			name: "Duplicate value",
			input: model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120004"),
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
				Rating:  decimal.NewFromFloat(71.00),
				Price:   decimal.NewFromFloat(36.90),
				InStock: true,
			},
			expected:      nil,
			expectedError: types.ErrorDuplicateValue,
		},
	}

	db, err := mongo.New(ctx, cfg)
	if err != nil {
		t.Errorf("Mongo connection throws an error: %v", err)
	}

	if err := clearDB(db); err != nil {
		t.Errorf("ClearDB function throws an error: %v", err)
	}

	repo := mongo.NewBookRepository(db)
	for _, testCase := range testCases {
		receivedBook, err := repo.Insert(ctx, &testCase.input)
		if err != nil {
			assert.Equal(t, testCase.expectedError, err)
		}

		assert.Equal(t, testCase.expected, receivedBook)
	}
}

func TestMongoBookRepository_Get(t *testing.T) {
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

	db, err := mongo.New(ctx, cfg)
	if err != nil {
		t.Errorf("Mongo connection throws an error: %v", err)
	}

	repo := mongo.NewBookRepository(db)
	for _, testCase := range testCases {
		receivedBook, err := repo.Get(ctx, testCase.input)
		if err != nil {
			assert.Equal(t, testCase.expectedError, err)
		}

		assert.Equal(t, testCase.expected, receivedBook)
	}
}

func TestMongoBookRepository_Update(t *testing.T) {
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
				Rating:  decimal.NewFromFloat(71.00),
				Price:   decimal.NewFromFloat(36.90),
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
				Rating:  decimal.NewFromFloat(71.00),
				Price:   decimal.NewFromFloat(36.90),
				InStock: true,
			},
			expectedError: nil,
		},
		{
			name:  "Duplicate value",
			input: uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
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
				Rating:  decimal.NewFromFloat(71.00),
				Price:   decimal.NewFromFloat(36.90),
				InStock: true,
			},
			expected:      nil,
			expectedError: types.ErrorDuplicateValue,
		},
		{
			name:          "Book not found",
			input:         uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120005"),
			expected:      nil,
			expectedError: types.ErrorNotFound,
		},
	}

	db, err := mongo.New(ctx, cfg)
	if err != nil {
		t.Errorf("Mongo connection throws an error: %v", err)
	}

	repo := mongo.NewBookRepository(db)
	for _, testCase := range testCases {
		updatedBook, err := repo.Update(ctx, testCase.input, &testCase.toUpdate)
		if err != nil {
			assert.Equal(t, testCase.expectedError, err)
		}

		assert.Equal(t, testCase.expected, updatedBook)
	}
}

func TestMongoBookRepository_Delete(t *testing.T) {
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
				Rating:  decimal.NewFromFloat(71.00),
				Price:   decimal.NewFromFloat(36.90),
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

	db, err := mongo.New(ctx, cfg)
	if err != nil {
		t.Errorf("Mongo connection throws an error: %v", err)
	}

	repo := mongo.NewBookRepository(db)
	for _, testCase := range testCases {
		deletedBook, err := repo.Delete(ctx, testCase.input)
		if err != nil {
			assert.Equal(t, testCase.expectedError, err)
		}

		assert.Equal(t, testCase.expected, deletedBook)
	}
}
