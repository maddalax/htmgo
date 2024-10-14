package dirutil

import (
	"fmt"
	"github.com/bmatcuk/doublestar/v4"
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
