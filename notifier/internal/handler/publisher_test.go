// Package handler_test contains tests for handlers.
package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/notifier/internal/handler"
	"github.com/ivyoverflow/pub-sub/notifier/internal/service"
	"github.com/ivyoverflow/pub-sub/platform/logger"
)

func TestPublish_handler(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		expected string
	}{
		{
			name:     "OK",
			body:     `{"topic":"news","message":"Ireland's 'brutally misogynistic culture' saw the death of 9,000 babies and children in mother and baby homes, report finds"}`,
			expected: "",
		},
		{
			name:     "OK",
			body:     `{"topic":"games","message":"New Indiana Jones Game Coming From Bethesda"}`,
			expected: "",
		},
		{
			name:     "Wrong JSON field type",
			body:     `{"topic":1,"message":"Hello World!"}`,
			expected: `{"error": {"statusCode": 400, "message": "json: cannot unmarshal number into Go struct field PublishRequest.topic of type string"}}`,
		},
		{
			name:     "Empty request body",
			body:     ``,
			expected: `{"error": {"statusCode": 400, "message": "EOF"}}`,
		},
	}

	svc := service.NewNotifier()
	for _, testCase := range testCases {
		log, err := logger.New()
		if err != nil {
			t.Errorf("Logger initialization throws an error: %v", err)
		}

		handl := handler.NewPublisher(svc, log)
		mux := http.NewServeMux()
		mux.HandleFunc("/publish", handl.Publish)

		rec := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/publish", bytes.NewBufferString(testCase.body))
		if err != nil {
			t.Errorf("HTTP request throws an error: %v", err)
		}

		mux.ServeHTTP(rec, req)

		assert.Equal(t, rec.Body.String(), testCase.expected)
	}
}
