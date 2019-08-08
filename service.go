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
	Write(*Entry)
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

func (svc *service) Write(log *Entry) {
	if log.Level < LevelDebug || log.Level >= svc.level {
		svc.w.Write(svc.f.Format(time.Now(), log))
	}
}

func (svc *service) Close() error {
	return svc.w.Close()
}
