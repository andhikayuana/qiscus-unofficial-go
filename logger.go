package qiscus

import (
	"fmt"
	"os"
)

// LogLevel is the logging level used by the Qiscus go library
type LogLevel uint32

const (
	// NoLogging sets a logger to not show the messages
	NoLogging LogLevel = 0

	// LogError sets a logger to show error messages only.
	LogError LogLevel = 1

	// LogInfo sets a logger to show information messages
	LogInfo LogLevel = 2

	// LogDebug sets a logger to show informational messages for debugging
	LogDebug LogLevel = 3
)

type Logger interface {
	// Error logs a warning message using Printf conventions.
	Error(format string, val ...interface{})

	// Info logs an informational message using Printf conventions.
	Info(format string, val ...interface{})

	// Debug logs debug message using Printf conventions.
	Debug(format string, val ...interface{})
}

// LoggerImpl is a logger interface implementation.
// It prints some info, errors message and debug message for debugging to `os.Stderr` and `os.Stdout`
type LoggerImpl struct {
	LogLevel LogLevel
}

// Error : Logs a warning message using Printf conventions.
func (l *LoggerImpl) Error(format string, val ...interface{}) {
	if l.LogLevel >= LogError {
		fmt.Fprintf(os.Stderr, "ERROR - "+format+"\n", val...)
	}
}

// Info : Logs information message using Printf conventions.
func (l *LoggerImpl) Info(format string, val ...interface{}) {
	if l.LogLevel >= LogInfo {
		fmt.Fprintf(os.Stdout, "INFO - "+format+"\n", val...)
	}
}

// Debug : Log debug message using Printf conventions.
func (l *LoggerImpl) Debug(format string, val ...interface{}) {
	if l.LogLevel >= LogDebug {
		fmt.Fprintf(os.Stdout, "DEBUG - "+format+"\n", val...)
	}
}
