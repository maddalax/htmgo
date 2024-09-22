package h

import "strings"

func BaseExtensions() string {
	extensions := []string{"path-deps", "response-targets", "mutation-error", "htmgo"}
	if IsDevelopment() {
		extensions = append(extensions, "livereload")
	}
	return strings.Join(extensions, ", ")
}
