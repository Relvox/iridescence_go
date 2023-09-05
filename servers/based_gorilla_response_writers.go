package servers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func redirectResponseWriter(w http.ResponseWriter, r *http.Request, url string) error {
	http.Redirect(w, r, url, http.StatusSeeOther)
	return nil
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
