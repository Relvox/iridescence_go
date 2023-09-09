package servers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type StaticFileServer string

func (s *StaticFileServer) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir(*s)).ServeHTTP(w, r)
	}))
}
