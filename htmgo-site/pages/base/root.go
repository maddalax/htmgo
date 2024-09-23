package base

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/partials"
)

func RootPage(children ...h.Ren) *h.Element {
	return h.Html(
		h.HxExtension(h.BaseExtensions()),
		h.Head(
			h.Meta("viewport", "width=device-width, initial-scale=1"),
			h.Link("/public/main.css", "stylesheet"),
			h.Script("/public/htmgo.js"),
			h.Raw(`
				<script async defer src="https://buttons.github.io/buttons.js"></script>
			`),
			h.Style(`
				html {
					scroll-behavior: smooth;
				}
			`),
		),
		h.Body(
			h.Class("bg-neutral-50 min-h-screen overflow-x-hidden"),
			partials.NavBar(false),
			h.Fragment(children...),
		),
	)
}
