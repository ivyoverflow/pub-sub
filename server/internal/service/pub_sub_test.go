// Package service_test cotains tests for Publisher-Subscriber pattern.
package service_test

import (
	"testing"

	"github.com/ivyoverflow/pub-sub/server/internal/service"
)

func TestPublisherSubscriber_service(t *testing.T) {
	testCases := []struct {
		name    string
		topic   string
		message interface{}
		result  interface{}
	}{
		{
			name:    "OK",
			topic:   "news",
			message: "Trump's disastrous end to his shocking presidency",
			result:  "Trump's disastrous end to his shocking presidency",
		},
		{
			name:    "OK",
			topic:   "news",
			message: "FBI warns 'armed protests' being planned at all 50 state capitols and in Washington DC",
			result:  "FBI warns 'armed protests' being planned at all 50 state capitols and in Washington DC",
		},
		{
			name:    "OK",
			topic:   "games",
			message: "Black PS5 Sale Canceled After Site Says It Received Threats To Safety",
			result:  "Black PS5 Sale Canceled After Site Says It Received Threats To Safety",
		},
		{
			name:    "OK",
			topic:   "games",
			message: "Star Wars Video Games Now Live Under The Lucasfilm Games Umbrella",
			result:  "Star Wars Video Games Now Live Under The Lucasfilm Games Umbrella",
		},
	}

	svc := service.NewPublisherSubscriber()
	for _, testCase := range testCases {
		message := svc.Subscribe(testCase.topic)
		svc.Publish(testCase.topic, testCase.message)
		if <-message != testCase.result {
			t.Errorf("The message received does not match what was expected. Expected: %s", testCase.result)
		}
	}
}
