package util

import (
	"os"
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
