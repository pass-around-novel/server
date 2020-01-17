package console

import (
	logger ".."
	"../../cmd"
)

var configDefault int

func defaultLevel() logger.LogLevel {
	return logger.LogLevel(configDefault) + logger.LogLevel(cmd.Quiet-cmd.Verbose)
}

func init() {
	cmd.GetInt("log.console.defaultLevel", &configDefault)
}
