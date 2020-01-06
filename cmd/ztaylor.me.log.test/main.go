package main

import (
	"fmt"
	"os"

	"ztaylor.me/cast"
	"ztaylor.me/log"
	"ztaylor.me/log/cmd/ztaylor.me.log.test/x"
	"ztaylor.me/log/cmd/ztaylor.me.log.test/x/c4"
	"ztaylor.me/log/cmd/ztaylor.me.log.test/xyz"
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
	logger = log.NewService(log.LevelDebug, log.DefaultFormatWithoutColor(), os.Stdout)
	logger.New().Debug("example debug")
	logger.New().Info("example info")
	logger.New().Warn("example warning")
	logger.New().Error("example error")
	logger.New().Trace("example trace")
	entry = logger.New()
	logger.Write(cast.Now(), entry, "example unknown")

	fmt.Println("-- Tag/Message/Source Test --")
	logger = log.StdOutService(log.LevelInfo)
	logger.New().Info()
	logger.New().Tag("tag").Info()
	logger.New().Info("message")
	logger.New().Tag("tag").Info("message")
	logger.New().Tag("tag").Tag("tag").Info()
	logger.New().Tag("tag").Tag("tag").Info("message")

	logger = log.StdOutService(log.LevelDebug)
	fmt.Println(`-- CutSourcePackage Test --`)
	logger.New().Source().Debug()
	c4.Hi(logger.New())
	x.Hi(logger.New())
	xyz.Hi(logger.New())
	logger.Formatter().CutSourcePath(1)
	fmt.Println(`logger.Formatter().CutSourcePath(1)`)
	logger.New().Source().Debug()
	c4.Hi(logger.New())
	x.Hi(logger.New())
	xyz.Hi(logger.New())
	logger.Formatter().CutSourcePath(0)
	fmt.Println(`logger.Formatter().CutSourcePath(0)`)
	logger.New().Source().Debug()
	c4.Hi(logger.New())
	x.Hi(logger.New())
	xyz.Hi(logger.New())
	x.CutSourcePath(logger)
	fmt.Println(`/x.CutSourcePath(logger)`)
	logger.New().Source().Debug()
	c4.Hi(logger.New())
	x.Hi(logger.New())
	xyz.Hi(logger.New())
}
