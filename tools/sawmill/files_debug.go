//go:build debug
// +build debug

package main

import (
	"os"

	"github.com/relvox/iridescence_go/handlers"
)

var staticFSDir = os.DirFS("./embeds/static/")

var templateFS = os.DirFS("./embeds/templates/")

var templateHandler = handlers.
	NewFSTemplateHandler(templateFS).
	WithDebug()
