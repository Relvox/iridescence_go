package servers

import (
	"html/template"
	"log/slog"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/relvox/iridescence_go/files"
	"github.com/relvox/iridescence_go/logging"
	"github.com/relvox/iridescence_go/servers"
)

type SawmillServer struct {
	Paths    []string
	TmplPath string

	SelectedFiles []string
	Query         string
	Trackers      map[string]*files.FileTracker

	Log *slog.Logger
}

func NewSawmillServer(logPath, tmplPath string, log *slog.Logger) *SawmillServer {
	logFiles, err := files.GetFilenames(logPath, "log")
	if err != nil {
		log.Error("failed to get template filenames", logging.Error(err))
	}
	return &SawmillServer{
		Paths:         logFiles,
		TmplPath:      tmplPath,
		SelectedFiles: make([]string, 0),
		Query:         "",
		Trackers:      map[string]*files.FileTracker{},
		Log:           log,
	}
}

func (s *SawmillServer) RegisterRoutes(r *mux.Router) {
	tmplFiles, err := files.GetFilenames(s.TmplPath, "html")
	if err != nil {
		s.Log.Error("failed to get template filenames", logging.Error(err))
	}
	rootTmpl := template.New("root").Funcs(template.FuncMap{
		"contains": func(s reflect.Value, k reflect.Value) bool {
			for i := 0; i < s.Len(); i++ {
				if s.Index(i).Interface() == k.Interface() {
					return true
				}
			}
			return false
		},
	})

	rootTmpl, err = rootTmpl.ParseFiles(tmplFiles...)
	if err != nil {
		panic(err)
	}
	templates := rootTmpl.Templates()
	for _, t := range templates {
		s.Log.Info("loaded template", slog.String("name", t.Name()))

		servers.RouterHandleTemplate[*SawmillServer](
			r, s.Log, servers.GET,
			"/"+files.IsolateName(t.Name()), t,
			func() (*SawmillServer, error) {
				return s, nil
			},
		)
	}

	type files_form struct {
		Files []string `schema:"files"`
	}
	servers.RouterHandleTemplatesRequest[map[string]any, files_form](
		r, s.Log,
		servers.POST, "/select", []*template.Template{
			rootTmpl.Lookup("log_file_selector.html"),
			rootTmpl.Lookup("result_table.html"),
		}, func(form files_form) (map[string]any, error) {
			s.SelectedFiles = form.Files
			return map[string]any{
				"log_file_selector.html": s,
				"result_table.html":      s,
			}, nil
		},
	)

	type query_form struct {
		Query string `schema:"query"`
	}

	servers.RouterHandleTemplateRequest[*SawmillServer](
		r, s.Log,
		servers.POST, "/query", rootTmpl.Lookup("result_table.html"),
		func(form query_form) (*SawmillServer, error) {
			s.Query = form.Query
			return s, nil
		},
	)
}
