//go:build !debug
// +build !debug

package main

import (
	"embed"
	"io/fs"

	"github.com/relvox/iridescence_go/files"
	"github.com/relvox/iridescence_go/handlers"
)

//go:embed static/* templates/*
var staticFS embed.FS

var staticFileServer = handlers.FSFileServer{
	FS: files.NewSubdirectoryFS(staticFS, "static"),
}

var getTemplatesFS = func() fs.FS { return files.NewSubdirectoryFS(staticFS, "templates") }
