// Package logger contains the logic to initialize the application logger.
package logger

import (
	"go.uber.org/zap"
)

// Logger represents application logger.
type Logger struct {
	*zap.Logger
}

// New returns a new configured Logger object.
func New() (*Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	logger.Info("The logger is successfully configured")

	return &Logger{logger}, nil
}
