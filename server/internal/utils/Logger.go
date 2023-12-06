package utils

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLogger() {
	// Create a Zap config
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// Initialize the custom_logger
	var err error
	if logger, err = cfg.Build(); err != nil {
		panic("can't initialize custom_logger: " + err.Error())
	}
}

// Close the custom_logger
func Close() error {
	return logger.Sync()
}

// GetLogger Get the initialized custom_logger
func GetLogger() *zap.Logger {
	return logger
}
