package mixers

import (
	logging "github.com/codemodify/systemkit-logging"
)

type asyncLogger struct {
	logger logging.Logger
}

// NewAsync -
func NewAsyncLogger(logger logging.Logger) logging.Logger {
	return &asyncLogger{
		logger: logger,
	}
}

func (thisRef asyncLogger) Log(logEntry logging.LogEntry) logging.LogEntry {
	go func(theLogEntryCopy logging.LogEntry) {
		thisRef.logger.Log(theLogEntryCopy)
	}(logEntry)

	return logging.LogEntry{}
}
