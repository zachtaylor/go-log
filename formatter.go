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
	Format(time.Time, *Entry, ...interface{}) []byte
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

func (f *format) Format(time time.Time, e *Entry, args ...interface{}) []byte {
	var sb strings.Builder
	sb.WriteString(time.Format(f.TimeFormat))
	sb.WriteByte(32) // space
	if f.Colors != nil {
		sb.WriteString(f.Colors[e.Level])
	} else {
		sb.WriteByte(e.Level.ByteCode())
		sb.WriteByte(32) // space
	}
	if msg := buildmessage(args...); e.Prefix == "" {
		fmt.Fprintf(&sb, f.MessageFormat, msg)
	} else if msg == "" {
		fmt.Fprintf(&sb, f.MessageFormat, e.Prefix)
	} else {
		fmt.Fprintf(&sb, f.MessageFormat, e.Prefix+": "+msg)
	}
	if f.Colors != nil {
		sb.WriteString(nocolor)
	}
	for _, k := range e.Fields.SortKeys() {
		fmt.Fprintf(&sb, "%s=%v ", k, e.Fields[k])
	}
	sb.WriteByte(10) // newline
	return []byte(sb.String())
}

func buildmessage(args ...interface{}) string {
	var msg string // build message
	if len(args) > 0 {
		for i, arg := range args {
			if i > 0 {
				msg += " "
			}
			msg += fmt.Sprintf("%v", arg)
		}
	}
	return msg
}
