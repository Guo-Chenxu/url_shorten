package logger

import (
	"url_shorten/conf"
	"io"
	"os"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func Init(logger conf.Logger) {
	hlog.SetLevel(getLevel(logger.Level))

	f, err := os.OpenFile(logger.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fileWriter := io.MultiWriter(f, os.Stdout)
	hlog.SetOutput(fileWriter)
}

func getLevel(level string) hlog.Level {
	switch level {
	case "trace":
		return hlog.LevelTrace
	case "debug":
		return hlog.LevelDebug
	case "info":
		return hlog.LevelInfo
	case "notice":
		return hlog.LevelNotice
	case "warn":
		return hlog.LevelWarn
	case "error":
		return hlog.LevelError
	case "fatal":
		return hlog.LevelFatal
	default:
		return hlog.LevelInfo
	}
}
