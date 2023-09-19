package sawmill

import (
	"encoding/json"
	"html"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/gorilla/mux"

	"github.com/relvox/iridescence_go/files"
	"github.com/relvox/iridescence_go/logging"
	"github.com/relvox/iridescence_go/servers"
	"github.com/relvox/iridescence_go/templates/funcs"
	"github.com/relvox/iridescence_go/utils"
)

type SawmillServer struct {
	TmplFS   fs.FS
	StaticFS fs.FS
	LogsFS   fs.FS

	Trackers map[string]*FSLineTracker

	SelectedFiles []string
	Query         string
	CurrentLines  []string

	Log *slog.Logger
}

func NewSawmillServer(logPath string, tmplFS, staticFS fs.FS, log *slog.Logger) *SawmillServer {
	return &SawmillServer{
		TmplFS:   tmplFS,
		StaticFS: staticFS,
		LogsFS:   os.DirFS(filepath.Dir(logPath)),

		Trackers: make(map[string]*FSLineTracker),

		SelectedFiles: make([]string, 0),
		Query:         "",
		CurrentLines:  []string{},

		Log: log,
	}
}

func (s *SawmillServer) UpdateSelection(selected []string) error {
	s.CurrentLines = make([]string, 0)
	s.SelectedFiles = selected
	for _, file := range selected {
		tracker, ok := s.Trackers[file]
		if !ok {
			s.Trackers[file] = NewFSLineTracker(s.LogsFS, file)
			tracker = s.Trackers[file]
		} else {
			if err := tracker.Refresh(); err != nil {
				return err
			}
		}
		s.CurrentLines = append(s.CurrentLines, tracker.Content...)
	}
	return nil
}

var onlyLowercase = regexp.MustCompile(`^[\"a-z0-9\s._\-\:]+$`)
var containsRegex = regexp.MustCompile(`[$*+?{}\[\]\\|()\^]`)
var splitRegex = regexp.MustCompile(`"(.*?)"|\S+`)

func (s *SawmillServer) ExecuteQuery(query string) error {
	query = html.UnescapeString(query)
	s.Query = query
	s.CurrentLines = make([]string, 0)

	var searchFunc func(line, query string) bool
	if containsRegex.MatchString(query) {
		regex, err := regexp.Compile(query)
		if err != nil {
			return err
		}
		searchFunc = func(line, query string) bool {
			return regex.MatchString(line)
		}
	} else {
		doLower := onlyLowercase.MatchString(query)
		matches := splitRegex.FindAllString(query, -1)
		for i, match := range matches {
			matches[i] = strings.Trim(match, "\"")
		}
		searchFunc = func(line, query string) bool {
			checkLine := line
			if doLower {
				checkLine = strings.ToLower(checkLine)
			}
			for _, match := range matches {
				if strings.Contains(checkLine, match) {
					return true
				}
			}
			return false
		}
	}

	trackerKeys := utils.SortedKeys(s.Trackers)
	slices.Reverse(trackerKeys)
	for _, file := range trackerKeys {
		if !slices.Contains(s.SelectedFiles, file) {
			continue
		}
		tracker := s.Trackers[file]
		tracker.Refresh()
		for _, line := range tracker.Content {
			if searchFunc(line, query) {
				s.CurrentLines = append(s.CurrentLines, line)
			}
		}
	}
	return nil
}

func (s *SawmillServer) RegisterRoutes(r *mux.Router) {
	type files_form struct {
		Files []string `schema:"files"`
	}
	servers.RouterHandleRedirectRequest[files_form](r, s.Log, servers.POST,
		"/sawmill/select", func(form files_form) (string, error) {
			err := s.UpdateSelection(form.Files)
			if err != nil {
				return "", err
			}
			if len(s.Query) > 0 {
				err = s.ExecuteQuery(s.Query)
			}
			return "/sawmill/tmpl/log_file_selector|result_table", err
		},
	)

	type query_form struct {
		Query string `schema:"query"`
	}
	servers.RouterHandleRedirectRequest[query_form](r, s.Log, servers.POST,
		"/sawmill/query", func(form query_form) (string, error) {
			return "/sawmill/tmpl/result_table", s.ExecuteQuery(form.Query)
		},
	)

	r.PathPrefix("/sawmill/tmpl").Handler(http.StripPrefix("/sawmill/tmpl",
		SawmillTemplateHandler.WithModelGetter(
			func(template string) any { return s },
		).WithFuncs(template.FuncMap{"contains": funcs.SliceContains}).Parse(),
	))

	r.PathPrefix("/sawmill").Handler(http.StripPrefix("/sawmill",
		http.FileServer(http.FS(s.StaticFS))),
	)
}

func (s *SawmillServer) Paths() []string {
	logFiles, err := files.ListFS(s.LogsFS, ".", "*.log")
	if err != nil {
		s.Log.Error("failed to get template filenames", logging.Error(err))
	}
	slices.Reverse(logFiles)
	return logFiles
}

func (s *SawmillServer) Events() []LogLine {
	var result []LogLine
	for _, line := range s.CurrentLines {
		if len(line) == 0 {
			continue
		}

		var obj map[string]any
		err := json.Unmarshal([]byte(line), &obj)
		if err != nil {
			s.Log.Error("failed to unmarshal", logging.Error(err), slog.String("line", line))
			continue
		}

		time := strings.Split(obj["time"].(string), "+")[0]
		time = strings.Replace(time, "T", " ", 1)
		time = strings.Replace(time, ".", " .", 1)
		delete(obj, "time")
		level := obj["level"].(string)
		delete(obj, "level")
		message := obj["msg"].(string)
		delete(obj, "msg")
		source, _ := json.MarshalIndent(obj["source"], "", "  ")
		delete(obj, "source")
		result = append(result, LogLine{
			Time:    time,
			Level:   level,
			Message: message,
			Source:  string(source),
			Object:  obj,
		})
	}
	return result
}
