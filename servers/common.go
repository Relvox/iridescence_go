package servers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
	PATCH  HttpMethod = "PATCH"
)

func (m HttpMethod) validateBody(requireBody bool) error {
	switch m {
	case GET, DELETE:
		if requireBody {
			return fmt.Errorf("GET|DELETE cannot require body")
		}
		return nil
	case POST, PUT, PATCH:
		if !requireBody {
			return fmt.Errorf("POST|PUT|PATCH require body")
		}
		return nil
	}
	return fmt.Errorf("bad method: %s", m)
}

func templateResponseWriterFor[TData any](template *template.Template) func(w http.ResponseWriter, r *http.Request, data TData) error {
	return func(w http.ResponseWriter, r *http.Request, data TData) error {
		return template.Execute(w, data)
	}
}

func templatesResponseWriterFor[TData ~map[string]any](templates []*template.Template) func(w http.ResponseWriter, r *http.Request, data TData) error {
	return func(w http.ResponseWriter, r *http.Request, data TData) error {
		for _, tmpl := range templates {
			err := tmpl.Execute(w, data[tmpl.Name()])
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func jsonResponseWriter[TOut any](w http.ResponseWriter, r *http.Request, response TOut) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func htmlResponseWriter[TOut any](w http.ResponseWriter, r *http.Request, response TOut) error {
	w.Header().Set("Content-Type", "text/html")
	_, err := w.Write([]byte(fmt.Sprint(response)))
	return err
}

func getTemplateNames(templates []*template.Template) string {
	names := []string{}
	for _, tmpl := range templates {
		names = append(names, tmpl.Name())
	}
	return strings.Join(names, ", ")
}

func redirectResponseWriter(w http.ResponseWriter, r *http.Request, url string) error {
	http.Redirect(w, r, url, http.StatusSeeOther)
	return nil
}
