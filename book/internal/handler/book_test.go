package handler_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/book/internal/handler"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/types"
	"github.com/ivyoverflow/pub-sub/book/internal/logger"
	"github.com/ivyoverflow/pub-sub/book/internal/model"
	mock_repository "github.com/ivyoverflow/pub-sub/book/internal/repository/mock"
	"github.com/ivyoverflow/pub-sub/book/internal/service"
	mock_service "github.com/ivyoverflow/pub-sub/book/internal/service/mock"
)

var (
	ctx = context.Background()
)

func TestBookHandler_Insert(t *testing.T) {
	testCases := []struct {
		name                    string
		inputString             string
		expectedString          string
		mockBehaviorIDGenerator func(*mock_service.MockIDGeneratorI)
		expectedJSON            *model.Book
		mockBehaviorBook        func(context.Context, *model.Book, *mock_repository.MockBookI)
		expectedStatusCode      int
	}{
		{
			name:           "OK",
			inputString:    `{"name":"Concurrency in Go: Tools and Techniques for Developers","dateOfIssue":"2017","author":"Katherine Cox-Buday","description": "Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming language makes working with concurrency tractable and even easy.","rating":71.00,"price":36.90,"inStock":true}`,
			expectedString: fmt.Sprintf(`{"id":"%s","name":"Concurrency in Go: Tools and Techniques for Developers","dateOfIssue":"2017","author":"Katherine Cox-Buday","description": "Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming language makes working with concurrency tractable and even easy.","rating":71.00,"price":36.90,"inStock":true}`, uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003")),
			mockBehaviorIDGenerator: func(gen *mock_service.MockIDGeneratorI) {
				gen.EXPECT().Generate().Return(uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"))
			},
			expectedJSON: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming language makes working with concurrency tractable and even easy.`,
				Rating:      decimal.NewFromFloat(71.0),
				Price:       decimal.NewFromFloat(36.9),
				InStock:     true,
			},
			mockBehaviorBook: func(ctx context.Context, expected *model.Book, repo *mock_repository.MockBookI) {
				repo.EXPECT().Insert(gomock.Any(), expected).Return(expected, nil)
			},
			expectedStatusCode: 201,
		},
		{
			name:           "Duplicate value",
			inputString:    `{"name":"Concurrency in Go: Tools and Techniques for Developers","dateOfIssue":"2017","author":"Katherine Cox-Buday","description": "Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming language makes working with concurrency tractable and even easy.","rating":71.00,"price":36.90,"inStock":true}`,
			expectedString: `{"error": {"statusCode": 409, "message": "duplicate value"}}`,
			mockBehaviorIDGenerator: func(gen *mock_service.MockIDGeneratorI) {
				gen.EXPECT().Generate().Return(uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120004"))
			},
			expectedJSON: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120004"),
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming language makes working with concurrency tractable and even easy.`,
				Rating:      decimal.NewFromFloat(71.0),
				Price:       decimal.NewFromFloat(36.9),
				InStock:     true,
			},
			mockBehaviorBook: func(ctx context.Context, expected *model.Book, repo *mock_repository.MockBookI) {
				repo.EXPECT().Insert(gomock.Any(), expected).Return(nil, types.ErrorDuplicateValue)
			},
			expectedStatusCode: 409,
		},
		{
			name:                    "Invalid JSON body",
			inputString:             `{}`,
			expectedString:          `{"error": {"statusCode": 409, "message": "bad request"}}`,
			mockBehaviorIDGenerator: func(gen *mock_service.MockIDGeneratorI) {},
			mockBehaviorBook:        func(ctx context.Context, expected *model.Book, repo *mock_repository.MockBookI) {},
			expectedStatusCode:      400,
		},
	}

	for _, testCase := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock_repository.NewMockBookI(ctrl)
		gen := mock_service.NewMockIDGeneratorI(ctrl)
		testCase.mockBehaviorIDGenerator(gen)
		testCase.mockBehaviorBook(ctx, testCase.expectedJSON, repo)

		svc := service.NewBook(repo, gen)
		log, err := logger.New()
		if err != nil {
			t.Errorf("Logger initialization throws an error: %v", err)
		}

		handl := handler.NewBook(ctx, svc, log)
		router := mux.NewRouter()
		router.HandleFunc("/v1/book/", handl.Insert)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/v1/book/", bytes.NewBufferString(testCase.inputString))

		router.ServeHTTP(rec, req)

		assert.Equal(t, rec.Code, testCase.expectedStatusCode)
		assert.Equal(t, rec.Body.String(), testCase.expectedString)
	}
}

// func TestBookHandler_Get(t *testing.T) {
// 	testCases := []struct {
// 		name               string
// 		input              uuid.UUID
// 		expectedJSON       *model.Book
// 		expectedString     string
// 		mockBehavior       func(context.Context, uuid.UUID, *model.Book, *mock_service.MockBookI)
// 		expectedStatusCode int
// 	}{
// 		{
// 			name:  "OK",
// 			input: uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
// 			expectedJSON: &model.Book{
// 				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
// 				Name:        "Concurrency in Go: Tools and Techniques for Developers",
// 				DateOfIssue: "2017",
// 				Author:      "Katherine Cox-Buday",
// 				Description: `Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming language makes working with concurrency tractable and even easy.`,
// 				Rating:      71.00,
// 				Price:       36.90,
// 				InStock:     true,
// 			},
// 			expectedString: `{"name":"Concurrency in Go: Tools and Techniques for Developers","dateOfIssue":"2017","author":"Katherine Cox-Buday","description": "Concurrency can be notoriously difficult to get right, but fortunately, the Go open source programming language makes working with concurrency tractable and even easy.","rating":71.00,"price":36.90,"inStock":true}`,
// 			mockBehavior: func(ctx context.Context, bookID uuid.UUID, expected *model.Book, repo *mock_service.MockBookI) {
// 				repo.EXPECT().Get(ctx, bookID).Return(expected, nil)
// 			},
// 			expectedStatusCode: 200,
// 		},
// 		{
// 			name:           "Book not found",
// 			input:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
// 			expectedJSON:   nil,
// 			expectedString: `{"error": {"statusCode": 404, "message": "not found"}}`,
// 			mockBehavior: func(ctx context.Context, bookID uuid.UUID, expected *model.Book, repo *mock_service.MockBookI) {
// 				repo.EXPECT().Get(ctx, bookID).Return(expected, types.ErrorNotFound)
// 			},
// 			expectedStatusCode: 404,
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		ctrl := gomock.NewController(t)
// 		defer ctrl.Finish()

// 		repo := mock_service.NewMockBookI(ctrl)
// 		gen := mock_service.NewMockIDGeneratorI(ctrl)
// 		testCase.mockBehavior(ctx, testCase.input, testCase.expectedJSON, repo)

// 		svc := service.NewBook(repo, gen)
// 		log, err := logger.New()
// 		if err != nil {
// 			t.Errorf("Logger initialization throws an error: %v", err)
// 		}

// 		handl := handler.NewBook(ctx, svc, log)
// 		router := mux.NewRouter()
// 		router.HandleFunc("/v1/book/{id}", handl.Insert)

// 		rec := httptest.NewRecorder()
// 		req := httptest.NewRequest("GET", fmt.Sprintf("/v1/book/%s", testCase.input.String()), nil)

// 		router.ServeHTTP(rec, req)

// 		assert.Equal(t, rec.Code, testCase.expectedStatusCode)
// 		assert.Equal(t, rec.Body.String(), testCase.expectedString)
// 	}
// }

// func TestBookHandler_Update(t *testing.T) {

// }

// func TestBookHandler_Delete(t *testing.T) {

// }
