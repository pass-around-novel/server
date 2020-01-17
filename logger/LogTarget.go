package logger

import (
	"sync"
	"time"
)

// LogTarget describes a location in which logs can be sent to by the Logger
type LogTarget interface {
	// GetLevel determines the logging level for a specific logger
	GetLevel(name string) *LogLevel
	// Log saves a message to the log.  It should already be filtered by level.
	Log(name string, level LogLevel, msg string, t time.Time)
}

var (
	targets     []LogTarget
	targetMutex sync.Mutex
)

func mapLoggerTarget(logger *logger, target LogTarget) {
	level := target.GetLevel(logger.name)
	logger.targetLevels = append(logger.targetLevels, level)
}

// AddTarget registers a new logging target
func AddTarget(target LogTarget) {
	targetMutex.Lock()
	defer targetMutex.Unlock()
	targets = append(targets, target)
	for _, logger := range loggers {
		mapLoggerTarget(logger, target)
	}
}
