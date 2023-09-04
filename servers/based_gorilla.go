package servers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type handleFunc[TIn, TOut any] struct {
	handler   func() (TOut, error)
	handlerV  func(vars map[string]string) (TOut, error)
	handlerR  func(request TIn) (TOut, error)
	handlerRV func(request TIn, vars map[string]string) (TOut, error)
}

func (h *handleFunc[TIn, TOut]) validate() error {
	handlers := 0
	if h.handler != nil {
		handlers++
	}
	if h.handlerV != nil {
		handlers++
	}
	if h.handlerR != nil {
		handlers++
	}
	if h.handlerRV != nil {
		handlers++
	}
	if handlers == 1 {
		return nil
	}
	return fmt.Errorf("exactly one handler must be set")
}

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
			err = json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				writeErrorResponse(log, r, w, ToBadRequestError(err))
				return
			}
			defer r.Body.Close()
			log.Debug("request body", slog.String("url", r.RequestURI), slog.Any("body", request))
			response, err = handler.handlerR(request)

		case handler.handlerRV != nil:
			vars := mux.Vars(r)
			log.Info("handle request", slog.String("url", r.RequestURI), slog.Any("vars", vars))

			var request TIn
			err = json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				writeErrorResponse(log, r, w, ToBadRequestError(err))
				return
			}
			defer r.Body.Close()
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

		log.Debug("request response", slog.String("url", r.RequestURI))
	}).Methods(string(method))
}
