package base

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/partials"
	"strings"
)

func Extensions() string {
	extensions := []string{"path-deps", "response-targets", "mutation-error"}
	if h.IsDevelopment() {
		extensions = append(extensions, "livereload")
	}
	return strings.Join(extensions, ", ")
}

func RootPage(children ...h.Ren) *h.Element {
	return h.Html(
		h.HxExtension(Extensions()),
		h.Head(
			h.Meta("viewport", "width=device-width, initial-scale=1"),
			h.Link("/public/main.css", "stylesheet"),
			h.Script("/public/htmgo.js"),
			h.Raw(`
				<script async defer src="https://buttons.github.io/buttons.js"></script>
			`),
		),
		h.Class("bg-neutral-50 min-h-screen overflow-x-hidden"),
		partials.NavBar(),
		h.Fragment(children...),
	)
}
