package log

import (
	"runtime"

	"ztaylor.me/cast"
)

// Source links to a go code instruction
type Source struct {
	file string
	line int
}

// NewSource creates a Source reference to the caller number # higher in the stack frame than the invoking call (0 is default)
func NewSource(history int) *Source {
	if history < 0 {
		return nil
	}
	_, file, line, _ := runtime.Caller(1 + history)

	// saved
	// pc, file, line, _ := runtime.Caller(1 + history)
	// _, fileName := path.Split(file)
	// parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	// pl := len(parts)
	// packageName := ""
	// funcName := parts[pl-1]
	// if parts[pl-2][0] == '(' {
	// funcName = parts[pl-2] + "." + funcName
	// packageName = strings.Join(parts[0:pl-2], ".")
	// } else {
	// packageName = strings.Join(parts[0:pl-1], ".")
	// }

	return &Source{
		// pkg: packageName,
		file: file,
		// fxn: funcName,
		line: line,
	}
}

// File returns Source.File
func (src *Source) File() string {
	return src.file
}

// Line returns Source.LineNo
func (src *Source) Line() int {
	return src.line
}

func (src *Source) String() string {
	msg := ""
	if src != nil {
		if file := src.File(); file != "" {
			if flen := len(file); flen > 4 && file[flen-3:] == ".go" {
				msg = file[:flen-3]
			} else {
				msg = file
			}
		}
		if lno := src.Line(); lno > 0 {
			msg += "#" + cast.StringI(lno)
		}
	}
	return msg
}
