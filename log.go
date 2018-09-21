package log

import (
	"github.com/sirupsen/logrus"
)

type Fields = logrus.Fields

type Log interface {
	Clone() Log
	Add(string, interface{}) Log
	WithFields(Fields) Log
	Protect(f func())
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
}

type log struct {
	*logrus.Entry
}

func wrapLog(entry *logrus.Entry) Log {
	return &log{entry}
}

func (log *log) Clone() Log {
	clone := wrapLogger(log.Entry.Logger).New()
	for key, val := range log.Entry.Data {
		clone.Add(key, val)
	}
	return clone
}

func (log *log) Add(k string, v interface{}) Log {
	log.Entry.Data[k] = v
	return log
}

func (log *log) Protect(f func()) {
	defer log.recoverPanic()
	f()
}

func (log *log) WithFields(f Fields) Log {
	log.Entry = log.Entry.WithFields(f)
	return log
}
