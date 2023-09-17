//go:build !debug
// +build !debug

package main

import (
	"embed"

	"github.com/relvox/iridescence_go/files"
	"github.com/relvox/iridescence_go/handlers"
)

//go:embed embeds/static/* embeds/templates/*
var staticFS embed.FS

var staticFSDir = files.NewSubdirectoryFS(staticFS, "embeds/static")

var templateFS = files.NewSubdirectoryFS(staticFS, "embeds/templates")

var templateHandler = handlers.
	NewFSTemplateHandler(templateFS)
