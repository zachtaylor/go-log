package log

import (
	"runtime"
	"strings"
)

type Source struct {
	Pkg  string
	F    string
	Line int
}

// NewSource creates a Source reference to the caller # number than default (0 is default)
func NewSource(history int) *Source {
	if history < 0 {
		return nil
	}
	pc, _, line, _ := runtime.Caller(1 + history)
	// pc, file, line, _ := runtime.Caller(1 + history)
	// _, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return &Source{
		Pkg:  packageName,
		F:    funcName,
		Line: line,
	}
}
