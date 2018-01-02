package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger() *Logger {
	logger := &Logger{logrus.New()}
	logger.logger.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "Jan 02 15:04:05",
	}
	return logger
}

func wrapLogger(logger *logrus.Logger) *Logger {
	return &Logger{logger}
}

func (logger *Logger) New() Log {
	return &log{logrus.NewEntry(logger.logger)}
}

func (logger *Logger) WithFields(fields Fields) Log {
	return &log{logger.logger.WithFields(logrus.Fields(fields))}
}

func (logger *Logger) StartRoller(path string) {
	go func() {
		for {
			file, err := os.OpenFile(path+time.Now().Format("20171225")+".log", os.O_CREATE|os.O_WRONLY, 0666)
			if err == nil {
				Info("proceeding on file")
			} else {
				Warn("failed to open log file")
			}
			logger.logger.Out = file
			<-time.After(24 * time.Hour)
			go file.Close()
		}
	}()
}

func (logger *Logger) SetFile(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		Info("log: proceeding on file")
	} else {
		Warn("failed to open log file")
	}
	logger.logger.Out = file
}

func (logger *Logger) SetStdOut() {
	logger.logger.Out = os.Stdout
}

func (logger *Logger) SetLevel(level string) {
	switch level {
	case "debug":
		logger.logger.Level = logrus.DebugLevel
		break
	case "info":
		logger.logger.Level = logrus.InfoLevel
		break
	case "warn":
		logger.logger.Level = logrus.WarnLevel
		break
	case "error":
		logger.logger.Level = logrus.ErrorLevel
		break
	case "fatal":
		logger.logger.Level = logrus.FatalLevel
		break
	case "panic":
		logger.logger.Level = logrus.PanicLevel
		break
	default:
		logger.logger.Level = logrus.InfoLevel
		Add("Level", level).Warn("log: level invalid")
		break
	}
}
