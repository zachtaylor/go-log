package log

import "ztaylor.me/cast"

const (
	nocolor = "\x1b[0m"
	red     = "\x1b[31m"
	green   = "\x1b[32m"
	yellow  = "\x1b[33m"
	purple  = "\x1b[35m"
	cyan    = "\x1b[36m"
)

// Formatting supports options for Formatter
type Formatting struct {
	// TimeFmt is the time format
	TimeFmt string
	// MsgFmt is the message format
	MsgFmt string
	// colors is the colors to use per level, or nil when in no-color mode
	Colors map[Level]string
	// PathCut
	PathCut []string
}

// CutSourcePath implements Formatter
func (f *Formatting) CutSourcePath(parentno int) {
	source := NewSource(1)
	filePath, _ := cast.SplitPath(source.File())
	for i := 0; cast.Contains(filePath, "/") && i <= parentno; i++ {
		filePath = filePath[:cast.LastIndex(filePath, "/")]
	}
	f.PathCut = append(f.PathCut, filePath)
	cast.SortSlice(f.PathCut, func(i, j int) bool {
		return f.PathCut[i] > f.PathCut[j]
	})
}

// Format implements Formatter
func (f *Formatting) Format(time cast.Time, e *Entry, args ...interface{}) []byte {
	var sb cast.StringBuilder
	sb.WriteString(time.Format(f.TimeFmt))
	sb.WriteByte(32) // space
	if f.Colors != nil {
		sb.WriteString(f.Colors[e.Level])
	} else {
		sb.WriteByte(e.Level.ByteCode())
		sb.WriteByte(32) // space
	}
	cast.Fprintf(&sb, f.MsgFmt, buildmessage(e, f, args...))
	if f.Colors != nil {
		sb.WriteString(nocolor)
	}
	for _, k := range e.Fields.GetKeys() {
		cast.Fprintf(&sb, "%s=%v ", k, e.Fields[k])
	}
	sb.WriteByte(10) // newline
	return []byte(sb.String())
}

func buildmessage(e *Entry, f *Formatting, args ...interface{}) string {
	msg := ""
	if src := e.GetSource(); src != nil {
		if file := src.File(); file != "" {
			if len(msg) > 0 {
				msg += "/"
			}
			msg += file
		}
		for _, path := range f.PathCut {
			if pkglen := len(path); pkglen > len(msg)+1 {
			} else if prefix := msg[:pkglen]; prefix != path {
			} else if len(msg) > pkglen {
				if msg[pkglen] == '/' {
					msg = msg[pkglen+1:]
				}
			} else {
				msg = ""
			}
		}
		if lno := src.Line(); lno > 0 {
			msg += "#" + cast.StringI(lno)
		}
	}
	for _, tag := range e.Tags {
		if len(msg) > 0 {
			msg += ": "
		}
		msg += tag
	}
	if len(msg) > 0 && len(args) > 0 {
		msg += ": "
	}
	return msg + cast.StringN(args...)
}
