package logger

import "go.uber.org/zap"

type Logger struct {
	*zap.Logger
}

func NewLogger() (*Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	defer logger.Sync()

	logger.Info("The logger is successfully configured")

	return &Logger{logger}, nil
}
