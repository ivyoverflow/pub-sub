package postgres_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/constants"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/types"
	"github.com/ivyoverflow/pub-sub/book/internal/repository/postgres"
)

func TestPostgres_New(t *testing.T) {
	testCases := []struct {
		name     string
		cfg      *config.PostgresConfig
		expected error
	}{
		{
			name: "OK",
			cfg: &config.PostgresConfig{
				Host:           constants.PostgresHost,
				Port:           constants.PostgresPort,
				Name:           constants.PostgresName,
				User:           constants.PostgresUser,
				Password:       constants.PostgresPassword,
				MigartionsPath: constants.PostgresMigartionsPath,
				SSLMode:        constants.PostgresSSLMode,
			},
			expected: nil,
		},
		{
			name:     "Invalid connection URI",
			cfg:      &config.PostgresConfig{},
			expected: types.ErrorPostgresConnectionRefused,
		},
		{
			name: "Invalid migrations path",
			cfg: &config.PostgresConfig{
				Host:           constants.PostgresHost,
				Port:           constants.PostgresPort,
				Name:           constants.PostgresName,
				User:           constants.PostgresUser,
				Password:       constants.PostgresPassword,
				MigartionsPath: "",
				SSLMode:        constants.PostgresSSLMode,
			},
			expected: types.ErrorMigrate,
		},
	}

	for _, testCase := range testCases {
		_, err := postgres.New(testCase.cfg)
		assert.Equal(t, testCase.expected, err)
	}
}
