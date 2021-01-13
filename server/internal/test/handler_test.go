package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/server/internal/handler"
	"github.com/ivyoverflow/pub-sub/server/internal/logger"
	"github.com/ivyoverflow/pub-sub/server/internal/model"
	"github.com/ivyoverflow/pub-sub/server/internal/service"
)

func TestPublish_handler(t *testing.T) {
	testCases := []struct {
		body       string
		statusCode int
	}{
		{
			body: `
			{
				"topic": "news",
				"message": "Ireland's 'brutally misogynistic culture' saw the death of 9,000 babies and children in mother and baby homes, report finds"
			}`,
			statusCode: 200,
		},
		{
			body: `
			{
				"topic": "games",
				"message": "New Indiana Jones Game Coming From Bethesda"
			}`,
			statusCode: 200,
		},
	}

	for _, testCase := range testCases {
		svc := service.NewPublisherSubscriber()
		log, err := logger.New()
		if err != nil {
			t.Errorf("Logger throws an error: %v", err)
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
		if recorder.Code != testCase.statusCode {
			t.Errorf("Wrong status code <<< %d >>>.\nExpected: %d", recorder.Code, testCase.statusCode)
		}
	}
}

func TestSubscribe_handler(t *testing.T) {
	testCases := []struct {
		body *model.SubscribeRequest
	}{
		{
			body: &model.SubscribeRequest{
				Topic: "news",
			},
		},
		{
			body: &model.SubscribeRequest{
				Topic: "news",
			},
		},
	}

	for _, testCase := range testCases {
		svc := service.NewPublisherSubscriber()
		log, err := logger.New()
		if err != nil {
			t.Errorf("Logger throws an error: %v", err)
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
