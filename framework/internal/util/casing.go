package util

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile("([a-z])([A-Z])")

// ConvertCamelToDash converts a camelCase string to dash-case
func ConvertCamelToDash(s string) string {
	// Find uppercase letters and prepend a dash
	s = re.ReplaceAllString(s, "$1-$2")
	// Convert the string to lower case
	return strings.ToLower(s)
}
