package utils

import (
	"go.uber.org/zap"
)

type Logger interface {
	Sync() error
	Close() error
	GetLogger() *zap.Logger
}

type ZapLogger struct {
	logger *zap.Logger
}

func (l *ZapLogger) Sync() error {
	return l.logger.Sync()
}

func (l *ZapLogger) Close() error {
	return l.logger.Sync()
}

// GetLogger Get the initialized custom_logger
func (l *ZapLogger) GetLogger() *zap.Logger {
	return l.logger
}

func NewZapLogger() (Logger, error) {
	// Create a Zap config
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// Initialize the custom_logger
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return &ZapLogger{
		logger: logger,
	}, nil
}
