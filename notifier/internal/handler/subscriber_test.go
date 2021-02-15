// Package handler_test contains tests for handlers.
package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/notifier/internal/handler"
	"github.com/ivyoverflow/pub-sub/notifier/internal/service"
	"github.com/ivyoverflow/pub-sub/platform/logger"
)

func TestSubscribe_handler(t *testing.T) {
	testCases := []struct {
		name     string
		subInput string
		pubInput string
		expected string
	}{
		{
			name:     "OK",
			subInput: `{"topic": "news"}`,
			pubInput: `{"topic": "news", "message": "..."}`,
			expected: ``,
		},
		{
			name:     "OK",
			subInput: `{"topic": "games"}`,
			pubInput: `{"topic": "games", "message": "..."}`,
			expected: ``,
		},
		{
			name:     "Wrong JSON field type",
			subInput: `{"topic": 1}`,
			expected: `{"error": {"statusCode": 400, "message": "json: cannot unmarshal number into Go struct field SubscribeRequest.topic of type string"}}`,
		},
		{
			name:     "Empty request body",
			subInput: ``,
			expected: `{"error": {"statusCode": 400, "message": "unexpected end of JSON input"}}`,
		},
	}

	for _, testCase := range testCases {
		svc := service.NewPublisherSubscriber()
		log, err := logger.New()
		if err != nil {
			t.Errorf("Logger initialization throws an error: %v", err)
		}

		pub := handler.NewPublisher(svc, log)
		sub := handler.NewSubscriber(svc, log)
		mux := http.NewServeMux()
		mux.Handle("/subscribe", websocket.Handler(sub.Subscribe))
		mux.HandleFunc("/publish", pub.Publish)

		subSrv := httptest.NewServer(websocket.Handler(sub.Subscribe))
		defer subSrv.Close()

		pubSrv := httptest.NewServer(http.HandlerFunc(pub.Publish))
		defer pubSrv.Close()

		url := "ws" + strings.TrimPrefix(subSrv.URL, "http")
		ws, err := websocket.Dial(url, "", subSrv.URL)
		if err != nil {
			t.Errorf("Websocket connection throws an error: %v", err)
		}

		defer ws.Close()

		if err := websocket.Message.Send(ws, testCase.subInput); err != nil {
			t.Errorf("Websocket request throws an error: %v", err)
		}
	}
}
