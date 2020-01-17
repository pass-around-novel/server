package logger

import (
	"fmt"
	"time"
)

// Logger allows code to write to the log
type Logger interface {
	// GetLevel gets the minimum level that logs are not discarded at
	GetLevel() LogLevel
	// Debug prints a message to the debug log
	Debug(msg string)
	// Debugf formats and prints a message to the debug log
	Debugf(fmt string, args ...interface{})
	// Info prints a message to the info log
	Info(msg string)
	// Intof formats and prints a message to the info log
	Infof(fmt string, args ...interface{})
	// Warn prints a message to the warning log
	Warn(msg string)
	// Warnf formats and prints a message to the warning log
	Warnf(fmt string, args ...interface{})
	// Error prints a message to the error log
	Error(msg string)
	// Errorf formats and prints a message to the error log
	Errorf(fmt string, args ...interface{})
}

type logger struct {
	name         string
	targetLevels []*LogLevel
}

var loggers map[string]*logger = make(map[string]*logger)

// Get a logger with the given name
func Get(name string) Logger {
	if logger, ok := loggers[name]; ok {
		return logger
	}
	logger := &logger{
		name:         name,
		targetLevels: []*LogLevel{},
	}
	targetMutex.Lock()
	defer targetMutex.Unlock()
	for _, target := range targets {
		mapLoggerTarget(logger, target)
	}
	loggers[name] = logger
	return logger
}

// GetLevel gets the minimum level that logs are not discarded at
func (l *logger) GetLevel() LogLevel {
	lowest := None
	for _, level := range l.targetLevels {
		lvl := *level
		if lvl < lowest {
			lowest = lvl
		}
	}
	return lowest
}

func (l *logger) log(level LogLevel, msg string) {
	var (
		t    time.Time
		init bool = false
	)
	for i, target := range l.targetLevels {
		if *target <= level {
			if !init {
				t = time.Now()
				init = true
			}
			targets[i].Log(l.name, level, msg, t)
		}
	}
}

func (l *logger) logf(level LogLevel, format string, args []interface{}) {
	var (
		t    time.Time
		msg  string
		init bool = false
	)
	for i, target := range l.targetLevels {
		if *target <= level {
			if !init {
				t = time.Now()
				msg = fmt.Sprintf(format, args...)
				init = true
			}
			targets[i].Log(l.name, level, msg, t)
		}
	}
}

// Debug prints a message to the debug log
func (l *logger) Debug(msg string) {
	l.log(Debug, msg)
}

// Debugf formats and prints a message to the debug log
func (l *logger) Debugf(fmt string, args ...interface{}) {
	l.logf(Debug, fmt, args)
}

// Info prints a message to the info log
func (l *logger) Info(msg string) {
	l.log(Info, msg)
}

// Intof formats and prints a message to the info log
func (l *logger) Infof(fmt string, args ...interface{}) {
	l.logf(Info, fmt, args)
}

// Warn prints a message to the warning log
func (l *logger) Warn(msg string) {
	l.log(Warning, msg)
}

// Warnf formats and prints a message to the warning log
func (l *logger) Warnf(fmt string, args ...interface{}) {
	l.logf(Warning, fmt, args)
}

// Error prints a message to the error log
func (l *logger) Error(msg string) {
	l.log(Error, msg)
}

// Errorf formats and prints a message to the error log
func (l *logger) Errorf(fmt string, args ...interface{}) {
	l.logf(Error, fmt, args)
}
