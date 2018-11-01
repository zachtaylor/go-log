package log

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"ztaylor.me/cast"
)

const (
	nocolor = "\x1b[0m"
	red     = "\x1b[31m"
	green   = "\x1b[32m"
	yellow  = "\x1b[33m"
	blue    = "\x1b[36m"
	gray    = "\x1b[37m"
)

type Formatter struct {
	Color bool
}

func (f *Formatter) Format(e *logrus.Entry) ([]byte, error) {
	var sb strings.Builder
	level, color := f.levelColor(e.Level)
	sb.WriteString(e.Time.Format("15:04:05"))
	if f.Color && color != "" {
		sb.WriteString(color)
	}
	sb.WriteString(level)
	sb.WriteString(fmt.Sprintf("%-42s", e.Message))
	if f.Color {
		sb.WriteString(nocolor)
	}
	for k, v := range e.Data {
		sb.WriteString(k + "=" + cast.String(v) + " ")
	}
	sb.WriteByte(10)
	return []byte(sb.String()), nil
}

func (f *Formatter) levelColor(level logrus.Level) (string, string) {
	switch level {
	case logrus.DebugLevel:
		return " D ", gray
	case logrus.InfoLevel:
		return " I ", blue
	case logrus.WarnLevel:
		return " W ", yellow
	case logrus.ErrorLevel:
		return " E ", red
	case logrus.PanicLevel:
		return " P ", red
	case logrus.FatalLevel:
		return " F ", red
	default:
		return " ? ", ""
	}
}
