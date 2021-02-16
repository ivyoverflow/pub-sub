package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/notifier/internal/config"
	"github.com/ivyoverflow/pub-sub/notifier/internal/redis"
	"github.com/ivyoverflow/pub-sub/notifier/internal/repository"
)

func TestPublish_repository(t *testing.T) {
	testCase := struct {
		name        string
		book        string
		messsage    interface{}
		expectedErr error
	}{
		name:        "OK",
		book:        "Go in Action",
		messsage:    `"Go in Action" book is available!`,
		expectedErr: nil,
	}

	ctx := context.Background()
	cfg := config.NewRedis()
	db, err := redis.NewDB(ctx, cfg)
	if err != nil {
		t.Errorf("Redis DB initialization throws an error")
	}

	notification := repository.NewNotification(db)
	err = notification.Publish(ctx, testCase.book, testCase.messsage)
	if err != nil {
		t.Errorf("Publish method throws an error")
	}

	assert.Equal(t, testCase.expectedErr, err)
}

func TestSubscribe_repository(t *testing.T) {
	testCase := struct {
		name string
		book string
	}{
		name: "OK",
		book: "Go in Action",
	}

	ctx := context.Background()
	cfg := config.NewRedis()
	db, err := redis.NewDB(ctx, cfg)
	if err != nil {
		t.Errorf("Redis DB initialization throws an error")
	}

	notification := repository.NewNotification(db)
	notification.Subscribe(ctx, testCase.book)
}
