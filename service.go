package log

import "ztaylor.me/cast"

// Service provides logging functionality
type Service interface {
	cast.Closer
	// New creates a log
	New() *Entry
	// Format returns the internal Format
	Format() *Format
	// Write flushes a log message
	Write(cast.Time, *Source, Level, cast.JSON, string)
}

// NewService creates a log service with the minimum Level, format function and output dest
func NewService(level Level, f *Format, w cast.WriteCloser) Service {
	return &service{
		level: level,
		f:     f,
		w:     w,
	}
}

type service struct {
	level Level
	f     *Format
	w     cast.WriteCloser
}

func (svc *service) New() *Entry {
	return NewEntry(svc)
}

func (svc *service) Format() *Format {
	return svc.f
}

func (svc *service) Write(t cast.Time, src *Source, lvl Level, flds cast.JSON, msg string) {
	if lvl >= svc.level {
		svc.w.Write(svc.f.Format(t, src, lvl, flds, msg))
	}
}

func (svc *service) Close() error {
	return svc.w.Close()
}
