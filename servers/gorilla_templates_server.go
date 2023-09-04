package servers

import (
	"html/template"
	"log/slog"

	"github.com/gorilla/mux"

	"github.com/relvox/iridescence_go/utils"
)

func RouterHandleTemplates[TData ~map[string]any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	templates []*template.Template,
	handler func() (TData, error),
) {
	hFunc := handleFunc[utils.Unit, TData]{handler: handler}
	unifiedRouteHandler[utils.Unit, TData](r, log, method, url, hFunc, templatesResponseWriterFor[TData](templates))
	log.Debug("handle HTML templates", slog.String("method", string(method)), slog.String("url", url), slog.String("templates", getTemplateNames(templates)))
}

func RouterHandleTemplatesVars[TData ~map[string]any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	templates []*template.Template,
	handler func(vars map[string]string) (TData, error),
) {
	hFunc := handleFunc[utils.Unit, TData]{handlerV: handler}
	unifiedRouteHandler[utils.Unit, TData](r, log, method, url, hFunc, templatesResponseWriterFor[TData](templates))
	log.Debug("handle HTML templates", slog.String("method", string(method)), slog.String("url", url), slog.String("templates", getTemplateNames(templates)))
}

func RouterHandleTemplatesRequest[TData ~map[string]any, TIn any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	templates []*template.Template,
	handler func(request TIn) (TData, error),
) {
	hFunc := handleFunc[TIn, TData]{handlerR: handler}
	unifiedRouteHandler[TIn, TData](r, log, method, url, hFunc, templatesResponseWriterFor[TData](templates))
	log.Debug("handle HTML templates", slog.String("method", string(method)), slog.String("url", url), slog.String("templates", getTemplateNames(templates)))
}

func RouterHandleTemplatesRequestVars[TData ~map[string]any, TIn any](
	r *mux.Router,
	log *slog.Logger,
	method HttpMethod,
	url string,
	templates []*template.Template,
	handler func(request TIn, vars map[string]string) (TData, error),
) {
	hFunc := handleFunc[TIn, TData]{handlerRV: handler}
	unifiedRouteHandler[TIn, TData](r, log, method, url, hFunc, templatesResponseWriterFor[TData](templates))
	log.Debug("handle HTML templates", slog.String("method", string(method)), slog.String("url", url), slog.String("templates", getTemplateNames(templates)))
}
