package console

import (
	logger ".."
	"../../cmd"
)

func defaultLevel() logger.LogLevel {
	return logger.Info + logger.LogLevel(cmd.Quiet-cmd.Verbose)
}
