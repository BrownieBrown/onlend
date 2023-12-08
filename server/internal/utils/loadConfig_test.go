package utils_test

import (
	"server/internal/utils"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Set up a defer function to unset environment variables after the test
	defer utils.UnsetEnvVars()

	// Test case: All environment variables are set
	utils.SetEnvVars()
	_, err := utils.LoadConfig()
	if err != nil {
		t.Errorf("LoadConfig() with all env vars should not error, got: %v", err)
	}

	// Test case: Missing environment variables
	utils.UnsetEnvVars()
	_, err = utils.LoadConfig()
	if err == nil {
		t.Error("LoadPostgresConfig() with missing env vars should error, got nil")
	}
}
