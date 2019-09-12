package log

import (
	"io"
	"time"
)

// Service provides logging functionality
type Service interface {
	io.Closer
	// New creates a log
	New() *Entry
	// Write flushes a log message
	Write(time.Time, *Entry, ...interface{})
}

// NewService creates a log service with the minimum Level, format function and output dest
func NewService(level Level, f Formatter, w io.WriteCloser) Service {
	return &service{
		level: level,
		f:     f,
		w:     w,
	}
}

type service struct {
	level Level
	f     Formatter
	w     io.WriteCloser
}

func (svc *service) New() *Entry {
	return NewEntry(svc)
}

func (svc *service) Write(t time.Time, log *Entry, args ...interface{}) {
	if log.Level < LevelDebug || log.Level >= svc.level {
		svc.w.Write(svc.f.Format(t, log, args...))
	}
}

func (svc *service) Close() error {
	return svc.w.Close()
}
