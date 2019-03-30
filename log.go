package log

// Fields is log data
type Fields map[string]interface{}

// DailyRollingService creates a log service that writes a rotating log file, named by the day
func DailyRollingService(level Level, path string) Service {
	return NewService(level, DefaultFormatter(false), NewRoller(path))
}

// StdOutService creates a log service that wraps std out
func StdOutService(level Level) Service {
	return NewService(level, DefaultFormatter(true), stdout())
}
