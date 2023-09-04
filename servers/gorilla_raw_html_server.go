package servers

import (
	"log/slog"

	"github.com/gorilla/mux"

	"github.com/relvox/iridescence_go/utils"
)

func RouterHandleHTML[TOut any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	handler func() (TOut, error),
) {
	hFunc := handleFunc[utils.Unit, TOut]{handler: handler}
	unifiedRouteHandler[utils.Unit, TOut](r, log, method, url, hFunc, htmlResponseWriter)
	log.Debug("handle HTML", slog.String("method", string(method)), slog.String("url", url))
}

func RouterHandleHTMLVars[TOut any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	handler func(vars map[string]string) (TOut, error),
) {
	hFunc := handleFunc[utils.Unit, TOut]{handlerV: handler}
	unifiedRouteHandler[utils.Unit, TOut](r, log, method, url, hFunc, htmlResponseWriter)
	log.Debug("handle HTML", slog.String("method", string(method)), slog.String("url", url))
}

func RouterHandleHTMLRequest[TIn any, TOut any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	handler func(request TIn) (TOut, error),
) {
	hFunc := handleFunc[TIn, TOut]{handlerR: handler}
	unifiedRouteHandler[TIn, TOut](r, log, method, url, hFunc, htmlResponseWriter)
	log.Debug("handle HTML", slog.String("method", string(method)), slog.String("url", url))
}

func RouterHandleHTMLRequestVars[TIn any, TOut any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	handler func(request TIn, vars map[string]string) (TOut, error),
) {
	hFunc := handleFunc[TIn, TOut]{handlerRV: handler}
	unifiedRouteHandler[TIn, TOut](r, log, method, url, hFunc, htmlResponseWriter)
	log.Debug("handle HTML", slog.String("method", string(method)), slog.String("url", url))
}
