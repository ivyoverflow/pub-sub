// Package service_test contains all tests for BookRepository.
package service_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/book/internal/model"
	"github.com/ivyoverflow/pub-sub/book/internal/service"
	"github.com/ivyoverflow/pub-sub/book/internal/store"
	mock "github.com/ivyoverflow/pub-sub/book/internal/store/mock"
)

func TestBook_Add(t *testing.T) {
	testCases := []struct {
		name        string
		input       model.Book
		expectation func(input *model.Book, bookRepository *mock.MockBookRepository)
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
			expectation: func(input *model.Book, bookRepository *mock.MockBookRepository) {
				bookRepository.EXPECT().Add(input).Return(input, nil)
			},
		},
		{
			name:  "Empty JSON body",
			input: model.Book{},
			expectation: func(input *model.Book, bookRepository *mock.MockBookRepository) {
				bookRepository.EXPECT().Add(input).Return(input, errors.New("EOF"))
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
			expectation: func(input *model.Book, bookRepository *mock.MockBookRepository) {
				bookRepository.EXPECT().Add(input).Return(input, errors.New("pq: duplicate key value violates unique constraint \"books_name_key\""))
			},
		},
	}

	for _, testCase := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookRepository := mock.NewMockBookRepository(ctrl)
		svc := service.NewBook(&store.Store{Book: bookRepository})
		testCase.expectation(&testCase.input, bookRepository)

		createdBook, err := svc.Add(&testCase.input)
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
		expectation func(bookID string, expected *model.Book, bookRepository *mock.MockBookRepository)
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
			expectation: func(bookID string, expected *model.Book, bookRepository *mock.MockBookRepository) {
				bookRepository.EXPECT().Get(bookID).Return(expected, nil)
			},
		},
		{
			name:     "Book not found",
			bookID:   service.GenerateUniqueID(),
			expected: model.Book{},
			expectation: func(bookID string, expected *model.Book, bookRepository *mock.MockBookRepository) {
				bookRepository.EXPECT().Get(bookID).Return(expected, errors.New("sql: no rows in result set"))
			},
		},
	}

	for _, testCase := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookRepository := mock.NewMockBookRepository(ctrl)
		svc := service.NewBook(&store.Store{Book: bookRepository})
		testCase.expectation(testCase.bookID, &testCase.expected, bookRepository)

		receivedBook, err := svc.Get(testCase.bookID)
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
		expectation func(bookID string, input *model.Book, bookRepository *mock.MockBookRepository)
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
			expectation: func(bookID string, input *model.Book, bookRepository *mock.MockBookRepository) {
				bookRepository.EXPECT().Update(bookID, input).Return(input, nil)
			},
		},
		{
			name:  "Empty JSON body",
			input: model.Book{},
			expectation: func(bookID string, input *model.Book, bookRepository *mock.MockBookRepository) {
				bookRepository.EXPECT().Update(bookID, input).Return(input, errors.New("EOF"))
			},
		},
		{
			name:   "Wrong book ID",
			bookID: "service.GenerateUniqueID()",
			expectation: func(bookID string, input *model.Book, bookRepository *mock.MockBookRepository) {
				bookRepository.EXPECT().Update(bookID, input).Return(input, errors.New("sql: no rows in result set"))
			},
		},
	}

	for _, testCase := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookRepository := mock.NewMockBookRepository(ctrl)
		svc := service.NewBook(&store.Store{Book: bookRepository})
		testCase.expectation(testCase.bookID, &testCase.input, bookRepository)

		updatedBook, err := svc.Update(testCase.bookID, &testCase.input)
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
		expectation func(bookID string, expected *model.Book, bookRepository *mock.MockBookRepository)
		err         error
	}{
		{
			name:   "OK",
			bookID: service.GenerateUniqueID(),
			expectation: func(bookID string, expected *model.Book, bookRepository *mock.MockBookRepository) {
				bookRepository.EXPECT().Delete(bookID).Return(expected, nil)
			},
		},
		{
			name:   "Wrong book ID",
			bookID: "service.GenerateUniqueID()",
			expectation: func(bookID string, expected *model.Book, bookRepository *mock.MockBookRepository) {
				bookRepository.EXPECT().Delete(bookID).Return(expected, errors.New("sql: no rows in result set"))
			},
		},
	}

	for _, testCase := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookRepository := mock.NewMockBookRepository(ctrl)
		svc := service.NewBook(&store.Store{Book: bookRepository})
		testCase.expectation(testCase.bookID, &testCase.expected, bookRepository)

		deletedBook, err := svc.Delete(testCase.bookID)
		if err != nil {
			if testCase.err != nil {
				assert.Equal(t, testCase.err.Error(), err.Error())
			}
		}

		assert.Equal(t, &testCase.expected, deletedBook)
	}
}
