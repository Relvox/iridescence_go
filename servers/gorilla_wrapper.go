package servers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/relvox/iridescence_go/logging"
	"github.com/relvox/iridescence_go/middleware"
)

var (
	DefaultHeaders = []string{"X-Requested-With", "content-type"}
	DefaultMethods = []string{"GET", "POST"}
)

type SupportedServer interface {
	RegisterRoutes(r *mux.Router)
}

func ConfigureAndListen(address string, headers, origins, methods []string, log *slog.Logger, logOpts middleware.LoggingOptions, servers ...SupportedServer) {
	router := mux.NewRouter()
	headersOk := handlers.AllowedHeaders(headers)
	originsOk := handlers.AllowedOrigins(origins)
	methodsOk := handlers.AllowedMethods(methods)

	for _, s := range servers {
		log.Debug("registering server routes", slog.String("server", fmt.Sprintf("%T", s)))
		s.RegisterRoutes(router)
	}

	log.Info("Started Listening", slog.String("address", address))
	err := http.ListenAndServe(address, handlers.CORS(originsOk, headersOk, methodsOk)(
		middleware.LoggingMiddleware(log, logOpts)(router)))
	if err != http.ErrServerClosed {
		log.Error("server crashed", logging.Error(err))
	}
	log.Info("Server Shutdown")
}
