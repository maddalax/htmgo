package dirutil

import (
	"fmt"
	"github.com/bmatcuk/doublestar/v4"
	"io/fs"
	"path/filepath"
)

func matchesAny(patterns []string, path string) bool {
	for _, pattern := range patterns {
		matched, err := doublestar.Match(pattern, path)
		if err != nil {
			fmt.Printf("Error matching pattern: %v\n", err)
			return false
		}
		if matched {
			return true
		}
	}
	return false
}

func IsGlobExclude(path string, excludePatterns []string) bool {
	return matchesAny(excludePatterns, path)
}

func IsGlobMatch(path string, patterns []string, excludePatterns []string) bool {
	if matchesAny(excludePatterns, path) {
		return false
	}
	return matchesAny(patterns, path)
}

func GlobMatchDirs(root string, includePatterns, excludePatterns []string, cb func(string)) {
	//directories := map[string]bool{}
	// Walk through the directory recursively
	_ = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if matchesAny(excludePatterns, path) {
			return fs.SkipDir
		}
		if matchesAny(includePatterns, path) {
			cb(path)
		}
		return nil
	})
}
