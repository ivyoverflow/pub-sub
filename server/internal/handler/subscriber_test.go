// Package handler_test contains tests for handlers.
package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/platform/logger"
	"github.com/ivyoverflow/pub-sub/server/internal/handler"
	"github.com/ivyoverflow/pub-sub/server/internal/service"
)

func TestSubscribe_handler(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		expected string
	}{
		{
			name:     "OK",
			body:     `{"topic":"news"}`,
			expected: ``,
		},
		{
			name:     "OK",
			body:     `{"topic":"games"}`,
			expected: ``,
		},
		{
			name:     "Wrong JSON field type",
			body:     `{"topic":1}`,
			expected: `{"error": {"statusCode": 400, "message": "json: cannot unmarshal number into Go struct field SubscribeRequest.topic of type string"}}`,
		},
		{
			name:     "Empty request body",
			body:     ``,
			expected: `{"error": {"statusCode": 400, "message": "unexpected end of JSON input"}}`,
		},
	}

	svc := service.NewPublisherSubscriber()
	for _, testCase := range testCases {
		log, err := logger.New()
		if err != nil {
			t.Errorf("Logger initialization throws an error: %v", err)
		}

		handl := handler.NewSubscriber(svc, log)
		mux := http.NewServeMux()
		mux.Handle("/subscribe", websocket.Handler(handl.Subscribe))

		srv := httptest.NewServer(websocket.Handler(handl.Subscribe))
		defer srv.Close()

		url := "ws" + strings.TrimPrefix(srv.URL, "http")
		ws, err := websocket.Dial(url, "", srv.URL)
		if err != nil {
			t.Errorf("Websocket connection throws an error: %v", err)
		}

		defer ws.Close()

		if err := websocket.JSON.Send(ws, testCase.body); err != nil {
			t.Errorf("Websocket request throws an error: %v", err)
		}
	}
}
