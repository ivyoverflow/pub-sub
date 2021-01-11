package logger

import (
	"log"

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

	defer func() {
		if err = logger.Sync(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	logger.Info("The logger is successfully configured")

	return &Logger{logger}, nil
}
