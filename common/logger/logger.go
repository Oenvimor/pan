package logger

import (
	"context"
	"log"
	"log/slog"
	"os"
)

var Logger *slog.Logger

type MultiHandler struct {
	handlers []slog.Handler
}

func (m *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range m.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (m *MultiHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, h := range m.handlers {
		_ = h.Handle(ctx, r) // 忽略错误或记录
	}
	return nil
}

func (m *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandlers := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		newHandlers[i] = h.WithAttrs(attrs)
	}
	return &MultiHandler{handlers: newHandlers}
}

func (m *MultiHandler) WithGroup(name string) slog.Handler {
	newHandlers := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		newHandlers[i] = h.WithGroup(name)
	}
	return &MultiHandler{handlers: newHandlers}
}

func InitLogger(LogPath string, LogLevel slog.Level) {
	file, err := os.OpenFile(LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件 err:%v", err)
	}

	fileHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: LogLevel,
	})

	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // 控制台仅输出 info+
	})

	// 自定义组合 handler
	multiHandler := &MultiHandler{
		handlers: []slog.Handler{fileHandler, consoleHandler},
	}

	Logger = slog.New(multiHandler)
	slog.SetDefault(Logger)
}
