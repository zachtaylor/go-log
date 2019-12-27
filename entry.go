package log

import "ztaylor.me/cast"

// NewEntry creates a new Entry
func NewEntry(service Service) *Entry {
	return &Entry{
		Service: service,
		Tags:    make([]string, 0),
		Fields:  make(cast.JSON),
	}
}

// Entry is an intermediate step in creating a log
type Entry struct {
	Service
	src  *Source
	Tags []string
	Level
	Fields cast.JSON
}

// Add writes any value to the Fields
func (log *Entry) Add(k string, v interface{}) *Entry {
	log.Fields[k] = v
	return log
}

// GetSource returns the Source if available
func (log *Entry) GetSource() *Source {
	return log.src
}

// Source sets the src to the moment Source is called
func (log *Entry) Source() *Entry {
	log.src = NewSource(1)
	return log
}

// Tag appends any number of tags to the entry
func (log *Entry) Tag(s ...string) *Entry {
	log.Tags = append(log.Tags, s...)
	return log
}

// With writes any value to the Fields
func (log *Entry) With(fields cast.JSON) *Entry {
	for k, v := range fields {
		log.Fields[k] = v
	}
	return log
}

// Copy duplicates the Entry
func (log *Entry) Copy() *Entry {
	fields := make(cast.JSON)
	for k, v := range log.Fields {
		fields[k] = v
	}
	tags := make([]string, 0, len(log.Tags))
	for _, tag := range log.Tags {
		tags = append(tags, tag)
	}
	return &Entry{
		Service: log.Service,
		Fields:  fields,
		Level:   log.Level,
		Tags:    tags,
	}
}

// Debug calls Write with LevelDebug
func (log *Entry) Debug(args ...interface{}) {
	log.write(LevelDebug, args...)
}

// Info calls Write with LevelInfo
func (log *Entry) Info(args ...interface{}) {
	log.write(LevelInfo, args...)
}

// Warn calls Write with LevelWarn
func (log *Entry) Warn(args ...interface{}) {
	log.write(LevelWarn, args...)
}

// Warn calls Write with LevelWarn
func (log *Entry) Error(args ...interface{}) {
	log.write(LevelError, args...)
}

// Trace calls Write with LevelTrace
func (log *Entry) Trace(args ...interface{}) {
	log.write(LevelTrace, args...)
}

func (log *Entry) write(lvl Level, args ...interface{}) {
	log.Level = lvl
	log.Service.Write(cast.Now(), log, args...)
}
