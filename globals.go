package log

var logger = NewLogger()

// New creates a Log from the global logger
func New() Log {
	return logger.New()
}

// Add is shorthand for New().Add(name, value)
func Add(name string, value interface{}) Log {
	return New().Add(name, value)
}

// Protect is shorthand for New().Protect(f)
func Protect(f func()) {
	New().Protect(f)
}

// Debug is shorthand for New().Debug(messages ...interface{})
func Debug(messages ...interface{}) {
	New().Debug(messages...)
}

// Info is shorthand for New().Info(messages ...interface{})
func Info(messages ...interface{}) {
	New().Info(messages...)
}

// Warn is shorthand for New().Warn(messages ...interface{})
func Warn(messages ...interface{}) {
	New().Warn(messages...)
}

// Error is shorthand for New().Error(messages ...interface{})
func Error(messages ...interface{}) {
	New().Error(messages...)
}

// WithFields is shorthand for New().WithFields(f)
func WithFields(f Fields) Log {
	return New().WithFields(f)
}

// SetLevel sets log level for the global logger
func SetLevel(level string) {
	logger.SetLevel(level)
}

// StartRoller sets log roller for the global logger
func StartRoller(path string) {
	logger.StartRoller(path)
}
