package x

import "ztaylor.me/log"

func CutSourcePath(logger log.Service) log.Service {
	logger.Format().CutPathSource()
	return logger
}
