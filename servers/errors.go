package servers

import (
	"log/slog"
	"net/http"

	"github.com/relvox/iridescence_go/logging"
)

type BadRequestError string

func ToBadRequestError(err error) BadRequestError { return BadRequestError(err.Error()) }

func (e BadRequestError) Error() string { return string(e) }

type InternalError string

func ToInternalError(err error) InternalError { return InternalError(err.Error()) }

func (e InternalError) Error() string { return string(e) }

type PanicError string

func ToPanicError(err error) PanicError { return PanicError(err.Error()) }

func (e PanicError) Error() string { return string(e) }

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
