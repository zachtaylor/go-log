package log

import (
	"github.com/sirupsen/logrus"
)

type Log interface {
	Clone() Log
	Add(string, interface{}) Log
	WithFields(map[string]interface{}) Log
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

func (log *log) WithFields(m map[string]interface{}) Log {
	log.Entry = log.Entry.WithFields(m)
	return log
}
