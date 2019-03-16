package log

import (
	"io"
	"os"
	"time"
)

// NewRollingWriter creates a log file io.Writer for a system path, which
// updates to a new file name every day
func NewRollingWriter(path string) io.Writer {
	w := &rollingWriter{
		PathPrefix: path,
		Publish:    make(chan []byte, 0),
		Close:      make(chan bool, 0),
	}
	go w.start()
	return w
}

type rollingWriter struct {
	PathPrefix string
	Publish    chan []byte
	Close      chan bool
	w          io.Writer
}

func (w *rollingWriter) Write(bytes []byte) (int, error) {
	go w.write(bytes)
	return 0, nil
}

func (w *rollingWriter) write(bytes []byte) {
	w.Publish <- bytes
}

func (w *rollingWriter) fileFormat(time time.Time) string {
	return w.PathPrefix + time.Format("2006_01_02")
}

func (w *rollingWriter) start() {
	roller := time.NewTicker(24 * time.Hour)
	for {
		fileTitle := w.fileFormat(time.Now())
		fileName := fileTitle + ".log"
		file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		w.w = file
		w.wait(roller)
		file.Close()
	}
}

func (w *rollingWriter) wait(roller *time.Ticker) {
	for {
		select {
		case <-roller.C:
			return
		case msg := <-w.Publish:
			w.Write(msg)
		}
	}
}
