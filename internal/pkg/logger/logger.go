package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

// New creates a new logger instance based on environment.
func New(env string) *Logger {
	var zapLogger *zap.Logger
	if env == "production" {
		zapLogger, _ = zap.NewProduction()
	} else {
		zapLogger, _ = zap.NewDevelopment()
	}

	return &Logger{zapLogger}
}

// Sync flushes any buffered log entries.
func (l *Logger) Sync() {
	_ = l.Logger.Sync()
}
