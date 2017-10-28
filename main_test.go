package log

import (
	"testing"
)

func TestDebug(*testing.T) {
	SetLevel("debug")
	Debug("example debug")
}

func TestInfo(*testing.T) {
	Info("example info")
}

func TestWarn(*testing.T) {
	Warn("example warning")
}
