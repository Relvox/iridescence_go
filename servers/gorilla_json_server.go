package servers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/relvox/iridescence_go/logging"
)

func writeErrorResponse(log *slog.Logger, r *http.Request, w http.ResponseWriter, err error) {
	switch err.(type) {
	case BadRequestError:
		log.Error("handler input error", slog.String("url", r.RequestURI), logging.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
	case PanicError:
		log.Error("handler panic", slog.String("url", r.RequestURI), logging.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	case InternalError:
		log.Error("handler error", slog.String("url", r.RequestURI), logging.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	default:
		log.Error("unknown error", slog.String("url", r.RequestURI), logging.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func panicRecovery(log *slog.Logger, r *http.Request, w http.ResponseWriter) {
	if er := recover(); er != nil {
		err, ok := er.(error)
		if !ok {
			panic(er)
		}
		writeErrorResponse(log, r, w, ToPanicError(err))
	}
}

func RouterHandleGet[TOut any](
	r *mux.Router,
	url string,
	log *slog.Logger,
	handler func() (TOut, error),
) {
	r.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		defer panicRecovery(log, r, w)

		log.Info("handle request", slog.String("url", r.RequestURI))

		response, err := handler()
		if err != nil {
			writeErrorResponse(log, r, w, ToInternalError(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(response); err != nil {
			writeErrorResponse(log, r, w, ToInternalError(err))
		}

		log.Debug("request response", slog.String("url", r.RequestURI), slog.Any("response", response))
	}).Methods("GET")
}

func RouterHandleGetHTML(
	r *mux.Router,
	url string,
	log *slog.Logger,
	handler func() (string, error),
) {
	r.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		defer panicRecovery(log, r, w)

		log.Info("handle request", slog.String("url", r.RequestURI))

		response, err := handler()
		if err != nil {
			writeErrorResponse(log, r, w, ToInternalError(err))
			return
		}

		w.Header().Set("Content-Type", "text/html")
		n, err := w.Write([]byte(response))
		if err != nil {
			writeErrorResponse(log, r, w, ToInternalError(err))
		}

		log.Debug("request response", slog.String("url", r.RequestURI), slog.Int("response", n))
	}).Methods("GET")
}

func RouterHandlePost[TIn any, TOut any](
	r *mux.Router,
	url string,
	log *slog.Logger,
	handler func(request TIn) (TOut, error),
) {
	r.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		defer panicRecovery(log, r, w)

		log.Info("handle request", slog.String("url", r.RequestURI))

		var request TIn
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			writeErrorResponse(log, r, w, ToBadRequestError(err))
			return
		}

		log.Debug("request body", slog.String("url", r.RequestURI), slog.Any("body", request))

		response, err := handler(request)
		if err != nil {
			writeErrorResponse(log, r, w, ToInternalError(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(response); err != nil {
			writeErrorResponse(log, r, w, ToInternalError(err))
		}

		log.Debug("request response", slog.String("url", r.RequestURI), slog.Any("response", response))
	}).Methods("POST")
}

func RouterHandlePostHTML[TIn any](
	r *mux.Router,
	url string,
	log *slog.Logger,
	handler func(request TIn) (string, error),
) {
	r.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		defer panicRecovery(log, r, w)

		log.Info("handle request", slog.String("url", r.RequestURI))

		var request TIn
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			writeErrorResponse(log, r, w, ToBadRequestError(err))
			return
		}

		log.Debug("request body", slog.String("url", r.RequestURI), slog.Any("body", request))

		response, err := handler(request)
		if err != nil {
			writeErrorResponse(log, r, w, ToInternalError(err))
			return
		}

		w.Header().Set("Content-Type", "text/html")
		n, err := w.Write([]byte(response))
		if err != nil {
			writeErrorResponse(log, r, w, ToInternalError(err))
		}

		log.Debug("request response", slog.String("url", r.RequestURI), slog.Int("response", n))
	}).Methods("POST")
}
