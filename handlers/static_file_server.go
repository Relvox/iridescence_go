package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type StaticFileServer string

func (s StaticFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir(s)).ServeHTTP(w, r)
}

func (s StaticFileServer) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/").Handler(s)
}
