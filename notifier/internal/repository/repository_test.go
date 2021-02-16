package repository_test

import (
	"context"
	"testing"

	driver "github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/notifier/internal/config"
	"github.com/ivyoverflow/pub-sub/notifier/internal/redis"
	"github.com/ivyoverflow/pub-sub/notifier/internal/repository"
)

func TestNotification_repository(t *testing.T) {
	testCase := struct {
		book            string
		message         string
		subscribedBook  string
		expectedMessage string
	}{
		book:            "Go in Action",
		message:         `"Go in Action" book is available`,
		subscribedBook:  "Go in Action",
		expectedMessage: `"Go in Action" book is available`,
	}

	ctx := context.Background()
	cfg := config.NewRedis()
	db, err := redis.NewDB(ctx, cfg)
	if err != nil {
		t.Errorf("Redis DB initialization throws an error: %v", err)
	}

	notification := repository.NewNotification(db)
	channel := notification.Subscribe(ctx, testCase.subscribedBook)
	err = notification.Publish(ctx, testCase.book, testCase.message)
	if err != nil {
		t.Errorf("Publish method throws an error: %v", err)
	}

	go func(channel <-chan *driver.Message) {
		result := make(chan string)
		for message := range channel {
			result <- message.Payload
		}

		assert.Equal(t, testCase.expectedMessage, <-result)
	}(channel)
}
