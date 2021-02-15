// Package storage contains repository interfaces and implementations.
package storage

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/api/internal/lib/types"
	"github.com/ivyoverflow/pub-sub/api/internal/model"
)

// Suite contains all repository tests.
type Suite struct {
	repo Booker
}

// NewSuite returns a new configured Suite object.
func NewSuite(repo Booker) *Suite {
	return &Suite{repo}
}

// Run starts all repository tests.
func (s *Suite) Run(t *testing.T) {
	s.testInsert(t)
	s.testGet(t)
	s.testUpdate(t)
	s.testDelete(t)
}

func (s *Suite) testInsert(t *testing.T) {
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

	for index := range testCases {
		ctx := context.Background()
		insertedBook, err := s.repo.Insert(ctx, &testCases[index].input)
		if err != nil {
			assert.Equal(t, testCases[index].expectedError, err)
		}

		assert.Equal(t, testCases[index].expected, insertedBook)
	}
}

func (s *Suite) testGet(t *testing.T) {
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

	for _, testCase := range testCases {
		ctx := context.Background()
		receivedBook, err := s.repo.Get(ctx, testCase.input)
		if err != nil {
			assert.Equal(t, testCase.expectedError, err)
		}

		assert.Equal(t, testCase.expected, receivedBook)
	}
}

func (s *Suite) testUpdate(t *testing.T) {
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

	for index := range testCases {
		ctx := context.Background()
		updatedBook, err := s.repo.Update(ctx, testCases[index].input, &testCases[index].toUpdate)
		if err != nil {
			assert.Equal(t, testCases[index].expectedError, err)
		}

		assert.Equal(t, testCases[index].expected, updatedBook)
	}
}

func (s *Suite) testDelete(t *testing.T) {
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

	for _, testCase := range testCases {
		ctx := context.Background()
		deletedBook, err := s.repo.Delete(ctx, testCase.input)
		if err != nil {
			assert.Equal(t, testCase.expectedError, err)
		}

		assert.Equal(t, testCase.expected, deletedBook)
	}
}
