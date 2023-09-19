package sawmill

import (
	"io"
	"io/fs"
	"log"
	"slices"
	"strings"
	"sync"
	"time"
)

type FSLineTracker struct {
	FS   fs.FS
	Path string

	lock          sync.RWMutex
	Content       []string
	LastWriteTime time.Time
}

func NewFSLineTracker(fs fs.FS, filepath string) *FSLineTracker {
	result := &FSLineTracker{
		FS:   fs,
		Path: filepath,

		lock:          sync.RWMutex{},
		Content:       make([]string, 0),
		LastWriteTime: time.Time{},
	}
	if err := result.Refresh(); err != nil {
		panic(err)
	}
	return result
}

func (f *FSLineTracker) Refresh() error {
	f.lock.Lock()
	defer f.lock.Unlock()

	file, err := f.FS.Open(f.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	if info.ModTime().Equal(f.LastWriteTime) {
		return nil
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	f.Content = strings.Split(string(data), "\n")
	slices.Reverse(f.Content)
	f.LastWriteTime = info.ModTime()
	return nil
}
