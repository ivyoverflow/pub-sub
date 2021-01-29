package mongo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/constants"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/types"
	"github.com/ivyoverflow/pub-sub/book/internal/repository/mongo"
)

func TestMongo_New(t *testing.T) {
	testCases := []struct {
		name     string
		cfg      *config.MongoConfig
		expected error
	}{
		{
			name: "OK",
			cfg: &config.MongoConfig{
				Host:     constants.MongoHost,
				Port:     constants.MongoPort,
				Name:     constants.MongoName,
				User:     constants.MongoUser,
				Password: constants.MongoPassword,
			},
			expected: nil,
		},
		{
			name:     "Invalid connection URI",
			cfg:      &config.MongoConfig{},
			expected: types.ErrorMongoConnectionRefused,
		},
	}

	for _, testCase := range testCases {
		ctx := context.Background()
		_, err := mongo.New(ctx, testCase.cfg)
		assert.Equal(t, testCase.expected, err)
	}
}
