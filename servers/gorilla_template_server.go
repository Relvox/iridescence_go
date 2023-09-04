package servers

import (
	"html/template"
	"log/slog"

	"github.com/gorilla/mux"

	"github.com/relvox/iridescence_go/utils"
)

func RouterHandleTemplate[TData any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	template *template.Template,
	handler func() (TData, error),
) {
	hFunc := handleFunc[utils.Unit, TData]{handler: handler}
	unifiedRouteHandler[utils.Unit, TData](r, log, method, url, hFunc, templateResponseWriterFor[TData](template))
	log.Debug("handle HTML template", slog.String("method", string(method)), slog.String("url", url), slog.String("template", template.Name()))
}

func RouterHandleTemplateVars[TData any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	template *template.Template,
	handler func(vars map[string]string) (TData, error),
) {
	hFunc := handleFunc[utils.Unit, TData]{handlerV: handler}
	unifiedRouteHandler[utils.Unit, TData](r, log, method, url, hFunc, templateResponseWriterFor[TData](template))
	log.Debug("handle HTML template", slog.String("method", string(method)), slog.String("url", url), slog.String("template", template.Name()))
}

func RouterHandleTemplateRequest[TData, TIn any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	template *template.Template,
	handler func(request TIn) (TData, error),
) {
	hFunc := handleFunc[TIn, TData]{handlerR: handler}
	unifiedRouteHandler[TIn, TData](r, log, method, url, hFunc, templateResponseWriterFor[TData](template))
	log.Debug("handle HTML template", slog.String("method", string(method)), slog.String("url", url), slog.String("template", template.Name()))
}

func RouterHandleTemplateRequestVars[TData, TIn any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	template *template.Template,
	handler func(request TIn, vars map[string]string) (TData, error),
) {
	hFunc := handleFunc[TIn, TData]{handlerRV: handler}
	unifiedRouteHandler[TIn, TData](r, log, method, url, hFunc, templateResponseWriterFor[TData](template))
	log.Debug("handle HTML template", slog.String("method", string(method)), slog.String("url", url), slog.String("template", template.Name()))
}
