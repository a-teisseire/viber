package viber

import (
	"fmt"
	"io"
	"log"
)

// Logger provides a standardized interface for logging
type Logger interface {
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(string, ...interface{})
}

// NewDefaultLogger creates a default logger which uses the standard logger
func NewDefaultLogger(w io.Writer) Logger {
	return &defaultLogger{
		logger: log.New(w, "Viber >>", 0),
	}
}

type defaultLogger struct {
	logger *log.Logger
}

func (l defaultLogger) Debug(format string, params ...interface{}) {
	l.logger.Printf(fmt.Sprintf("[D] %s", format), params...)
}

func (l defaultLogger) Info(format string, params ...interface{}) {
	l.logger.Printf(fmt.Sprintf("[I] %s", format), params...)
}

func (l defaultLogger) Warn(format string, params ...interface{}) {
	l.logger.Printf(fmt.Sprintf("[W] %s", format), params...)
}

func (l defaultLogger) Error(format string, params ...interface{}) {
	l.logger.Printf(fmt.Sprintf("[E] %s", format), params...)
}
