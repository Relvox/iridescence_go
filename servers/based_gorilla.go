package servers

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func unifiedRouteHandler[TIn any, TOut any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	handler handleFunc[TIn, TOut],
	responseWriter func(w http.ResponseWriter, r *http.Request, response TOut) error,
) {
	requireBody := handler.handlerR != nil || handler.handlerRV != nil
	if err := method.validateBody(requireBody); err != nil {
		panic(err)
	}
	if err := handler.validate(); err != nil {
		panic(err)
	}

	r.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		defer panicRecovery(log, r, w)

		var err error
		var response TOut

		switch {
		case handler.handler != nil:
			response, err = handler.handler()

		case handler.handlerV != nil:
			vars := mux.Vars(r)
			log.Info("handle request", slog.String("url", r.RequestURI), slog.Any("vars", vars))
			response, err = handler.handlerV(vars)

		case handler.handlerR != nil:
			var request TIn
			request, err = handler.decodeRequest(r)
			if err != nil {
				writeErrorResponse(log, r, w, ToBadRequestError(err))
				return
			}
			log.Debug("request body", slog.String("url", r.RequestURI), slog.Any("body", request))
			response, err = handler.handlerR(request)

		case handler.handlerRV != nil:
			vars := mux.Vars(r)
			log.Info("handle request", slog.String("url", r.RequestURI), slog.Any("vars", vars))

			var request TIn
			request, err = handler.decodeRequest(r)
			if err != nil {
				writeErrorResponse(log, r, w, ToBadRequestError(err))
				return
			}
			log.Debug("request body", slog.String("url", r.RequestURI), slog.Any("body", request))
			response, err = handler.handlerRV(request, vars)
		}

		if err != nil {
			writeErrorResponse(log, r, w, ToInternalError(err))
			return
		}

		if err = responseWriter(w, r, response); err != nil {
			writeErrorResponse(log, r, w, ToInternalError(err))
		}
	}).Methods(string(method))
}
