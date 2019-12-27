package log

import "ztaylor.me/cast"

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
	Format(cast.Time, *Entry, ...interface{}) []byte
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

func (f *format) Format(time cast.Time, e *Entry, args ...interface{}) []byte {
	var sb cast.StringBuilder
	sb.WriteString(time.Format(f.TimeFormat))
	sb.WriteByte(32) // space
	if f.Colors != nil {
		sb.WriteString(f.Colors[e.Level])
	} else {
		sb.WriteByte(e.Level.ByteCode())
		sb.WriteByte(32) // space
	}
	cast.Fprintf(&sb, f.MessageFormat, buildmessage(e, args...))
	if f.Colors != nil {
		sb.WriteString(nocolor)
	}
	for _, k := range e.Fields.GetKeys() {
		cast.Fprintf(&sb, "%s=%v ", k, e.Fields[k])
	}
	sb.WriteByte(10) // newline
	return []byte(sb.String())
}

func buildmessage(e *Entry, args ...interface{}) string {
	msg := ""
	if src := e.GetSource(); src != nil {
		if src.Pkg != "" {
			msg = src.Pkg
		}
		if src.F != "" {
			if len(msg) > 0 {
				msg += "/"
			}
			msg += src.F
		}
		if src.Line > 0 {
			msg += "#" + cast.StringI(src.Line)
		}
	}
	for _, tag := range e.Tags {
		if len(msg) > 0 {
			msg += ": "
		}
		msg += tag
	}
	if len(msg) > 0 && len(args) > 0 {
		msg += ": "
	}
	return msg + cast.StringN(args...)
}
