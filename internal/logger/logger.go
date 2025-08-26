package logger

import (
	"log/slog"
	"os"
)

func init() {
	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: false,
	}

	baseHandler := slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(baseHandler)

	slog.SetDefault(logger)
}
