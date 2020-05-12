package x

import "ztaylor.me/log"

func Hi(e *log.Entry) {
	func() {
		e.Info()
	}()
}
