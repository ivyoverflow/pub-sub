package main

import (
	"testing"
)

func TestPublisherSubscriber(t *testing.T) {
	pubSub := newPublisherSubscriber()
	channel := make(chan interface{})
	testCases := []struct {
		topic   string
		message interface{}
		result  interface{}
	}{
		{"news", "Ohhhh...", "Ohhhh..."},
		{"news", "Uhhhh...", "Uhhhh..."},
		{"games", "Jhhhh...", "Jhhhh..."},
	}

	for _, testCase := range testCases {
		channel = pubSub.subscribe(testCase.topic)
		pubSub.publish(testCase.topic, testCase.message)
		if <-channel != testCase.result {
			t.Error("Channel contains different message")
		}
	}
}
