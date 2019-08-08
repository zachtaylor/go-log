package log_test

import (
	"fmt"
	"os"
	"testing"

	"ztaylor.me/log"
)

func TestColors(*testing.T) {
	fmt.Println("-- Color Test --")
	log := log.StdOutService(log.LevelDebug)
	log.New().Debug("example debug")
	log.New().Info("example info")
	log.New().Warn("example warning")
	log.New().Error("example error")
	log.New().Trace("example trace")
	entry := log.New()
	entry.Message = "example unknown"
	log.Write(entry)
}

func TestNoColors(*testing.T) {
	fmt.Println("-- Non-Color Test --")
	log := log.NewService(log.LevelDebug, log.DefaultFormatter(false), os.Stdout)
	log.New().Debug("example debug")
	log.New().Info("example info")
	log.New().Warn("example warning")
	log.New().Error("example error")
	log.New().Trace("example trace")
	entry := log.New()
	entry.Message = "example unknown"
	log.Write(entry)
}
