//go:build debug
// +build debug

package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/relvox/iridescence_go/handlers"
	"github.com/relvox/iridescence_go/templates/funcs"
)

var staticFileHandler = http.FileServer(http.FS(
	os.DirFS("./embeds/static/"),
))

var templateFS = os.DirFS("./embeds/templates/")

var templateHandler = handlers.
	NewFSTemplateHandler(templateFS).
	WithFuncs(template.FuncMap{"contains": funcs.SliceContains}).
	WithDebug()
