// Package handler_test contains tests for handlers.
package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/server/internal/handler"
	"github.com/ivyoverflow/pub-sub/server/internal/logger"
	"github.com/ivyoverflow/pub-sub/server/internal/service"
)

func TestPublish_handler(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		expected string
	}{
		{
			name: "OK",
			body: `
			{
				"topic": "news",
				"message": "Ireland's 'brutally misogynistic culture' saw the death of 9,000 babies and children in mother and baby homes, report finds"
			}`,
			expected: ``,
		},
		{
			name: "OK",
			body: `
			{
				"topic": "games",
				"message": "New Indiana Jones Game Coming From Bethesda"
			}`,
			expected: ``,
		},
		{
			name: "Wrong JSON field type",
			body: `
			{
				"topic": 1,
				"message": "Hello World!"	
			}`,
			expected: `{"error": {"statusCode": 400, "message": "json: cannot unmarshal number into Go struct field PublishRequest.topic of type string"}}`,
		},
		{
			name:     "Empty request body",
			body:     ``,
			expected: `{"error": {"statusCode": 400, "message": "EOF"}}`,
		},
		{
			name: "JSON syntax error",
			body: `
			{
				"topic": "news"
				"message": "House impeaches Trump for 'incitement of insurrection'"
			}`,
			expected: `{"error": {"statusCode": 400, "message": "invalid character '"' after object key:value pair"}}`,
		},
	}

	svc := service.NewPublisherSubscriber()
	for _, testCase := range testCases {
		log, err := logger.New()
		if err != nil {
			t.Errorf("Logger initialization throws an error: %v", err)
		}

		handl := handler.NewPublisher(svc, log)
		mux := http.NewServeMux()
		mux.HandleFunc("/publish", handl.Publish)

		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("POST", "/publish", bytes.NewBufferString(testCase.body))
		if err != nil {
			t.Errorf("HTTP request throws an error: %v", err)
		}

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Body.String(), testCase.expected)
	}
}
