package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	repomock "github.com/ivyoverflow/pub-sub/notifier/internal/repository/mock"
	"github.com/ivyoverflow/pub-sub/notifier/internal/service"
)

func TestPublish_service(t *testing.T) {
	testCase := struct {
		name         string
		book         string
		message      interface{}
		mockBehavior func(context.Context, string, interface{}, *repomock.MockNotifierRepository)
		expectedErr  error
	}{
		name:    "OK",
		book:    "Go in Action",
		message: `"Go in Action" book is available!`,
		mockBehavior: func(ctx context.Context, book string, message interface{}, repo *repomock.MockNotifierRepository) {
			repo.EXPECT().Publish(ctx, book, message).Return(nil)
		},
		expectedErr: nil,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	repo := repomock.NewMockNotifierRepository(ctrl)
	testCase.mockBehavior(ctx, testCase.book, testCase.message, repo)
	svc := service.NewNotification(repo)
	err := svc.Publish(ctx, testCase.book, testCase.message)
	assert.Equal(t, testCase.expectedErr, err)
}
