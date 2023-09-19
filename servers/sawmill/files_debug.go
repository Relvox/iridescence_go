//go:build debug
// +build debug

package sawmill

import (
	"os"

	"github.com/relvox/iridescence_go/handlers"
)

var SawmillStaticFSDir = os.DirFS("./embeds/static/")

var SawmillTemplateFS = os.DirFS("./embeds/templates/")

var SawmillTemplateHandler = handlers.
	NewFSTemplateHandler(SawmillTemplateFS).
	WithDebug()
