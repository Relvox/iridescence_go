package logging

import (
	"context"
	"errors"

	"log/slog"
)

type TeeHandler struct {
	Handlers []slog.Handler
}

func NewTeeHandler(handlers ...slog.Handler) *TeeHandler {
	return &TeeHandler{
		Handlers: handlers,
	}
}

func (h *TeeHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	for _, handler := range h.Handlers {
		if handler.Enabled(ctx, lvl) {
			return true
		}
	}
	return false
}

func (h *TeeHandler) Handle(ctx context.Context, record slog.Record) error {
	var errs []error
	for _, handler := range h.Handlers {
		if !handler.Enabled(ctx, record.Level) {
			continue
		}
		err := handler.Handle(ctx, record)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.Join(errs...)
}

func (h *TeeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	var newHandlers []slog.Handler = make([]slog.Handler, 0, len(h.Handlers))
	for _, handler := range h.Handlers {
		newHandlers = append(newHandlers, handler.WithAttrs(attrs))
	}
	return NewTeeHandler(newHandlers...)
}

func (h *TeeHandler) WithGroup(name string) slog.Handler {
	var newHandlers []slog.Handler = make([]slog.Handler, 0, len(h.Handlers))
	for _, handler := range h.Handlers {
		newHandlers = append(newHandlers, handler.WithGroup(name))
	}
	return NewTeeHandler(newHandlers...)
}
