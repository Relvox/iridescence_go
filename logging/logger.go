package logging

import (
	"fmt"
	"os"

	"log/slog"
)

func NewLogger(logPath string, fileLevel, stdoutLevel slog.Level) *slog.Logger {
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("create or append to log file: %w", err))
	}
	fileHandler := slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		AddSource: true,
		Level:     fileLevel,
	})
	stdoutHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: stdoutLevel,
	})

	return slog.New(NewTeeHandler(stdoutHandler, fileHandler))
}
