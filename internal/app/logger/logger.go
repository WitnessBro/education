package logger

import (
	"log/slog"
	"os"
)

func NewLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	slog.Info("Info message")
	slog.Info("server started", "port", "8000")
	slog.Debug("debug message", "key", "hui")
	slog.Error("Fatal Error")
}
