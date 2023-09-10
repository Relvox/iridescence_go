package main

import (
	"flag"
	"log/slog"

	"github.com/gorilla/mux"
	"github.com/relvox/iridescence_go/logging"
	"github.com/relvox/iridescence_go/middleware"
	"github.com/relvox/iridescence_go/servers"
)

func main() {
	addrFlag := flag.String("addr", ":8080", "Address including port to listen on")
	dirFlag := flag.String("dir", ".", "Directory to serve logs from")
	logFlag := flag.String("log", "sawmill.log", "name of own log")
	flag.Parse()
	log := logging.NewLogger(
		logging.LoggingOptions{},
		logging.LoggingOptions{
			Target:                 *logFlag,
			PrefixFilenameWithTime: true,
			AddSource:              true,
			Level:                  slog.LevelDebug,
			JsonHandler:            true,
		})

	sawmillServ := NewSawmillServer(*dirFlag, getTemplatesFS, log)

	mwLogging := middleware.LoggingMiddleware(log, middleware.AllOptions-middleware.UserAgent-middleware.Response-middleware.RequestID)

	servers.ConfigureAndListen(*addrFlag,
		servers.DefaultHeaders, []string{"*"}, servers.DefaultMethods,
		log, []mux.MiddlewareFunc{mwLogging},
		sawmillServ,
		staticFileServer,
	)
}
