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
	Service Service
	Level   Level
	Tags    []string
	Fields  cast.JSON
}

// Add writes any value to the Fields
func (log *Entry) Add(k string, v interface{}) *Entry {
	log.Fields[k] = v
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
	tags := make([]string, len(log.Tags))
	for i, tag := range log.Tags {
		tags[i] = tag
	}
	return &Entry{
		Service: log.Service,
		Fields:  fields,
		Level:   log.Level,
		Tags:    tags,
	}
}

// Trace calls Write with LevelTrace
func (log *Entry) Trace(args ...interface{}) {
	log.Service.Write(cast.Now(), NewSource(1), LevelTrace, log.Fields, parseargs(log.Tags, args...))
}

// Debug calls Write with LevelDebug
func (log *Entry) Debug(args ...interface{}) {
	log.Service.Write(cast.Now(), NewSource(1), LevelDebug, log.Fields, parseargs(log.Tags, args...))
}

// Info calls Write with LevelInfo
func (log *Entry) Info(args ...interface{}) {
	log.Service.Write(cast.Now(), NewSource(1), LevelInfo, log.Fields, parseargs(log.Tags, args...))
}

// Warn calls Write with LevelWarn
func (log *Entry) Warn(args ...interface{}) {
	log.Service.Write(cast.Now(), NewSource(1), LevelWarn, log.Fields, parseargs(log.Tags, args...))
}

// Warn calls Write with LevelWarn
func (log *Entry) Error(args ...interface{}) {
	log.Service.Write(cast.Now(), NewSource(1), LevelError, log.Fields, parseargs(log.Tags, args...))
}

// Out calls Write with LevelOut
func (log *Entry) Out(args ...interface{}) {
	log.Service.Write(cast.Now(), NewSource(1), LevelOut, log.Fields, parseargs(log.Tags, args...))
}

func parseargs(tags []string, args ...interface{}) string {
	sargs := make([]string, len(args))
	for i, arg := range args {
		sargs[i] = cast.String(arg)
	}
	tags = append(tags, sargs...)
	var sb cast.StringBuilder
	for _, arg := range tags {
		if sb.Len() > 0 {
			sb.WriteString(": ")
		}
		sb.WriteString(arg)
	}
	return sb.String()
}
