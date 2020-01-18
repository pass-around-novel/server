package console

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	logger ".."
	"../../cmd"
)

type target struct{}

var (
	timeFormat            string
	primaryLevelFormats   []string
	secondaryLevelFormats []string
	levels                = make(map[string]*logger.LogLevel)
	mutex                 sync.Mutex
)

const checkStackFrames = 10

// GetLevel determines the logging level for a specific logger
func (target) GetLevel(name string) *logger.LogLevel {
	if level, ok := levels[name]; ok {
		return level
	}
	mutex.Lock()
	defer mutex.Unlock()
	level := computeLevel(name)
	levels[name] = &level
	return &level
}

// Log saves a message to the log.  It should already be filtered by level.
func (target) Log(name string, level logger.LogLevel, msg string, t time.Time) {
	if level < 0 || int(level) >= len(primaryLevelFormats) {
		pc := make([]uintptr, checkStackFrames)
		n := runtime.Callers(1, pc)
		frames := runtime.CallersFrames(pc[:n])
		baseFrame, more := frames.Next()
		for more {
			var frame runtime.Frame
			frame, more = frames.Next()
			if baseFrame.Func == frame.Func {
				fmt.Printf("Unknown logging level '%d'.\n", level)
				return
			}
		}
		l.Warnf("Unknown logging level '%d'.", level)
	} else {
		tmstr := t.Format(timeFormat)
		var stream *os.File
		if level >= logger.Warning {
			stream = os.Stderr
		} else {
			stream = os.Stdout
		}
		levelFormats := primaryLevelFormats
		for _, line := range strings.Split(msg, "\n") {
			fmt.Fprintf(stream, levelFormats[level], tmstr, line)
			levelFormats = secondaryLevelFormats
		}
	}
}

func init() {
	logger.AddTarget(&target{})
	cmd.ConfigUpdated = append(cmd.ConfigUpdated, configUpdated)
	cmd.GetString("log.console.timeFormat", &timeFormat)
	cmd.GetStringSlice("log.console.primaryFormats", &primaryLevelFormats)
	cmd.GetStringSlice("log.console.secondaryFormats", &secondaryLevelFormats)
}

// Init causes the package to load
func Init() {}

func configUpdated() {
	mutex.Lock()
	defer mutex.Unlock()
	for name, level := range levels {
		*level = computeLevel(name)
	}
}
