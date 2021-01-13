package test

import (
	"testing"

	"github.com/ivyoverflow/pub-sub/server/internal/service"
)

func TestPublisherSubscriber_service(t *testing.T) {
	testCases := []struct {
		topic   string
		message interface{}
		result  interface{}
	}{
		{"news", "Trump's disastrous end to his shocking presidency", "Trump's disastrous end to his shocking presidency"},
		{"news", "FBI warns 'armed protests' being planned at all 50 state capitols and in Washington DC",
			"FBI warns 'armed protests' being planned at all 50 state capitols and in Washington DC"},
		{"games", "Black PS5 Sale Canceled After Site Says It Received Threats To Safety",
			"Black PS5 Sale Canceled After Site Says It Received Threats To Safety"},
		{"games", "Star Wars Video Games Now Live Under The Lucasfilm Games Umbrella",
			"Star Wars Video Games Now Live Under The Lucasfilm Games Umbrella"},
		{"games", "Jared Leto's Spider-Man Spinoff Morbius Delayed Again", "Jared Leto's Spider-Man Spinoff Morbius Delayed Again"},
	}

	pubSub := service.NewPublisherSubscriber()
	for _, testCase := range testCases {
		message := pubSub.Subscribe(testCase.topic)
		pubSub.Publish(testCase.topic, testCase.message)
		if <-message != testCase.result {
			t.Errorf("The message received does not match what was expected. Expected: %s", testCase.result)
		}
	}
}
