package utils_test

import (
	"os"
	"server/internal/utils"
	"testing"
)

func TestLoadPostgresConfig(t *testing.T) {
	// Set up a defer function to unset environment variables after the test
	defer unsetEnvVars()

	// Test case: All environment variables are set
	setEnvVars()
	_, err := utils.LoadPostgresConfig()
	if err != nil {
		t.Errorf("LoadPostgresConfig() with all env vars should not error, got: %v", err)
	}

	// Test case: Missing environment variables
	unsetEnvVars()
	_, err = utils.LoadPostgresConfig()
	if err == nil {
		t.Error("LoadPostgresConfig() with missing env vars should error, got nil")
	}
}

func setEnvVars() {
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_USER", "user")
	os.Setenv("POSTGRES_PASSWORD", "password")
	os.Setenv("POSTGRES_NAME", "dbname")
	os.Setenv("POSTGRES_SSL_MODE", "disable")
}

func unsetEnvVars() {
	os.Unsetenv("POSTGRES_HOST")
	os.Unsetenv("POSTGRES_PORT")
	os.Unsetenv("POSTGRES_USER")
	os.Unsetenv("POSTGRES_PASSWORD")
	os.Unsetenv("POSTGRES_NAME")
	os.Unsetenv("POSTGRES_SSL_MODE")
}
