package console

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	logger ".."
	"../../cmd"
)

type target struct{}

var (
	timeFormat          = "15:04:05"
	primaryLevelFormats = []string{
		"[%s] [DBG ] %s\n",
		"[%s] [INFO] %s\n",
		"[%s] [WARN] %s\n",
		"[%s] [ERR ] %s\n",
	}
	secondaryLevelFormats = []string{
		"           [DBG ] %[2]s\n",
		"           [INFO] %[2]s\n",
		"           [WARN] %[2]s\n",
		"           [ERR ] %[2]s\n",
	}
	levels = make(map[string]*logger.LogLevel)
	mutex  sync.Mutex
)

func computeLevel(name string) logger.LogLevel {
	return defaultLevel()
}

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
