package log

import "ztaylor.me/cast"

// Formatter encodes a log
type Formatter interface {
	// CutSourcePath adds the calling source's path ancestor number to the path cutset
	CutSourcePath(int)
	// Format creates writable output
	Format(cast.Time, *Entry, ...interface{}) []byte
}

// NewFormat creates Formatter
func NewFormat(tfmt, mfmt string, colors map[Level]string, pkgcut []string) *Formatting {
	if pkgcut == nil {
		pkgcut = make([]string, 0)
	}
	return &Formatting{
		TimeFmt: tfmt,
		MsgFmt:  mfmt,
		Colors:  colors,
		PathCut: pkgcut,
	}
}

// DefaultFormatWithoutColor uses NewFormatter with no colors
func DefaultFormatWithoutColor() *Formatting {
	return NewFormatWithoutColor(DefaultTimeFormat(), DefaultMessageFormat())
}

// DefaultFormatWithColor uses NewFormatWithDefaultColor with DefaultTimeFormat and DefaultMessageFormat
func DefaultFormatWithColor() *Formatting {
	return NewFormatWithDefaultColor(DefaultTimeFormat(), DefaultMessageFormat())
}

// NewFormatWithoutColor uses NewFormatter with no colors
func NewFormatWithoutColor(tfmt, mfmt string) *Formatting {
	return NewFormat(tfmt, mfmt, nil, nil)
}

// NewFormatWithDefaultColor uses NewFormatter with DefaultColors
func NewFormatWithDefaultColor(tfmt, mfmt string) *Formatting {
	return NewFormat(tfmt, mfmt, DefaultColors(), nil)
}

// DefaultTimeFormat returns "15:04:05" (24-hour format)
func DefaultTimeFormat() string {
	return "15:04:05"
}

// DefaultMessageFormat returns "%-63s " (min 64-char right-padded with last char is space)
func DefaultMessageFormat() string {
	return "%-63s "
}

// DefaultColors returns the default color set
func DefaultColors() map[Level]string {
	return map[Level]string{
		LevelDebug: green,
		LevelInfo:  cyan,
		LevelWarn:  yellow,
		LevelError: red,
		LevelTrace: purple,
	}
}
