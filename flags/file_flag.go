package flags

import (
	"fmt"
	"os"
)

// FileFlagValue holds the content of the file provided as a flag.
type FileFlagValue string

// String is the content of the file.
func (ffv *FileFlagValue) String() string {
	return string(*ffv)
}

// Set opens the file, reads its content, and stores it in the type.
func (ffv *FileFlagValue) Set(value string) error {
	*ffv = FileFlagValue(value)
	_, err := os.ReadFile(string(*ffv))
	if err != nil {
		return fmt.Errorf("read file '%s': %w", *ffv, err)
	}
	return nil
}

// Read opens the file and returns its content
func (ffv *FileFlagValue) Read() []byte {
	content, _ := os.ReadFile(string(*ffv))
	return content
}
