package handler_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/notifier/internal/handler"
	repomock "github.com/ivyoverflow/pub-sub/notifier/internal/repository/mock"
	"github.com/ivyoverflow/pub-sub/notifier/internal/service"
	"github.com/ivyoverflow/pub-sub/platform/logger"
)

func TestPublish_handler(t *testing.T) {
	testCases := []struct {
		name         string
		body         string
		book         string
		message      interface{}
		mockBehavior func(context.Context, string, interface{}, *repomock.MockNotifierRepository)
		expected     string
	}{
		{
			name:    "OK",
			body:    `{"book":"Go in Action","message":"Go in Action is available!"}`,
			book:    "Go in Action",
			message: "Go in Action is available!",
			mockBehavior: func(ctx context.Context, book string, message interface{}, repo *repomock.MockNotifierRepository) {
				repo.EXPECT().Publish(gomock.Any(), book, message).Return(nil)
			},
			expected: "",
		},
		{
			name:         "Wrong JSON field type",
			body:         `{"book":1,"message":"Hello World!"}`,
			mockBehavior: func(ctx context.Context, book string, message interface{}, repo *repomock.MockNotifierRepository) {},
			expected:     `{"error": {"statusCode": 400, "message": "json: cannot unmarshal number into Go struct field PublishRequest.book of type string"}}`,
		},
		{
			name:         "Empty request body",
			body:         ``,
			mockBehavior: func(ctx context.Context, book string, message interface{}, repo *repomock.MockNotifierRepository) {},
			expected:     `{"error": {"statusCode": 400, "message": "EOF"}}`,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()
			repo := repomock.NewMockNotifierRepository(ctrl)
			testCase.mockBehavior(ctx, testCase.book, testCase.message, repo)
			log, err := logger.New()
			if err != nil {
				t.Errorf("Logger initialization throws an error: %v", err)
			}

			svc := service.NewNotification(repo)
			notificationCtrl := handler.NewNotification(svc, log)
			mux := http.NewServeMux()
			mux.HandleFunc("/publish", notificationCtrl.Publish)

			rec := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/publish", bytes.NewBufferString(testCase.body))
			if err != nil {
				t.Errorf("HTTP request throws an error: %v", err)
			}

			mux.ServeHTTP(rec, req)

			assert.Equal(t, rec.Body.String(), testCase.expected)
		})
	}
}
