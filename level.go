package log

import "ztaylor.me/cast"

// Level is used to rank the importance of logs
type Level uint8

const (
	// LevelDebug is the lowest level
	LevelDebug = iota + 1
	// LevelInfo is the default level
	LevelInfo
	// LevelWarn is a raised level
	LevelWarn
	// LevelError is the considered the top level
	LevelError
	// LevelTrace is a sentinal value that is guaranteed to print
	LevelTrace
)

// GetLevel returns the level named, if valid
func GetLevel(level string) (Level, error) {
	switch level {
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warn":
		return LevelWarn, nil
	case "error":
		return LevelError, nil
	case "trace":
		return LevelTrace, nil
	default:
		return LevelDebug, cast.NewError(nil, "log level unknown: "+level)
	}
}

// ByteCode returns an ASCII byte code for this level
func (level Level) ByteCode() byte {
	switch level {
	case LevelDebug:
		return 68 // D
	case LevelInfo:
		return 73 // I
	case LevelWarn:
		return 87 // W
	case LevelError:
		return 69 // E
	case LevelTrace:
		return 84 // T
	default:
		return 63 // ?
	}
}
