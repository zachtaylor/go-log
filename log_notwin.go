// +build !windows

package log

import (
	"io"
	"os"
)

func stdout() io.Writer {
	return os.Stdout
}
