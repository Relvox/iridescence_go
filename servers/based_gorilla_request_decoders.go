package servers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"gopkg.in/yaml.v3"
)

func nullDecoder[TIn any](r *http.Request) (TIn, error) {
	var none TIn
	return none, nil
}

func jsonRequestDecoder[TIn any](r *http.Request) (TIn, error) {
	var request TIn
	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	return request, err
}

func yamlRequestDecoder[TIn any](r *http.Request) (TIn, error) {
	var request TIn
	err := yaml.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	return request, err
}

func formRequestDecoder[TIn any](r *http.Request) (TIn, error) {
	var request TIn
	err := r.ParseForm()
	if err != nil {
		return request, err
	}
	err = schema.NewDecoder().Decode(&request, r.PostForm)
	return request, err
}
