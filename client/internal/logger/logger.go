package logger

import "go.uber.org/zap"

// Logger represents application logger.
type Logger struct {
	*zap.Logger
}

// NewLogger returns a new configured Logger object.
func NewLogger() (*Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	defer logger.Sync()

	logger.Info("The logger is successfully configured")

	return &Logger{logger}, nil
}
