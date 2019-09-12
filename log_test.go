package log_test

import (
	"fmt"
	"os"
	"testing"
	"time"

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
	log.Write(time.Now(), entry, "example unknown")
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
	log.Write(time.Now(), entry, "example unknown")
}

func TestTagMessage(*testing.T) {
	fmt.Println("-- Tag/Message Test --")
	log := log.StdOutService(log.LevelInfo)
	log.New().Tag("tag").Info()
	log.New().Info("message")
	log.New().Tag("tag1").Tag("tag2").Info()
	log.New().Tag("tag").Info("message")
	log.New().Tag("tag1").Tag("tag2").Info("message")
	log.New().Tag("tag1").Tag("tag2").Info("message1", "message2")
	log.New().Tag("tag1").Tag("tag2").Tag("tag3").Info()
	log.New().Tag("tag1").Tag("tag2").Tag("tag3").Info("message1", "message2", "message3")
}

func TestMessageCast(*testing.T) {

}
