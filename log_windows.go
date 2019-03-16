// +build windows

package log

import (
	"io"

	colorable "github.com/mattn/go-colorable"
)

func stdout() io.Writer {
	return colorable.NewColorableStdout()
}
