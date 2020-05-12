package log

import (
	"os"
	"time"

	"ztaylor.me/cast"
)

// DailyRollingService creates a Service that writes a rotating log file, named by the day
func DailyRollingService(level Level, f *Format, path string) Service {
	return NewService(level, f, NewRoller(path))
}

// StdOutService creates a log service that wraps std out
func StdOutService(level Level, f *Format) Service {
	return NewService(level, f, os.Stdout)
}

// DefaultFormatWithColor creates a new Format with colors, and default time and message formats
func DefaultFormatWithColor() *Format {
	f := NewFormatWithDefaultColor(DefaultTimeFormat, DefaultSourceFormat, DefaultMessageFormat)
	f.CutPathWith(NewSource(1), 0)
	return f
}

// DefaultFormatWithoutColor creates a new Format without colors, and default time and message formats
func DefaultFormatWithoutColor() *Format {
	f := NewFormatWithoutColor(DefaultTimeFormat, DefaultSourceFormat, DefaultMessageFormat)
	f.CutPathWith(NewSource(1), 0)
	return f
}

// NewFormatWithoutColor creates a new Format without colors, and specified time and message formats
func NewFormatWithoutColor(tfmt func(time.Time) string, srcfmt, mfmt func(string) string) *Format {
	return NewFormat(tfmt, srcfmt, mfmt, nil)
}

// NewFormatWithDefaultColor creates a new Format with default colors, and specified time and message formats
func NewFormatWithDefaultColor(tfmt func(time.Time) string, srcfmt, mfmt func(string) string) *Format {
	return NewFormat(tfmt, srcfmt, mfmt, DefaultColors())
}

// DefaultTimeFormat returns "15:04:05" (24-hour format)
func DefaultTimeFormat(time time.Time) string {
	return time.Format("15:04:05")
}

// DefaultSourceFormat formats a string to max length 24, elipses the beginning
func DefaultSourceFormat(src string) string {
	const maxlen = 24
	lensrc := len(src)
	if lensrc <= maxlen {
		return src
	}
	lendif := lensrc - maxlen
	buf := make([]byte, maxlen)
	buf[0] = '.'
	buf[1] = '.'
	buf[2] = '.'
	for i := 3; i < maxlen; i++ {
		buf[i] = src[lendif+i]
	}
	return cast.StringBytes(buf)
}

// DefaultMessageFormat formats a string to "%-25s " (minimum length 26 right-padded and last char is space)
func DefaultMessageFormat(msg string) string {
	return cast.Sprintf("%-25s ", msg)
}

// DefaultColors returns the default color set
func DefaultColors() map[Level]string {
	return map[Level]string{
		LevelTrace: purple,
		LevelDebug: green,
		LevelInfo:  blue,
		LevelWarn:  yellow,
		LevelError: red,
		LevelOut:   white,
	}
}
