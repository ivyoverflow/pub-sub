package handler_test

import (
	"bytes"
	"context"
	"errors"
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
	"github.com/ivyoverflow/pub-sub/book/internal/model"
	"github.com/ivyoverflow/pub-sub/book/internal/service"
	svcmock "github.com/ivyoverflow/pub-sub/book/internal/service/mock"
	repomock "github.com/ivyoverflow/pub-sub/book/internal/storage/mock"
	"github.com/ivyoverflow/pub-sub/platform/logger"
)

func TestBookHandler_Insert(t *testing.T) {
	testCases := []struct {
		name                    string
		inputString             string
		expectedString          string
		mockBehaviorIDGenerator func(*svcmock.MockGeneratorService)
		expectedJSON            *model.Book
		mockBehaviorBook        func(context.Context, *model.Book, *repomock.MockBookerRepository)
		expectedStatusCode      int
	}{
		{
			name:        "OK",
			inputString: `{"name":"Concurrency in Go: Tools and Techniques for Developers","dateOfIssue":"2017","author":"Katherine Cox-Buday","description":"...","rating":99.99,"price":199.99,"inStock":true}`,
			expectedString: fmt.Sprintf(`{"id":"%s","name":"Concurrency in Go: Tools and Techniques for Developers","dateOfIssue":"2017","author":"Katherine Cox-Buday","description":"...","rating":"99.99","price":"199.99","inStock":true}
`,
				uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003")),
			mockBehaviorIDGenerator: func(gen *svcmock.MockGeneratorService) {
				gen.EXPECT().GenerateUUID().Return(uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"))
			},
			expectedJSON: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			mockBehaviorBook: func(ctx context.Context, expected *model.Book, repo *repomock.MockBookerRepository) {
				repo.EXPECT().Insert(gomock.Any(), expected).Return(expected, nil)
			},
			expectedStatusCode: 201,
		},
		{
			name:        "OK",
			inputString: `{"name":"Introduction to Go","dateOfIssue":"2017","author":"Katherine Cox-Buday","description":"...","rating":99.99,"price":199.99,"inStock":true}`,
			expectedString: fmt.Sprintf(`{"id":"%s","name":"Introduction to Go","dateOfIssue":"2017","author":"Katherine Cox-Buday","description":"...","rating":"99.99","price":"199.99","inStock":true}
`,
				uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120004")),
			mockBehaviorIDGenerator: func(gen *svcmock.MockGeneratorService) {
				gen.EXPECT().GenerateUUID().Return(uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120004"))
			},
			expectedJSON: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120004"),
				Name:        "Introduction to Go",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			mockBehaviorBook: func(ctx context.Context, expected *model.Book, repo *repomock.MockBookerRepository) {
				repo.EXPECT().Insert(gomock.Any(), expected).Return(expected, nil)
			},
			expectedStatusCode: 201,
		},
		{
			name:           "Insert method throws an error: duplicate value",
			inputString:    `{"name":"Concurrency in Go: Tools and Techniques for Developers","dateOfIssue":"2017","author":"Katherine Cox-Buday","description": "...","rating":99.99,"price":199.99,"inStock":true}`,
			expectedString: `{"error": {"statusCode": 409, "message": "duplicate value"}}`,
			mockBehaviorIDGenerator: func(gen *svcmock.MockGeneratorService) {
				gen.EXPECT().GenerateUUID().Return(uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120004"))
			},
			expectedJSON: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120004"),
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			mockBehaviorBook: func(ctx context.Context, expected *model.Book, repo *repomock.MockBookerRepository) {
				repo.EXPECT().Insert(gomock.Any(), expected).Return(nil, types.ErrorDuplicateValue)
			},
			expectedStatusCode: 409,
		},
		{
			name:                    "Insert method throws an error: invalid JSON value type",
			inputString:             `{"name":"jfjwoaopfopwa","dateOfIssue":"2017","author":"Katherine Cox-Buday","description": 111,"rating":99.99,"price":199.99,"inStock":true}`,
			expectedString:          `{"error": {"statusCode": 400, "message": "bad request"}}`,
			mockBehaviorIDGenerator: func(gen *svcmock.MockGeneratorService) {},
			mockBehaviorBook:        func(ctx context.Context, expected *model.Book, repo *repomock.MockBookerRepository) {},
			expectedStatusCode:      400,
		},
		{
			name:                    "Insert method throws an error: invalid JSON body",
			inputString:             `{}`,
			expectedString:          `{"error": {"statusCode": 400, "message": "received JSON is invalid"}}`,
			mockBehaviorIDGenerator: func(gen *svcmock.MockGeneratorService) {},
			mockBehaviorBook:        func(ctx context.Context, expected *model.Book, repo *repomock.MockBookerRepository) {},
			expectedStatusCode:      400,
		},
		{
			name:           "Insert method throws an error: internal service error",
			inputString:    `{"name":"Hello World","dateOfIssue":"2017","author":"John Bob","description":"...","rating":99.99,"price":199.99,"inStock":true}`,
			expectedString: `{"error": {"statusCode": 500, "message": "internal server error"}}`,
			expectedJSON: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120005"),
				Name:        "Hello World",
				DateOfIssue: "2017",
				Author:      "John Bob",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			mockBehaviorIDGenerator: func(gen *svcmock.MockGeneratorService) {
				gen.EXPECT().GenerateUUID().Return(uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120005"))
			},
			mockBehaviorBook: func(ctx context.Context, expected *model.Book, repo *repomock.MockBookerRepository) {
				repo.EXPECT().Insert(gomock.Any(), expected).Return(nil, errors.New("something went wrong"))
			},
			expectedStatusCode: 500,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := repomock.NewMockBookerRepository(ctrl)
			gen := svcmock.NewMockGeneratorService(ctrl)
			testCase.mockBehaviorIDGenerator(gen)
			ctx := context.Background()
			testCase.mockBehaviorBook(ctx, testCase.expectedJSON, repo)
			svc := service.NewBookController(repo, gen)
			log, err := logger.New()
			if err != nil {
				t.Errorf("Logger initialization throws an error: %v", err)
			}

			handl := handler.NewBookController(ctx, svc, log)
			router := mux.NewRouter()
			router.HandleFunc("/v1/book/", handl.Insert)
			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/v1/book/", bytes.NewBufferString(testCase.inputString))
			router.ServeHTTP(rec, req)

			assert.Equal(t, testCase.expectedStatusCode, rec.Code)
			assert.Equal(t, testCase.expectedString, rec.Body.String())
		})
	}
}

func TestBookHandler_Get(t *testing.T) {
	testCases := []struct {
		name               string
		inputStringID      string
		inputUUID          uuid.UUID
		expectedJSON       *model.Book
		expectedString     string
		mockBehavior       func(context.Context, uuid.UUID, *model.Book, *svcmock.MockBookerService)
		expectedStatusCode int
	}{
		{
			name:          "OK",
			inputStringID: "7a2f922c-073a-11eb-adc1-0242ac120003",
			inputUUID:     uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
			expectedJSON: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			expectedString: fmt.Sprintf(`{"id":"%s","name":"Concurrency in Go: Tools and Techniques for Developers","dateOfIssue":"2017","author":"Katherine Cox-Buday","description":"...","rating":"99.99","price":"199.99","inStock":true}
`,
				uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003")),
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, expected *model.Book, repo *svcmock.MockBookerService) {
				repo.EXPECT().Get(gomock.Any(), bookID).Return(expected, nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:               "Book Get service method throws an error: invalid UUID ID",
			inputStringID:      "wakldlkawdlklakwdlk",
			expectedJSON:       nil,
			expectedString:     `{"error": {"statusCode": 500, "message": "internal server error"}}`,
			mockBehavior:       func(ctx context.Context, bookID uuid.UUID, expected *model.Book, repo *svcmock.MockBookerService) {},
			expectedStatusCode: 500,
		},
		{
			name:           "Book Get service method throws an error: book not found",
			inputStringID:  "7a2f922c-073a-11eb-adc1-0242ac120002",
			inputUUID:      uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
			expectedJSON:   nil,
			expectedString: `{"error": {"statusCode": 404, "message": "not found"}}`,
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, expected *model.Book, repo *svcmock.MockBookerService) {
				repo.EXPECT().Get(gomock.Any(), bookID).Return(nil, types.ErrorNotFound)
			},
			expectedStatusCode: 404,
		},
		{
			name:          "Book Get service method throws an error: internal service error",
			inputStringID: "7a2f922c-073a-11eb-adc1-0242ac120003",
			inputUUID:     uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
			expectedJSON: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			expectedString: `{"error": {"statusCode": 500, "message": "internal server error"}}`,
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, expected *model.Book, repo *svcmock.MockBookerService) {
				repo.EXPECT().Get(gomock.Any(), bookID).Return(nil, errors.New("something went wrong"))
			},
			expectedStatusCode: 500,
		},
	}

	for _, testCase := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := svcmock.NewMockBookerService(ctrl)
		gen := svcmock.NewMockGeneratorService(ctrl)
		ctx := context.Background()

		testCase.mockBehavior(ctx, testCase.inputUUID, testCase.expectedJSON, repo)

		svc := service.NewBookController(repo, gen)
		log, err := logger.New()
		if err != nil {
			t.Errorf("Logger initialization throws an error: %v", err)
		}

		handl := handler.NewBookController(ctx, svc, log)
		router := mux.NewRouter()
		router.HandleFunc("/v1/book/{id}", handl.Get)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest("GET", fmt.Sprintf("/v1/book/%s", testCase.inputStringID), nil)

		router.ServeHTTP(rec, req)

		assert.Equal(t, testCase.expectedStatusCode, rec.Code)
		assert.Equal(t, testCase.expectedString, rec.Body.String())
	}
}

func TestBookHandler_Update(t *testing.T) {
	testCases := []struct {
		name               string
		inputStringID      string
		inputUUID          uuid.UUID
		inputString        string
		expectedString     string
		toUpdate           model.Book
		expectedJSON       *model.Book
		mockBehavior       func(context.Context, uuid.UUID, *model.Book, *model.Book, *repomock.MockBookerRepository)
		expectedStatusCode int
	}{
		{
			name:          "OK",
			inputStringID: "7a2f922c-073a-11eb-adc1-0242ac120003",
			inputUUID:     uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
			inputString:   `{"name":"Concurrency in Go: Tools and Techniques for Developers","dateOfIssue":"2017","author":"Katherine Cox-Buday","description":"...","rating":99.99,"price":199.99,"inStock":true}`,
			expectedString: fmt.Sprintf(`{"id":"%s","name":"Concurrency in Go: Tools and Techniques for Developers","dateOfIssue":"2017","author":"Katherine Cox-Buday","description":"...","rating":"99.99","price":"199.99","inStock":true}
`,
				uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003")),
			toUpdate: model.Book{
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			expectedJSON: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, book *model.Book, expected *model.Book, repo *repomock.MockBookerRepository) {
				repo.EXPECT().Update(gomock.Any(), bookID, book).Return(expected, nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:           "Book Update service method throws an error: duplicate value",
			inputStringID:  "7a2f922c-073a-11eb-adc1-0242ac120003",
			inputUUID:      uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
			inputString:    `{"name":"Concurrency in Go: Tools and Techniques for Developers","dateOfIssue":"2017","author":"Katherine Cox-Buday","description": "...","rating":99.99,"price":199.99,"inStock":true}`,
			expectedString: `{"error": {"statusCode": 409, "message": "duplicate value"}}`,
			toUpdate: model.Book{
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			expectedJSON: nil,
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, book *model.Book, expected *model.Book, repo *repomock.MockBookerRepository) {
				repo.EXPECT().Update(gomock.Any(), bookID, book).Return(nil, types.ErrorDuplicateValue)
			},
			expectedStatusCode: 409,
		},
		{
			name:               "Book Update service method throws an error: invalid UUID ID",
			inputStringID:      "wakldlkawdlklakwdlk",
			expectedJSON:       nil,
			expectedString:     `{"error": {"statusCode": 500, "message": "internal server error"}}`,
			mockBehavior:       func(context.Context, uuid.UUID, *model.Book, *model.Book, *repomock.MockBookerRepository) {},
			expectedStatusCode: 500,
		},
		{
			name:               "Book Update service method throws an error: invalid JSON value type",
			inputStringID:      "7a2f922c-073a-11eb-adc1-0242ac120003",
			inputString:        `{"name":"jfjwoaopfopwa","dateOfIssue":"2017","author":"Katherine Cox-Buday","description": 111,"rating":99.99,"price":199.99,"inStock":true}`,
			expectedString:     `{"error": {"statusCode": 400, "message": "bad request"}}`,
			mockBehavior:       func(context.Context, uuid.UUID, *model.Book, *model.Book, *repomock.MockBookerRepository) {},
			expectedStatusCode: 400,
		},
		{
			name:           "Book Update service method throws an error: invalid JSON body",
			inputStringID:  "7a2f922c-073a-11eb-adc1-0242ac120003",
			inputString:    `{}`,
			expectedString: `{"error": {"statusCode": 400, "message": "received JSON is invalid"}}`,
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, book *model.Book, expected *model.Book, repo *repomock.MockBookerRepository) {
			},
			expectedStatusCode: 400,
		},
		{
			name:          "Book Update service method throws an error: book not found",
			inputStringID: "7a2f922c-073a-11eb-adc1-0242ac120006",
			inputUUID:     uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120006"),
			inputString:   `{"name":"Concurrency in Go: Tools and Techniques for Developers","dateOfIssue":"2017","author":"Katherine Cox-Buday","description": "...","rating":99.99,"price":199.99,"inStock":true}`,
			expectedJSON:  nil,
			toUpdate: model.Book{
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			expectedString: `{"error": {"statusCode": 404, "message": "not found"}}`,
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, book *model.Book, expected *model.Book, repo *repomock.MockBookerRepository) {
				repo.EXPECT().Update(gomock.Any(), bookID, book).Return(expected, types.ErrorNotFound)
			},
			expectedStatusCode: 404,
		},
		{
			name:           "Update method throws an error: internal service error",
			inputStringID:  "7a2f922c-073a-11eb-adc1-0242ac120003",
			inputUUID:      uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
			inputString:    `{"name":"Concurrency in Go: Tools and Techniques for Developers","dateOfIssue":"2017","author":"Katherine Cox-Buday","description":"...","rating":99.99,"price":199.99,"inStock":true}`,
			expectedString: `{"error": {"statusCode": 500, "message": "internal server error"}}`,
			toUpdate: model.Book{
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, book *model.Book, expected *model.Book, repo *repomock.MockBookerRepository) {
				repo.EXPECT().Update(gomock.Any(), bookID, book).Return(nil, errors.New("something went wrong"))
			},
			expectedStatusCode: 500,
		},
	}

	for _, testCase := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := repomock.NewMockBookerRepository(ctrl)
		gen := svcmock.NewMockGeneratorService(ctrl)
		ctx := context.Background()

		testCase.mockBehavior(ctx, testCase.inputUUID, &testCase.toUpdate, testCase.expectedJSON, repo)

		svc := service.NewBookController(repo, gen)
		log, err := logger.New()
		if err != nil {
			t.Errorf("Logger initialization throws an error: %v", err)
		}

		handl := handler.NewBookController(ctx, svc, log)
		router := mux.NewRouter()
		router.HandleFunc("/v1/book/{id}", handl.Update)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest("PUT", fmt.Sprintf("/v1/book/%s", testCase.inputStringID), bytes.NewBufferString(testCase.inputString))

		router.ServeHTTP(rec, req)

		assert.Equal(t, testCase.expectedStatusCode, rec.Code)
		assert.Equal(t, testCase.expectedString, rec.Body.String())
	}
}

func TestBookHandler_Delete(t *testing.T) {
	testCases := []struct {
		name               string
		inputStringID      string
		inputUUID          uuid.UUID
		expectedJSON       *model.Book
		expectedString     string
		mockBehavior       func(context.Context, uuid.UUID, *model.Book, *svcmock.MockBookerService)
		expectedStatusCode int
	}{
		{
			name:          "OK",
			inputStringID: "7a2f922c-073a-11eb-adc1-0242ac120003",
			inputUUID:     uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
			expectedJSON: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			expectedString: fmt.Sprintf(`{"id":"%s","name":"Concurrency in Go: Tools and Techniques for Developers","dateOfIssue":"2017","author":"Katherine Cox-Buday","description":"...","rating":"99.99","price":"199.99","inStock":true}
`,
				uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003")),
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, expected *model.Book, repo *svcmock.MockBookerService) {
				repo.EXPECT().Delete(gomock.Any(), bookID).Return(expected, nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:               "Delete service method throws an error: invalid UUID ID",
			inputStringID:      "wakldlkawdlklakwdlk",
			expectedJSON:       nil,
			expectedString:     `{"error": {"statusCode": 500, "message": "internal server error"}}`,
			mockBehavior:       func(context.Context, uuid.UUID, *model.Book, *svcmock.MockBookerService) {},
			expectedStatusCode: 500,
		},
		{
			name:           "Book not found",
			inputStringID:  "7a2f922c-073a-11eb-adc1-0242ac120002",
			inputUUID:      uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120002"),
			expectedJSON:   nil,
			expectedString: `{"error": {"statusCode": 404, "message": "not found"}}`,
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, expected *model.Book, repo *svcmock.MockBookerService) {
				repo.EXPECT().Delete(gomock.Any(), bookID).Return(expected, types.ErrorNotFound)
			},
			expectedStatusCode: 404,
		},
		{
			name:          "Service error",
			inputStringID: "7a2f922c-073a-11eb-adc1-0242ac120003",
			inputUUID:     uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
			expectedJSON: &model.Book{
				ID:          uuid.MustParse("7a2f922c-073a-11eb-adc1-0242ac120003"),
				Name:        "Concurrency in Go: Tools and Techniques for Developers",
				DateOfIssue: "2017",
				Author:      "Katherine Cox-Buday",
				Description: `...`,
				Rating:      model.Decimal{Decimal: decimal.NewFromFloat(99.99)},
				Price:       model.Decimal{Decimal: decimal.NewFromFloat(199.99)},
				InStock:     true,
			},
			expectedString: `{"error": {"statusCode": 500, "message": "internal server error"}}`,
			mockBehavior: func(ctx context.Context, bookID uuid.UUID, expected *model.Book, repo *svcmock.MockBookerService) {
				repo.EXPECT().Delete(gomock.Any(), bookID).Return(nil, errors.New("something went wrong"))
			},
			expectedStatusCode: 500,
		},
	}

	for _, testCase := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := svcmock.NewMockBookerService(ctrl)
		gen := svcmock.NewMockGeneratorService(ctrl)
		ctx := context.Background()

		testCase.mockBehavior(ctx, testCase.inputUUID, testCase.expectedJSON, repo)

		svc := service.NewBookController(repo, gen)
		log, err := logger.New()
		if err != nil {
			t.Errorf("Logger initialization throws an error: %v", err)
		}

		handl := handler.NewBookController(ctx, svc, log)
		router := mux.NewRouter()
		router.HandleFunc("/v1/book/{id}", handl.Delete)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest("DELETE", fmt.Sprintf("/v1/book/%s", testCase.inputStringID), nil)

		router.ServeHTTP(rec, req)

		assert.Equal(t, testCase.expectedStatusCode, rec.Code)
		assert.Equal(t, testCase.expectedString, rec.Body.String())
	}
}
