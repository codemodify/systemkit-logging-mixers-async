package mixers

import (
	logging "github.com/codemodify/systemkit-logging"
)

type asyncLogger struct {
	logger logging.CoreLogger
}

// NewAsync -
func NewAsyncLogger(logger logging.CoreLogger) logging.CoreLogger {
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
