package servers

import (
	"encoding/json"
	"html/template"
	"log/slog"
	"reflect"
	"regexp"
	"slices"
	"strings"

	"github.com/gorilla/mux"
	"github.com/relvox/iridescence_go/files"
	"github.com/relvox/iridescence_go/logging"
	"github.com/relvox/iridescence_go/servers"
	"github.com/relvox/iridescence_go/utils"
)

type SawmillServer struct {
	Paths    []string
	TmplPath string

	SelectedFiles []string
	Query         string
	Trackers      map[string]*files.FileTracker

	CurrentLines []string

	Log *slog.Logger
}

func NewSawmillServer(logPath, tmplPath string, log *slog.Logger) *SawmillServer {
	logFiles, err := files.GetFilenames(logPath, "log")
	if err != nil {
		log.Error("failed to get template filenames", logging.Error(err))
	}
	slices.Reverse(logFiles)
	return &SawmillServer{
		Paths:         logFiles,
		TmplPath:      tmplPath,
		SelectedFiles: make([]string, 0),
		Query:         "",
		Trackers:      map[string]*files.FileTracker{},
		Log:           log,
	}
}

func (s *SawmillServer) UpdateSelection(selected []string) error {
	s.CurrentLines = make([]string, 0)
	s.SelectedFiles = selected
	for _, file := range selected {
		tracker, ok := s.Trackers[file]
		if !ok {
			s.Trackers[file] = files.NewFileTracker(file)
			tracker = s.Trackers[file]
		} else {
			if err := tracker.Refresh(); err != nil {
				return err
			}
		}
		lines := strings.Split(string(tracker.GetContent()), "\n")
		slices.Reverse(lines)
		s.CurrentLines = append(s.CurrentLines, lines...)
	}
	return nil
}

var onlyLowercase = regexp.MustCompile(`^[\"a-z0-9\s._\-\:]+$`)
var containsRegex = regexp.MustCompile(`[$*+?{}\[\]\\|()]`)
var splitRegex = regexp.MustCompile(`"(.*?)"|\S+`)

func (s *SawmillServer) ExecuteQuery(query string) error {
	s.Query = query
	s.CurrentLines = make([]string, 0)
	if containsRegex.Match([]byte(query)) {
		regex, err := regexp.Compile(query)
		if err != nil {
			return err
		}
		trackerKeys := utils.SortedKeys(s.Trackers)
		slices.Reverse(trackerKeys)
		for _, file := range trackerKeys {
			tracker := s.Trackers[file]
			if !slices.Contains(s.SelectedFiles, file) {
				continue
			}
			tracker.Refresh()
			lines := strings.Split(string(tracker.GetContent()), "\n")
			slices.Reverse(lines)
			for _, line := range lines {
				if regex.Match([]byte(line)) {
					s.CurrentLines = append(s.CurrentLines, line)
				}
			}
		}
	} else {
		doLower := onlyLowercase.Match([]byte(query))
		matches := splitRegex.FindAllString(query, -1)
		for i, match := range matches {
			matches[i] = strings.Trim(match, "\"")
		}
		trackerKeys := utils.SortedKeys(s.Trackers)
		slices.Reverse(trackerKeys)
		for _, file := range trackerKeys {
			tracker := s.Trackers[file]
			if !slices.Contains(s.SelectedFiles, file) {
				continue
			}
			tracker.Refresh()
			lines := strings.Split(string(tracker.GetContent()), "\n")
			slices.Reverse(lines)
			for _, line := range lines {
				matched := -1
				checkLine := line
				for i, match := range matches {
					if doLower {
						checkLine = strings.ToLower(checkLine)
					}
					if !strings.Contains(checkLine, match) {
						continue
					}
					matched = i
					break
				}
				if matched >= 0 {
					s.CurrentLines = append(s.CurrentLines, line)
				}
			}
		}
	}
	return nil
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
			s.UpdateSelection(form.Files)
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
			s.ExecuteQuery(form.Query)
			return s, nil
		},
	)
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
