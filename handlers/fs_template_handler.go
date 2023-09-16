package handlers

import (
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"strings"

	"github.com/relvox/iridescence_go/logging"
	"github.com/relvox/iridescence_go/middleware"
)

// var dotSuffixRegex = regexp.MustCompile(`[^\.]\.\s*$`)

type TemplateModelGetter func(template string) any

type FSTemplateHandler struct {
	fs.FS
	FuncMap   template.FuncMap
	GetModel  TemplateModelGetter
	Root      *template.Template
	DebugMode bool
	Log       *slog.Logger
}

func NewFSTemplateHandler(fs fs.FS) FSTemplateHandler {
	return FSTemplateHandler{
		FS:   fs,
		Root: template.New("root"),
	}
}

func (s FSTemplateHandler) WithModelGetter(getter TemplateModelGetter) FSTemplateHandler {
	s.GetModel = getter
	return s
}

func (s FSTemplateHandler) WithModelMap(modelMap map[string]any) FSTemplateHandler {
	s.GetModel = func(template string) any {
		return modelMap[template]
	}
	return s
}

func (s FSTemplateHandler) WithFuncs(funcs template.FuncMap) FSTemplateHandler {
	s.FuncMap = funcs
	s.Root.Funcs(s.FuncMap)
	return s
}

func (s FSTemplateHandler) WithLog(log *slog.Logger) FSTemplateHandler {
	s.Log = log
	return s
}

func (s FSTemplateHandler) WithDebug() FSTemplateHandler {
	s.DebugMode = true
	return s
}

func (s FSTemplateHandler) Parse() FSTemplateHandler {
	s.Root = template.Must(s.Root.ParseFS(s.FS, "*.gohtml"))
	return s
}

func (s FSTemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	contextLog := r.Context().Value(middleware.Key("log"))
	if contextLog != nil {
		s.Log = contextLog.(*slog.Logger)
	}
	if s.DebugMode {
		s.Root = template.Must(template.New("root").Funcs(s.FuncMap).ParseFS(s.FS, "*.gohtml"))
	}

	requestPath := r.URL.Path
	tmplName := strings.TrimLeft(requestPath, "/") + ".gohtml"
	model := s.GetModel(requestPath)
	if model == nil {
		s.Log.Error("missing model for template", slog.String("template", tmplName))
	}

	tmpl := s.Root.Lookup(tmplName)
	if tmpl == nil {
		s.Log.Error("Failed to lookup template", slog.String("template", tmplName))
		return
	}

	err := tmpl.Execute(w, model)
	if err != nil {
		s.Log.Error("Failed to execute template", slog.String("template", tmplName), logging.Error(err))
		return
	}
}
