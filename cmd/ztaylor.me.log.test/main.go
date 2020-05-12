package main

import (
	"fmt"

	"ztaylor.me/log"
	"ztaylor.me/log/cmd/ztaylor.me.log.test/x"
	"ztaylor.me/log/cmd/ztaylor.me.log.test/x/c4"
	"ztaylor.me/log/cmd/ztaylor.me.log.test/xyz"
)

func main() {
	fmt.Println("-- Color Test --")
	logger := log.StdOutService(log.LevelTrace, log.DefaultFormatWithColor())
	logger.New().Trace("example trace")
	logger.New().Debug("example debug")
	logger.New().Info("example info")
	logger.New().Warn("example warning")
	logger.New().Error("example error")
	logger.New().Out("example out")

	fmt.Println("-- Non-Color Test --")
	logger = log.StdOutService(log.LevelTrace, log.DefaultFormatWithoutColor())
	logger.New().Trace("example trace")
	logger.New().Debug("example debug")
	logger.New().Info("example info")
	logger.New().Warn("example warning")
	logger.New().Error("example error")
	logger.New().Out("example out")

	fmt.Println("-- Tag/Message/Source Test --")
	logger = log.StdOutService(log.LevelTrace, log.DefaultFormatWithColor())
	logger.New().Info()

	logger.New().Tag("tag").Info()
	logger.New().Info("message")
	logger.New().Tag("tag").Info("message")
	logger.New().Tag("tag").Tag("tag").Info()
	logger.New().Tag("tag").Tag("tag").Info("message")

	logger = log.StdOutService(log.LevelTrace, log.DefaultFormatWithColor())
	fmt.Println(`-- CutSourcePackage Test --`)
	logger.New().Debug()
	c4.Hi(logger.New())
	x.Hi(logger.New())
	xyz.Hi(logger.New())
	x.CutSourcePath(logger)
	fmt.Println(`/x.CutSourcePath(logger)`)
	logger.New().Debug()
	c4.Hi(logger.New())
	x.Hi(logger.New())
	xyz.Hi(logger.New())
}
