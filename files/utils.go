package files

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func GetFilenames(directory string, extension string) ([]string, error) {
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), extension) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func Split(file string) (dir string, name string, ext string) {
	dir = filepath.Dir(file)
	ext = filepath.Ext(file)
	name = strings.TrimSuffix(filepath.Base(file), ext)
	return
}

func IsolateName(file string) string {
	return strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
}

func ListFS(fsys fs.FS, root string) ([]string, error) {
	var result []string
	err := fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			result = append(result, path)
		}
		return nil
	})
	return result, err
}
