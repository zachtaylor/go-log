package log

import "time"

// NewEntry creates a new Entry
func NewEntry(service Service) *Entry {
	return &Entry{
		Service: service,
		Fields:  make(Fields),
	}
}

// Entry is an intermediate step in creating a log
type Entry struct {
	Service
	Fields
	Level
	Prefix string
}

// Add writes any value to the Fields
func (log *Entry) Add(k string, v interface{}) *Entry {
	log.Fields[k] = v
	return log
}

// With writes any value to the Fields
func (log *Entry) With(fields Fields) *Entry {
	for k, v := range fields {
		log.Fields[k] = v
	}
	return log
}

// Copy duplicates the Entry
func (log *Entry) Copy() *Entry {
	fields := make(Fields)
	for k, v := range log.Fields {
		fields[k] = v
	}
	return &Entry{
		Service: log.Service,
		Fields:  fields,
		Level:   log.Level,
		Prefix:  log.Prefix,
	}
}

// Tag sets the message prefix
func (log *Entry) Tag(tag string) *Entry {
	if log.Prefix == "" {
		log.Prefix = tag
	} else {
		log.Prefix += ": " + tag
	}
	return log
}

// Debug calls Write with LevelDebug
func (log *Entry) Debug(args ...interface{}) {
	log.Level = LevelDebug
	log.write(args...)
}

// Info calls Write with LevelInfo
func (log *Entry) Info(args ...interface{}) {
	log.Level = LevelInfo
	log.write(args...)
}

// Warn calls Write with LevelWarn
func (log *Entry) Warn(args ...interface{}) {
	log.Level = LevelWarn
	log.write(args...)
}

// Warn calls Write with LevelWarn
func (log *Entry) Error(args ...interface{}) {
	log.Level = LevelError
	log.write(args...)
}

// Trace calls Write with LevelTrace
func (log *Entry) Trace(args ...interface{}) {
	log.Level = LevelTrace
	log.write(args...)
}

func (log *Entry) write(args ...interface{}) {
	log.Service.Write(time.Now(), log, args...)
}
