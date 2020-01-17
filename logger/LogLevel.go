package logger

// LogLevel describes how important a log message is
type LogLevel int

const (
	// Debug contains messages intended for debugging only
	Debug LogLevel = 0
	// Info contains informational messages for the server administrator
	Info LogLevel = 1
	// Warning contains messages that may pose problems to the server
	Warning LogLevel = 2
	// Error contains messages describing when the server cannot operate
	Error LogLevel = 3
	// None is never logged to and it means all logs are discarded
	None LogLevel = 4
)
