package main

import (
	"fmt"
	"os"

	"ztaylor.me/cast"
	"ztaylor.me/log"
)

func main() {
	fmt.Println("-- Color Test --")
	logger := log.StdOutService(log.LevelDebug)
	logger.New().Debug("example debug")
	logger.New().Info("example info")
	logger.New().Warn("example warning")
	logger.New().Error("example error")
	logger.New().Trace("example trace")
	entry := logger.New()
	logger.Write(cast.Now(), entry, "example unknown")

	fmt.Println("-- Non-Color Test --")
	logger = log.NewService(log.LevelDebug, log.DefaultFormatter(false), os.Stdout)
	logger.New().Debug("example debug")
	logger.New().Info("example info")
	logger.New().Warn("example warning")
	logger.New().Error("example error")
	logger.New().Trace("example trace")
	entry = logger.New()
	logger.Write(cast.Now(), entry, "example unknown")

	fmt.Println("-- Tag/Message Test --")
	logger = log.StdOutService(log.LevelInfo)
	logger.New().Info()
	logger.New().Tag("tag").Info()
	logger.New().Info("message")
	logger.New().Source().Info()
	logger.New().Tag("tag").Info("message")
	logger.New().Source().Tag("tag").Info()
	logger.New().Source().Tag("tag").Info("message")
}
