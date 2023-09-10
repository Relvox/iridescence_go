package handlers

import (
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

type FSFileServer struct {
	fs.FS
}

func (s FSFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.FS(s)).ServeHTTP(w, r)
}

func (s FSFileServer) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/").Handler(s)
}
