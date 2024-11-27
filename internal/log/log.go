package log

import (
	"log/slog"
	"os"
)

func InitLog() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	return slog.New(handler)
}
