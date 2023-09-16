//go:build !debug
// +build !debug

package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/relvox/iridescence_go/files"
	"github.com/relvox/iridescence_go/handlers"
	"github.com/relvox/iridescence_go/templates/funcs"
)

//go:embed embeds/static/* embeds/templates/*
var staticFS embed.FS

var staticFileHandler = http.FileServer(http.FS(
	files.NewSubdirectoryFS(staticFS, "embeds/static"),
))

var templateFS = files.NewSubdirectoryFS(staticFS, "embeds/templates")

var templateHandler = handlers.
	NewFSTemplateHandler(templateFS).
	WithFuncs(template.FuncMap{"contains": funcs.SliceContains})
