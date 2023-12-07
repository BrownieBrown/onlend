package utils_test

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"server/internal/utils"
	"testing"
)

func TestNewZapLoggerReturnsNonNilLoggerAndNoError(t *testing.T) {
	logger, err := utils.NewZapLogger()

	assert.NotNil(t, logger, "Logger should not be nil")
	assert.NoError(t, err, "Should not return an error")
}

func TestLoggerConfigurations(t *testing.T) {
	logger, err := utils.NewZapLogger()
	assert.NoError(t, err)

	zapLogger := logger.GetLogger()

	currentLevel := zapLogger.Core().Enabled(zap.InfoLevel)
	assert.True(t, currentLevel)
}
