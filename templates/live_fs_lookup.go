package templates

import (
	"html/template"
	"io/fs"

	"golang.org/x/exp/slices"
)

// LiveLookup bla bla
//
// Deprecated: foo
func LiveLookup(funcMap template.FuncMap, getFS func() fs.FS, files ...string) []*template.Template {
	root := template.New("root").Funcs(funcMap)
	root = template.Must(root.ParseFS(getFS(), "*"))
	var result []*template.Template
	for _, t := range root.Templates() {
		if !slices.Contains(files, t.Name()) {
			continue
		}
		result = append(result, t)
	}
	return result
}
