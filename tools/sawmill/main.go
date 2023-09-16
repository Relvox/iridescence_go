package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"

	mux_handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/relvox/iridescence_go/logging"
	"github.com/relvox/iridescence_go/middleware"
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
		},
	)
	mwLogging := middleware.LoggingMiddleware(log, middleware.AllOptions-middleware.UserAgent-middleware.Response-middleware.RequestID)

	sawmillServ := NewSawmillServer(*dirFlag, func() fs.FS { return templateFS }, log)
	log.Debug("registering server routes", slog.String("server", fmt.Sprintf("%T", sawmillServ)))

	router := mux.NewRouter()
	sawmillServ.RegisterRoutes(router)
	router.PathPrefix("/tmpl").Handler(http.StripPrefix("/tmpl",
		templateHandler.Parse().WithModelGetter(
			func(template string) any { return sawmillServ },
		),
	))
	router.PathPrefix("/").Handler(staticFileHandler)

	log.Info("Started Listening", slog.String("address", *addrFlag))
	router.Use(mwLogging)
	headersOk := mux_handlers.AllowedHeaders([]string{"X-Requested-With", "content-type"})
	originsOk := mux_handlers.AllowedOrigins([]string{"*"})
	methodsOk := mux_handlers.AllowedMethods([]string{"GET", "POST"})
	err := http.ListenAndServe(*addrFlag, mux_handlers.CORS(originsOk, headersOk, methodsOk)(router))
	if err != http.ErrServerClosed {
		log.Error("server crashed", logging.Error(err))
	}
	log.Info("Server Shutdown")
}
