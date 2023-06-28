package servers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func writeErrorResponse(log *zap.Logger, r *http.Request, w http.ResponseWriter, err error) {
	switch err.(type) {
	case BadRequestError:
		log.Error("handler input error", zap.String("url", r.RequestURI), zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
	case PanicError:
		log.Error("handler panic", zap.String("url", r.RequestURI), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	case InternalError:
		log.Error("handler error", zap.String("url", r.RequestURI), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	default:
		log.Error("unknown error", zap.String("url", r.RequestURI), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func panicRecovery(log *zap.Logger, r *http.Request, w http.ResponseWriter) {
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
	log *zap.Logger,
	handler func() (TOut, error),
) {
	r.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		defer panicRecovery(log, r, w)

		log.Info("handle request", zap.String("url", r.RequestURI))

		response, err := handler()
		if err != nil {
			writeErrorResponse(log, r, w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(response); err != nil {
			writeErrorResponse(log, r, w, ToInternalError(err))
		}

	}).Methods("GET")
}

func RouterHandlePost[TIn any, TOut any](
	r *mux.Router,
	url string,
	log *zap.Logger,
	handler func(request TIn) (TOut, error),
) {
	r.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		defer panicRecovery(log, r, w)

		log.Info("handle request", zap.String("url", r.RequestURI))

		var request TIn
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			writeErrorResponse(log, r, w, ToBadRequestError(err))
			return
		}

		response, err := handler(request)
		if err != nil {
			writeErrorResponse(log, r, w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(response); err != nil {
			writeErrorResponse(log, r, w, ToInternalError(err))
		}

	}).Methods("POST")
}

