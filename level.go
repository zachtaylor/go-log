package log

// Level is used to rank the importance of logs
type Level uint8

const (
	// LevelDebug is the lowest level
	LevelDebug = iota
	// LevelInfo is the default level
	LevelInfo
	// LevelWarn is a raised level
	LevelWarn
	// LevelError is the considered the top level
	LevelError
	// LevelTrace is a sentinal value that is guaranteed to print
	LevelTrace
)
