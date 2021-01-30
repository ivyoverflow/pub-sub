package postgres_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/constants"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/types"
	"github.com/ivyoverflow/pub-sub/book/internal/repository/postgres"
)

func TestPostgres_RunMigration(t *testing.T) {
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
			name: "Empty migrations path",
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
		{
			name: "Invalid migrations path",
			cfg: &config.PostgresConfig{
				Host:           constants.PostgresHost,
				Port:           constants.PostgresPort,
				Name:           constants.PostgresName,
				User:           constants.PostgresUser,
				Password:       constants.PostgresPassword,
				MigartionsPath: "PostgresMigartionsPath",
				SSLMode:        constants.PostgresSSLMode,
			},
			expected: types.ErrorMigrate,
		},
		{
			name: "Invalid Postgres host",
			cfg: &config.PostgresConfig{
				Host:           "",
				Port:           constants.PostgresPort,
				Name:           constants.PostgresName,
				User:           constants.PostgresUser,
				Password:       constants.PostgresPassword,
				MigartionsPath: constants.PostgresMigartionsPath,
				SSLMode:        constants.PostgresSSLMode,
			},
			expected: types.ErrorMigrate,
		},
	}

	for _, testCase := range testCases {
		err := postgres.RunMigration(testCase.cfg)
		assert.Equal(t, testCase.expected, err)
	}
}
