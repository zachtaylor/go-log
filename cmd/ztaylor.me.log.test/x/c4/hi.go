package c4

import "ztaylor.me/log"

func Hi(e *log.Entry) {
	e.Source().Trace()
}
