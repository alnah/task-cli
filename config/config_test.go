package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	t.Parallel()
	t.Run("Valid config", func(t *testing.T) {
		testEnvFile := "test.env"
		err := os.WriteFile(testEnvFile, []byte(`
POSTGRES_NAME=test_db
POSTGRES_USER=test_user
POSTGRES_PASSWORD=test_pass
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
`), 0644)
		require.NoError(t, err)
		defer os.Remove(testEnvFile)

		config, err := LoadConfig(".", "test")
		require.NoError(t, err)

		require.Equal(t, "test_db", config.PostgresName)
		require.Equal(t, "test_user", config.PostgresUser)
		require.Equal(t, "test_pass", config.PostgresPassword)
		require.Equal(t, "localhost", config.PostgresHost)
		require.Equal(t, "5432", config.PostgresPort)
	})

	t.Run("Missing File", func(t *testing.T) {
		_, err := LoadConfig(".", "invalid")
		require.Error(t, err)
	})

	t.Run("Invalid Format", func(t *testing.T) {
		testEnvFile := "invalid.env"
		err := os.WriteFile(testEnvFile, []byte(`
POSTGRES_NAME=test_db
POSTGRES_USER=test_user
POSTGRES_PASSWORD=test_pass
INVALID_LINE
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
`), 0644)
		require.NoError(t, err)
		defer os.Remove(testEnvFile)

		_, err = LoadConfig(".", "invalid")
		require.Error(t, err)
	})
	t.Run("Empty File", func(t *testing.T) {
		testEnvFile := "empty.env"
		err := os.WriteFile(testEnvFile, []byte(``), 0644)
		require.NoError(t, err)
		defer os.Remove(testEnvFile)

		_, err = LoadConfig(".", "empty")
		require.Error(t, err)
	})
}
