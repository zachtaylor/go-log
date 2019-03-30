// +build windows

package log

import (
	"syscall"

	"github.com/konsorten/go-windows-terminal-sequences"
)

func init() {
	sequences.EnableVirtualTerminalProcessing(syscall.Stdout, true)
}
