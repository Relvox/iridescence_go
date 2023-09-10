//go:build !debug
// +build !debug

package main

import (
	"embed"
	"io/fs"

	"github.com/relvox/iridescence_go/handlers"
	"github.com/relvox/iridescence_go/utils"
)

//go:embed static/* templates/*
var staticFS embed.FS

var staticFileServer = handlers.FSFileServer{
	FS: utils.NewSubdirectoryFS(staticFS, "static"),
}

var getTemplatesFS = func() fs.FS { return utils.NewSubdirectoryFS(staticFS, "templates") }
