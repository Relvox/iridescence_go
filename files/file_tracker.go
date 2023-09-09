package files

import (
	"os"
	"sync"
	"time"
)

type FileTracker struct {
	FilePath string

	mu            sync.RWMutex
	content       []byte
	LastWriteTime time.Time
}

func NewFileTracker(filepath string) *FileTracker {
	result := &FileTracker{
		FilePath:      filepath,
		mu:            sync.RWMutex{},
		content:       []byte{},
		LastWriteTime: time.Time{},
	}
	if err := result.Refresh(); err != nil {
		panic(err)
	}
	return result
}

func (fw *FileTracker) Refresh() error {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	info, err := os.Stat(fw.FilePath)
	if err != nil {
		return err
	}

	if info.ModTime().Equal(fw.LastWriteTime) {
		return nil
	}

	content, err := os.ReadFile(fw.FilePath)
	if err != nil {
		return err
	}

	fw.content = content
	fw.LastWriteTime = info.ModTime()
	return nil
}

func (fw *FileTracker) GetContent() []byte {
	fw.mu.RLock()
	defer fw.mu.RUnlock()
	return fw.content
}
