package logging

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"log/slog"

	"github.com/relvox/iridescence_go/utils"
)

func NewLogger(logPath string, fileLevel, stdoutLevel slog.Level) *slog.Logger {
	dir, name  := filepath.Dir(logPath), filepath.Base(logPath)
	file := fmt.Sprintf("%s/%s_%s", dir, utils.TimestampNow(), name)
	os.MkdirAll(dir, 0644)
	logFile, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
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

type NullHandler utils.Unit

func (h NullHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}
func (h NullHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}
func (h NullHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}
func (h NullHandler) WithGroup(_ string) slog.Handler {
	return h
}

func NewNullLogger() *slog.Logger {
	return slog.New(NullHandler{})
}
