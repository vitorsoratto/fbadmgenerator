package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	error   *log.Logger
	warn    *log.Logger
	writter io.Writer
}

func NewLogger() *Logger {
	writer := io.MultiWriter(os.Stdout)

	logger := &Logger{
		debug:   log.New(writer, "[DEBUG] ", log.Ldate|log.Ltime),
		info:    log.New(writer, "[INFO] ", log.Ldate|log.Ltime),
		error:   log.New(writer, "[ERROR] ", log.Ldate|log.Ltime),
		warn:    log.New(writer, "[WARN] ", log.Ldate|log.Ltime),
		writter: writer,
	}

	return logger
}

// Raw logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.error.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

// Formatted logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.error.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}
