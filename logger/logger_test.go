package logger

import (
	"log/slog"
	"testing"
)

func TestLog(t *testing.T) {
	slog.Debug("debug message")
	slog.Info("info message")
	slog.Warn("warning message")
	slog.Error("error message")
}
