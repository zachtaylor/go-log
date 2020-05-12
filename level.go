package log

import "ztaylor.me/cast"

// Level is used to rank the importance of logs
type Level uint8

const (
	// LevelTrace is the lowest level
	LevelTrace = iota
	// LevelDebug is a more detailed value
	LevelDebug
	// LevelInfo is the default level
	LevelInfo
	// LevelWarn is a raised level
	LevelWarn
	// LevelError is the considered the top level
	LevelError
	// LevelOut is the highest value, a sentinal value
	LevelOut
)

// GetLevel returns the level named, if valid
func GetLevel(level string) (Level, error) {
	switch level {
	case "trace":
		return LevelTrace, nil
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warn":
		return LevelWarn, nil
	case "error":
		return LevelError, nil
	case "out":
		return LevelOut, nil
	default:
		return LevelDebug, cast.NewError(nil, "log level unknown: "+level)
	}
}

// ByteCode returns an ASCII byte code for this level
func (level Level) ByteCode() byte {
	switch level {
	case LevelTrace:
		return 84 // T
	case LevelDebug:
		return 68 // D
	case LevelInfo:
		return 73 // I
	case LevelWarn:
		return 87 // W
	case LevelError:
		return 69 // E
	case LevelOut:
		return 79 // O
	default:
		return 63 // ?
	}
}
