package log

var logger = NewLogger()

func New() Log {
	return logger.New()
}

func Add(name string, value interface{}) Log {
	return New().Add(name, value)
}

func Debug(message string) {
	New().Debug(message)
}

func Info(message string) {
	New().Info(message)
}

func Warn(message string) {
	New().Warn(message)
}

func Error(message interface{}) {
	New().Error(message)
}

func SetLevel(level string) {
	logger.SetLevel(level)
}

func StartRoller(path string) {
	logger.StartRoller(path)
}
