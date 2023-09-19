//go:build !debug
// +build !debug

package sawmill

import (
	"embed"

	"github.com/relvox/iridescence_go/files"
	"github.com/relvox/iridescence_go/handlers"
)

//go:embed embeds/static/* embeds/templates/*
var staticFS embed.FS

var SawmillStaticFSDir = files.NewSubdirectoryFS(staticFS, "embeds/static")

var SawmillTemplateFS = files.NewSubdirectoryFS(staticFS, "embeds/templates")

var SawmillTemplateHandler = handlers.
	NewFSTemplateHandler(SawmillTemplateFS)
