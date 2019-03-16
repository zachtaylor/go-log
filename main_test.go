package log_test

import (
	"testing"

	"ztaylor.me/log"
)

func TestDebug(*testing.T) {
	log := log.StdOutService(log.LevelDebug)
	log.New().Debug("example debug")
}

func TestInfo(*testing.T) {
	log := log.StdOutService(log.LevelDebug)
	log.New().Info("example info")
}

func TestWarn(*testing.T) {
	log := log.StdOutService(log.LevelDebug)
	log.New().Warn("example warning")
}

func TestError(*testing.T) {
	log := log.StdOutService(log.LevelDebug)
	log.New().Error("example error")
}

func TestTrace(*testing.T) {
	log := log.StdOutService(log.LevelDebug)
	log.New().Trace("example trace")
}
