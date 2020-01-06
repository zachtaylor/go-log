package log

import "os"

// DailyRollingService creates a log service that writes a rotating log file, named by the day
func DailyRollingService(level Level, path string) Service {
	return NewService(level, DefaultFormatWithoutColor(), NewRoller(path))
}

// StdOutService creates a log service that wraps std out
func StdOutService(level Level) Service {
	return NewService(level, DefaultFormatWithColor(), os.Stdout)
}
