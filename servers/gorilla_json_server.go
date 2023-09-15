package servers

import (
	"log/slog"

	"github.com/gorilla/mux"
	"github.com/relvox/iridescence_go/sets"
)

// RouterHandleJSON JSON -> JSON
func RouterHandleJSON[TOut any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	handler func() (TOut, error),
) {
	hFunc := handleFunc[sets.Unit, TOut]{handler: handler}
	unifiedRouteHandler[sets.Unit, TOut](r, log, method, url, hFunc, jsonResponseWriter)
	log.Debug("handle JSON", slog.String("method", string(method)), slog.String("url", url))
}

// RouterHandleJSONVars JSON -> JSON
func RouterHandleJSONVars[TOut any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	handler func(vars map[string]string) (TOut, error),
) {
	hFunc := handleFunc[sets.Unit, TOut]{handlerV: handler}
	unifiedRouteHandler[sets.Unit, TOut](r, log, method, url, hFunc, jsonResponseWriter)
	log.Debug("handle JSON", slog.String("method", string(method)), slog.String("url", url))
}

// RouterHandleJSONRequest JSON -> JSON
func RouterHandleJSONRequest[TIn any, TOut any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	handler func(request TIn) (TOut, error),
) {
	hFunc := handleFunc[TIn, TOut]{handlerR: handler, decodeRequest: jsonRequestDecoder[TIn]}
	unifiedRouteHandler[TIn, TOut](r, log, method, url, hFunc, jsonResponseWriter)
	log.Debug("handle JSON", slog.String("method", string(method)), slog.String("url", url))
}

// RouterHandleJSONRequestVars JSON -> JSON
func RouterHandleJSONRequestVars[TIn any, TOut any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	handler func(request TIn, vars map[string]string) (TOut, error),
) {
	hFunc := handleFunc[TIn, TOut]{handlerRV: handler, decodeRequest: jsonRequestDecoder[TIn]}
	unifiedRouteHandler[TIn, TOut](r, log, method, url, hFunc, jsonResponseWriter)
	log.Debug("handle JSON", slog.String("method", string(method)), slog.String("url", url))
}
