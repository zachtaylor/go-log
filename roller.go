package log

import (
	"os"

	"ztaylor.me/cast"
)

// NewRoller creates a log file io.WriteCloser for a system path, which
// updates to a new file name every day
func NewRoller(path string) cast.WriteCloser {
	w := &roller{
		PathPrefix: path,
		publish:    make(chan []byte, 0),
		done:       make(chan bool, 0),
	}
	go w.start()
	return w
}

type roller struct {
	PathPrefix string
	publish    chan []byte
	done       chan bool
	w          cast.WriteCloser
}

func (w *roller) Write(bytes []byte) (int, error) {
	go w.write(bytes)
	return 0, nil
}

func (w *roller) write(bytes []byte) {
	w.publish <- bytes
}

func (w *roller) Close() error {
	go w.close()
	return nil
}

func (w *roller) close() {
	close(w.done)
}

func (w *roller) fileFormat(time cast.Time) string {
	return w.PathPrefix + time.Format("2006_01_02")
}

func (w *roller) start() {
	roller := cast.NewTicker(24 * cast.Hour)
	for {
		fileTitle := w.fileFormat(cast.Now())
		fileName := fileTitle + ".log"
		file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		w.w = file
		err := w.wait(roller)
		file.Close()
		if err == cast.EOF {
			return
		}
	}
}

// wait ends with the timer (nil) or when w.done is closed (io.EOF)
func (w *roller) wait(roller *cast.Ticker) error {
	for {
		select {
		case <-roller.C:
			return nil
		case msg := <-w.publish:
			if _, err := w.w.Write(msg); err != nil {
				return err
			}
		case <-w.done:
			return cast.EOF
		}
	}
}
