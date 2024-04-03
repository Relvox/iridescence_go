package files

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func Split(file string) (dir string, name string, ext string) {
	dir = filepath.Dir(file)
	ext = filepath.Ext(file)
	name = strings.TrimSuffix(filepath.Base(file), ext)
	return
}

func IsolateName(file string) string {
	return strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
}

// ListFS
//
// Deprecated: moved to utilgo
func ListFS(fsys fs.FS, root, glob string) ([]string, error) {
	var result []string
	err := fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			if ok, err := filepath.Match(glob, path); !ok || err != nil {
				return err
			}
			result = append(result, path)
		}
		return nil
	})
	return result, err
}
