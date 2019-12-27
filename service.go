package log

import "ztaylor.me/cast"

// Service provides logging functionality
type Service interface {
	cast.Closer
	// New creates a log
	New() *Entry
	// Write flushes a log message
	Write(cast.Time, *Entry, ...interface{})
}

// NewService creates a log service with the minimum Level, format function and output dest
func NewService(level Level, f Formatter, w cast.WriteCloser) Service {
	return &service{
		level: level,
		f:     f,
		w:     w,
	}
}

type service struct {
	level Level
	f     Formatter
	w     cast.WriteCloser
}

func (svc *service) New() *Entry {
	return NewEntry(svc)
}

func (svc *service) Write(t cast.Time, log *Entry, args ...interface{}) {
	if log.Level < LevelDebug || log.Level >= svc.level {
		svc.w.Write(svc.f.Format(t, log, args...))
	}
}

func (svc *service) Close() error {
	return svc.w.Close()
}
