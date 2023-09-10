package logging

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/relvox/iridescence_go/utils"
)

// LoggingOptions describes a handler
type LoggingOptions struct {
	// Filepath | "" or "stderr" for stderr | "stdout" for stdout (default "")
	Target string
	// adds a "2006_01_02_15_04_05" timestamp to the log file name
	// Not compatible with stderr/stdout (default false)
	PrefixFilenameWithTime bool
	// Whether to add the source location of the call to log (default false)
	AddSource bool
	// Minimum level to log for this handler (default INFO)
	Level slog.Level
	// false means TextHandler, true means JsonHandler
	// overridden by CustomHandler (default false)
	JsonHandler bool
	// ReplaceAttrs are called to rewrite each attribute before it is logged. (default nil)
	ReplaceAttrs []func(attrGroups []string, a slog.Attr) slog.Attr
	// Your custom handler, overrides all other fields (default nil)
	CustomHandler slog.Handler
}

// Handler instantiates a new handle based on the requested options
func (o LoggingOptions) Handler() slog.Handler {
	if o.CustomHandler != nil {
		return o.CustomHandler
	}
	var file *os.File
	switch strings.ToLower(o.Target) {
	case "", "stderr":
		file = os.Stderr
	case "stdout":
		file = os.Stdout
	default:
		dir, name := filepath.Dir(o.Target), filepath.Base(o.Target)
		if o.PrefixFilenameWithTime {
			o.Target = fmt.Sprintf("%s/%s_%s", dir, utils.TimestampNow(), name)
		}
		os.MkdirAll(dir, 0644)
		var err error
		file, err = os.OpenFile(o.Target, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
	}
	handlerOptions := &slog.HandlerOptions{
		AddSource: o.AddSource,
		Level:     o.Level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			for _, rep := range o.ReplaceAttrs {
				a = rep(groups, a)
			}
			return a
		},
	}
	if o.JsonHandler {
		return slog.NewJSONHandler(file, handlerOptions)
	} else {
		return slog.NewTextHandler(file, handlerOptions)
	}
}
