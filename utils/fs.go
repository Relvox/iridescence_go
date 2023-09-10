package utils

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type SubdirectoryFS struct {
	fs.FS
	Subdirectory string
}

func NewSubdirectoryFS(src fs.FS, subDir string) *SubdirectoryFS {
	return &SubdirectoryFS{
		FS:           src,
		Subdirectory: subDir,
	}
}

func (s SubdirectoryFS) Open(name string) (fs.File, error) {
	newPath := filepath.Join(s.Subdirectory, filepath.Clean("/"+name))
	newPath = strings.ReplaceAll(newPath, "\\", "/")
	return s.FS.Open(newPath)
}
