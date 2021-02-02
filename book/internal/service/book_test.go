package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/book/internal/lib/types"
	"github.com/ivyoverflow/pub-sub/book/internal/model"
	mock "github.com/ivyoverflow/pub-sub/book/internal/repository/mock"
	"github.com/ivyoverflow/pub-sub/book/internal/service"
)

func TestBookService_Insert(t *testing.T) {
	testCases := []struct {
		name          string
		input         model.Book
		expected      *model.Book
		mockBehavior  func(context.Context, *model.Book, *model.Book, *mock.MockBookI)
		expectedError error
	}{
		{
			name: "OK",
			input: model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			expected: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			mockBehavior: func(ctx context.Context, book *model.Book, expected *model.Book, repo *mock.MockBookI) {
				repo.EXPECT().Insert(ctx, book).Return(expected, nil)
			},
			expectedError: nil,
		},
		{
			name: "Duplicate value",
			input: model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			expected: nil,
			mockBehavior: func(ctx context.Context, book *model.Book, expected *model.Book, repo *mock.MockBookI) {
				repo.EXPECT().Insert(ctx, book).Return(expected, types.ErrorDuplicateValue)
			},
			expectedError: types.ErrorDuplicateValue,
		},
		{
			name: "Invalid body",
			input: model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
			},
			expected:      nil,
			mockBehavior:  func(ctx context.Context, book *model.Book, expected *model.Book, repo *mock.MockBookI) {},
			expectedError: types.ErrorValidation,
		},
	}

	for _, testCase := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock.NewMockBookI(ctrl)
		gen := service.NewIDGenerator()
		svc := service.NewBook(repo, gen)
		ctx := context.Background()

		testCase.mockBehavior(ctx, &testCase.input, testCase.expected, repo)

		insertedBook, err := svc.Insert(ctx, &testCase.input)
		if err != nil {
			assert.Equal(t, testCase.expectedError, err)
		}

		assert.Equal(t, testCase.expected, insertedBook)
	}
}

func TestBookService_Get(t *testing.T) {
	testCases := []struct {
		name          string
		input         uuid.UUID
		expected      *model.Book
		mockBehavior  func(context.Context, uuid.UUID, *model.Book, *mock.MockBookI)
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
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, expected *model.Book, repo *mock.MockBookI) {
				repo.EXPECT().Get(ctx, bookID).Return(expected, nil)
			},
			expectedError: nil,
		},
		{
			name:     "Book not found",
			input:    uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120004"),
			expected: nil,
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, expected *model.Book, repo *mock.MockBookI) {
				repo.EXPECT().Get(ctx, bookID).Return(expected, types.ErrorNotFound)
			},
			expectedError: types.ErrorNotFound,
		},
	}

	for _, testCase := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock.NewMockBookI(ctrl)
		gen := service.NewIDGenerator()
		svc := service.NewBook(repo, gen)
		ctx := context.Background()

		testCase.mockBehavior(ctx, testCase.input, testCase.expected, repo)

		insertedBook, err := svc.Get(ctx, testCase.input)
		if err != nil {
			assert.Equal(t, testCase.expectedError, err)
		}

		assert.Equal(t, testCase.expected, insertedBook)
	}
}

func TestBookService_Update(t *testing.T) {
	testCases := []struct {
		name          string
		input         uuid.UUID
		toUpdate      model.Book
		expected      *model.Book
		mockBehavior  func(context.Context, uuid.UUID, *model.Book, *model.Book, *mock.MockBookI)
		expectedError error
	}{
		{
			name:  "OK",
			input: uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
			toUpdate: model.Book{
				Name:        "Concurrency in Go: TTD",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			expected: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
				Name:        "Concurrency in Go: TTD",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, book *model.Book, expected *model.Book, repo *mock.MockBookI) {
				repo.EXPECT().Update(ctx, bookID, book).Return(expected, nil)
			},
			expectedError: nil,
		},
		{
			name:  "Book not found",
			input: uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120004"),
			toUpdate: model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120004"),
				Name:        "Concurrency in Go: TTD",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			expected: nil,
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, book *model.Book, expected *model.Book, repo *mock.MockBookI) {
				repo.EXPECT().Update(ctx, bookID, book).Return(expected, types.ErrorNotFound)
			},
			expectedError: types.ErrorNotFound,
		},
		{
			input: uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
			name:  "Invalid body",
			toUpdate: model.Book{
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
			},
			expected:      nil,
			mockBehavior:  func(context.Context, uuid.UUID, *model.Book, *model.Book, *mock.MockBookI) {},
			expectedError: types.ErrorValidation,
		},
	}

	for _, testCase := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock.NewMockBookI(ctrl)
		gen := service.NewIDGenerator()
		svc := service.NewBook(repo, gen)
		ctx := context.Background()

		testCase.mockBehavior(ctx, testCase.input, &testCase.toUpdate, testCase.expected, repo)

		insertedBook, err := svc.Update(ctx, testCase.input, &testCase.toUpdate)
		if err != nil {
			assert.Equal(t, testCase.expectedError, err)
		}

		assert.Equal(t, testCase.expected, insertedBook)
	}
}

func TestBookService_Delete(t *testing.T) {
	testCases := []struct {
		name          string
		input         uuid.UUID
		expected      *model.Book
		mockBehavior  func(context.Context, uuid.UUID, *model.Book, *mock.MockBookI)
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
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, expected *model.Book, repo *mock.MockBookI) {
				repo.EXPECT().Delete(ctx, bookID).Return(expected, nil)
			},
			expectedError: nil,
		},
		{
			name:     "Book not found",
			input:    uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120004"),
			expected: nil,
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, expected *model.Book, repo *mock.MockBookI) {
				repo.EXPECT().Delete(ctx, bookID).Return(expected, types.ErrorNotFound)
			},
			expectedError: types.ErrorNotFound,
		},
	}

	for _, testCase := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock.NewMockBookI(ctrl)
		gen := service.NewIDGenerator()
		svc := service.NewBook(repo, gen)
		ctx := context.Background()

		testCase.mockBehavior(ctx, testCase.input, testCase.expected, repo)

		insertedBook, err := svc.Delete(ctx, testCase.input)
		if err != nil {
			assert.Equal(t, testCase.expectedError, err)
		}

		assert.Equal(t, testCase.expected, insertedBook)
	}
}
