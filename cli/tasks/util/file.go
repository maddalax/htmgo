package util

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func ReplaceTextInFile(file string, text string, replacement string) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	str := string(bytes)
	updated := strings.ReplaceAll(str, text, replacement)
	return os.WriteFile(file, []byte(updated), 0644)
}

func ReplaceTextInDirRecursive(dir string, text string, replacement string, filter func(file string) bool) error {
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if filter(path) {
			_ = ReplaceTextInFile(path, text, replacement)
		}
		return nil
	})
}
