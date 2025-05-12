package logger

import (
	"log"
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger(LogPath string, LogLevel slog.Level) {
	file, err := os.OpenFile(LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件 err:%v", err)
	}
	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: LogLevel,
	})
	Logger = slog.New(handler)
	slog.SetDefault(Logger)

}
