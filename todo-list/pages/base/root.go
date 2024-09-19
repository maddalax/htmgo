package base

import (
	"github.com/maddalax/htmgo/framework/h"
	"strings"
)

func Extensions() string {
	extensions := []string{"path-deps", "response-targets", "mutation-error"}
	if h.IsDevelopment() {
		extensions = append(extensions, "livereload")
	}
	return strings.Join(extensions, ", ")
}

func RootPage(children ...h.Renderable) h.Renderable {
	return h.Html(
		h.HxExtension(Extensions()),
		h.Head(
			h.Link("/public/main.css", "stylesheet"),
			h.Script("/public/htmgo.js"),
		),
		h.Fragment(children...),
	)
}
