package x

import "ztaylor.me/log"

func CutSourcePath(logger log.Service) log.Service {
	logger.Formatter().CutSourcePath(0)
	// logger.Formatter().CutSourcePath(1) // parent of caller, doesn't cut this package
	return logger
}
