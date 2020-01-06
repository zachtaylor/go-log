package xyz

import "ztaylor.me/log"

func Hi(e *log.Entry) {
	e.Source().Warn()
}
