//go:build debug
// +build debug

package main

import (
	"io/fs"
	"os"

	"github.com/relvox/iridescence_go/handlers"
)

var staticFileServer = handlers.StaticFileServer("./static/")

var getTemplatesFS = func() fs.FS { return os.DirFS("./templates/") }
