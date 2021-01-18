// Package service_test contains all tests for BookRepository.
package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/book/internal/model"
	mock "github.com/ivyoverflow/pub-sub/book/internal/repository/mock"
	"github.com/ivyoverflow/pub-sub/book/internal/service"
)

func TestBook_Insert(t *testing.T) {
	testCases := []struct {
		name        string
		input       model.Book
		expectation func(ctx context.Context, input *model.Book, bookRepository *mock.MockBookI)
		err         error
	}{
		{
			name: "OK",
			input: model.Book{
				ID:          service.GenerateUniqueID(),
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
			expectation: func(ctx context.Context, input *model.Book, bookRepository *mock.MockBookI) {
				bookRepository.EXPECT().Insert(ctx, input).Return(input, nil)
			},
		},
		{
			name:  "Empty JSON body",
			input: model.Book{},
			expectation: func(ctx context.Context, input *model.Book, bookRepository *mock.MockBookI) {
				bookRepository.EXPECT().Insert(ctx, input).Return(input, errors.New("EOF"))
			},
		},
		{
			name: "Duplicate book name",
			input: model.Book{
				ID:          service.GenerateUniqueID(),
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
			expectation: func(ctx context.Context, input *model.Book, bookRepository *mock.MockBookI) {
				bookRepository.EXPECT().Insert(ctx, input).Return(input, errors.New("pq: duplicate key value violates unique constraint \"books_name_key\""))
			},
		},
	}

	for _, testCase := range testCases {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookRepository := mock.NewMockBookI(ctrl)
		svc := service.NewBook(bookRepository)
		testCase.expectation(ctx, &testCase.input, bookRepository)

		createdBook, err := svc.Insert(ctx, &testCase.input)
		if err != nil {
			if testCase.err != nil {
				assert.Equal(t, testCase.err.Error(), err.Error())
			}
		}

		assert.Equal(t, &testCase.input, createdBook)
	}
}

func TestBook_Get(t *testing.T) {
	testCases := []struct {
		name        string
		bookID      string
		expected    model.Book
		expectation func(ctx context.Context, bookID string, expected *model.Book, bookRepository *mock.MockBookI)
		err         error
	}{
		{
			name:   "OK",
			bookID: service.GenerateUniqueID(),
			expected: model.Book{
				ID:          service.GenerateUniqueID(),
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
			expectation: func(ctx context.Context, bookID string, expected *model.Book, bookRepository *mock.MockBookI) {
				bookRepository.EXPECT().Get(ctx, bookID).Return(expected, nil)
			},
		},
		{
			name:     "Book not found",
			bookID:   service.GenerateUniqueID(),
			expected: model.Book{},
			expectation: func(ctx context.Context, bookID string, expected *model.Book, bookRepository *mock.MockBookI) {
				bookRepository.EXPECT().Get(ctx, bookID).Return(expected, errors.New("sql: no rows in result set"))
			},
		},
	}

	for _, testCase := range testCases {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookRepository := mock.NewMockBookI(ctrl)
		svc := service.NewBook(bookRepository)
		testCase.expectation(ctx, testCase.bookID, &testCase.expected, bookRepository)

		receivedBook, err := svc.Get(ctx, testCase.bookID)
		if err != nil {
			if testCase.err != nil {
				assert.Equal(t, testCase.err.Error(), err.Error())
			}
		}

		assert.Equal(t, &testCase.expected, receivedBook)
	}
}

func TestBook_Update(t *testing.T) {
	testCases := []struct {
		name        string
		bookID      string
		input       model.Book
		expectation func(ctx context.Context, bookID string, input *model.Book, bookRepository *mock.MockBookI)
		err         error
	}{
		{
			name:   "OK",
			bookID: service.GenerateUniqueID(),
			input: model.Book{
				ID:          service.GenerateUniqueID(),
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
				InStock: false,
			},
			expectation: func(ctx context.Context, bookID string, input *model.Book, bookRepository *mock.MockBookI) {
				bookRepository.EXPECT().Update(ctx, bookID, input).Return(input, nil)
			},
		},
		{
			name:  "Empty JSON body",
			input: model.Book{},
			expectation: func(ctx context.Context, bookID string, input *model.Book, bookRepository *mock.MockBookI) {
				bookRepository.EXPECT().Update(ctx, bookID, input).Return(input, errors.New("EOF"))
			},
		},
		{
			name:   "Wrong book ID",
			bookID: "service.GenerateUniqueID()",
			expectation: func(ctx context.Context, bookID string, input *model.Book, bookRepository *mock.MockBookI) {
				bookRepository.EXPECT().Update(ctx, bookID, input).Return(input, errors.New("sql: no rows in result set"))
			},
		},
	}

	for _, testCase := range testCases {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookRepository := mock.NewMockBookI(ctrl)
		svc := service.NewBook(bookRepository)
		testCase.expectation(ctx, testCase.bookID, &testCase.input, bookRepository)

		updatedBook, err := svc.Update(ctx, testCase.bookID, &testCase.input)
		if err != nil {
			if testCase.err != nil {
				assert.Equal(t, testCase.err.Error(), err.Error())
			}
		}

		assert.Equal(t, &testCase.input, updatedBook)
	}
}

func TestBook_Delete(t *testing.T) {
	testCases := []struct {
		name        string
		bookID      string
		expected    model.Book
		expectation func(ctx context.Context, bookID string, expected *model.Book, bookRepository *mock.MockBookI)
		err         error
	}{
		{
			name:   "OK",
			bookID: service.GenerateUniqueID(),
			expectation: func(ctx context.Context, bookID string, expected *model.Book, bookRepository *mock.MockBookI) {
				bookRepository.EXPECT().Delete(ctx, bookID).Return(expected, nil)
			},
		},
		{
			name:   "Wrong book ID",
			bookID: "service.GenerateUniqueID()",
			expectation: func(ctx context.Context, bookID string, expected *model.Book, bookRepository *mock.MockBookI) {
				bookRepository.EXPECT().Delete(ctx, bookID).Return(expected, errors.New("sql: no rows in result set"))
			},
		},
	}

	for _, testCase := range testCases {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookRepository := mock.NewMockBookI(ctrl)
		svc := service.NewBook(bookRepository)
		testCase.expectation(ctx, testCase.bookID, &testCase.expected, bookRepository)

		deletedBook, err := svc.Delete(ctx, testCase.bookID)
		if err != nil {
			if testCase.err != nil {
				assert.Equal(t, testCase.err.Error(), err.Error())
			}
		}

		assert.Equal(t, &testCase.expected, deletedBook)
	}
}
