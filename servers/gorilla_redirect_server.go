package servers

import (
	"log/slog"

	"github.com/gorilla/mux"
	"github.com/relvox/iridescence_go/sets"
)

func RouterHandleRedirect(
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	handler func() (string, error),
) {
	hFunc := handleFunc[sets.Unit, string]{handler: handler}
	unifiedRouteHandler[sets.Unit, string](r, log, method, url, hFunc, redirectResponseWriter)
	log.Debug("handle Redirect", slog.String("method", string(method)), slog.String("url", url))
}

func RouterHandleRedirectVars(
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	handler func(vars map[string]string) (string, error),
) {
	hFunc := handleFunc[sets.Unit, string]{handlerV: handler}
	unifiedRouteHandler[sets.Unit, string](r, log, method, url, hFunc, redirectResponseWriter)
	log.Debug("handle Redirect", slog.String("method", string(method)), slog.String("url", url))
}

func RouterHandleRedirectRequest[TIn any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	handler func(request TIn) (string, error),
) {
	hFunc := handleFunc[TIn, string]{handlerR: handler, decodeRequest: formRequestDecoder[TIn]}
	unifiedRouteHandler[TIn, string](r, log, method, url, hFunc, redirectResponseWriter)
	log.Debug("handle Redirect", slog.String("method", string(method)), slog.String("url", url))
}

func RouterHandleRedirectRequestVars[TIn any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	handler func(request TIn, vars map[string]string) (string, error),
) {
	hFunc := handleFunc[TIn, string]{handlerRV: handler, decodeRequest: formRequestDecoder[TIn]}
	unifiedRouteHandler[TIn, string](r, log, method, url, hFunc, redirectResponseWriter)
	log.Debug("handle Redirect", slog.String("method", string(method)), slog.String("url", url))
}
