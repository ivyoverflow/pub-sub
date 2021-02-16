package redis_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/notifier/internal/config"
	"github.com/ivyoverflow/pub-sub/notifier/internal/redis"
)

func TestNewDB_redis(t *testing.T) {
	testCases := []struct {
		name        string
		cfg         config.RedisConfig
		expectedErr error
	}{
		{
			name: "OK",
			cfg: config.RedisConfig{
				Addr: "localhost:6379",
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			ctx := context.Background()
			_, err := redis.NewDB(ctx, &testCase.cfg)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}
