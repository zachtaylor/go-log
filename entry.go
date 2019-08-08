package log

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
	Prefix  string
	Message string
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
		log.Prefix = tag + ": "
	} else {
		log.Prefix += tag + ": "
	}
	return log
}

// Debug calls Write with LevelDebug
func (log *Entry) Debug(v string) {
	log.Level = LevelDebug
	log.write(v)
}

// Info calls Write with LevelInfo
func (log *Entry) Info(v string) {
	log.Level = LevelInfo
	log.write(v)
}

// Warn calls Write with LevelWarn
func (log *Entry) Warn(v string) {
	log.Level = LevelWarn
	log.write(v)
}

// Warn calls Write with LevelWarn
func (log *Entry) Error(v string) {
	log.Level = LevelError
	log.write(v)
}

// Trace calls Write with LevelTrace
func (log *Entry) Trace(v string) {
	log.Level = LevelTrace
	log.write(v)
}

func (log *Entry) write(v string) {
	log.Message = v
	log.Service.Write(log)
}

// Protect calls a func and writes an error if the func causes a panic
//
// Adds "Error", "Source" values to log
func (log *Entry) Protect(f func()) {
	defer func() {
		if err := recover().(error); err != nil {
			log.With(Fields{
				"Source": identifyPanic(),
				"Error":  err,
			}).Error("panic stopped")
		}
	}()
	f()
}
