package log

import colorable "github.com/mattn/go-colorable"

func init() {
	logger.logger.SetOutput(colorable.NewColorableStdout())
	logger.logger.Formatter = &Formatter{
		Color: true,
	}
}
