package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "\033[32m[DEBUG] \033[0m", logger.Flags()),
		info:    log.New(writer, "\033[35m[INFO] \033[0m", logger.Flags()),
		warning: log.New(writer, "\033[33m[WARNING] \033[0m", logger.Flags()),
		err:     log.New(writer, "\033[31m[ERROR] \033[0m", logger.Flags()),
		writer:  writer,
	}
}

// Create Non-Formatted Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

// Create Non-Formatted Logs
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

// Create Non-Formatted Logs
func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println(v...)
}

// Create Non-Formatted Logs
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Create Formatted logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

// Create Formatted logs
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

// Create Formatted logs
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

// Create Formatted logs
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
