package console

import (
	logger ".."
	"../../cmd"
)

var overrides = make(map[string]int)

func computeLevel(name string) logger.LogLevel {
	if level, ok := overrides[name]; ok {
		return logger.LogLevel(level)
	}
	return defaultLevel()
}

func init() {
	cmd.GetStringMapInt("log.console.levelOverrides", &overrides)
}
