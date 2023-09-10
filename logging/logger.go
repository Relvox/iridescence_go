package logging

import (
	"log/slog"
)

func NewLogger(opts ...LoggingOptions) *slog.Logger {
	var handlers []slog.Handler
	for _, opt := range opts {
		handlers = append(handlers, opt.Handler())
	}
	return slog.New(NewTeeHandler(handlers...))
}
