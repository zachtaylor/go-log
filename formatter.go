package log

import (
	"fmt"
	"strings"
	"time"
)

const (
	nocolor = "\x1b[0m"
	red     = "\x1b[31m"
	green   = "\x1b[32m"
	yellow  = "\x1b[33m"
	purple  = "\x1b[35m"
	cyan    = "\x1b[36m"
)

// Formatter encodes a log
type Formatter interface {
	// Format creates writable output
	Format(time.Time, *Entry) []byte
}

// DefaultFormatter creates a basic Formatter with color printing on or off
func DefaultFormatter(color bool) Formatter {
	var colors map[Level]string
	if color {
		colors = map[Level]string{
			LevelDebug: green,
			LevelInfo:  cyan,
			LevelWarn:  yellow,
			LevelError: red,
			LevelTrace: purple,
		}
	}
	return NewFormatter("15:04:05", "%-42s", colors)
}

// NewFormatter creates a log Formatter, using time.Format, fmt formatting, and
// optional color mode(s)
func NewFormatter(timeFormat, msgFormat string, colors map[Level]string) Formatter {
	return &format{
		TimeFormat:    timeFormat,
		MessageFormat: msgFormat,
		Colors:        colors,
	}
}

type format struct {
	TimeFormat    string
	MessageFormat string
	Colors        map[Level]string
}

func (f *format) Format(time time.Time, e *Entry) []byte {
	var sb strings.Builder
	sb.WriteString(time.Format(f.TimeFormat))
	sb.WriteByte(32)
	if f.Colors != nil {
		sb.WriteString(f.Colors[e.Level])
	} else {
		sb.WriteString(f.levelString(e.Level))
	}
	fmt.Fprintf(&sb, f.MessageFormat, e.Prefix+e.Message)
	if f.Colors != nil {
		sb.WriteString(nocolor)
	}
	for k, v := range e.Fields {
		fmt.Fprintf(&sb, "%s=%v ", k, v)
	}
	sb.WriteByte(10)
	return []byte(sb.String())
}

func (f *format) levelString(level Level) string {
	switch level {
	case LevelDebug:
		return "D "
	case LevelInfo:
		return "I "
	case LevelWarn:
		return "W "
	case LevelError:
		return "E "
	case LevelTrace:
		return "T "
	default:
		return "? "
	}
}
