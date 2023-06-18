package utils

import (
	"log"
	"os"
)

// Logger is a wrapper around the standard log package to provide consistent logging functionality.
type Logger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

// NewLogger creates a new instance of the Logger.
func NewLogger(infoHandle, warningHandle, errorHandle *os.File) *Logger {
	return &Logger{
		infoLogger:    log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime),
		warningLogger: log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime),
		errorLogger:   log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info logs an informational message.
func (l *Logger) Info(message string) {
	l.infoLogger.Println(message)
}

// Warning logs a warning message.
func (l *Logger) Warning(message string) {
	l.warningLogger.Println(message)
}

// Error logs an error message along with the associated error.
func (l *Logger) Error(message string, err error) {
	l.errorLogger.Printf("%s: %v", message, err)
}
