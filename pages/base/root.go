package base

import (
	"mhtml/h"
	"mhtml/partials"
	"mhtml/partials/sheet"
)

func RootPage(children ...h.Renderable) h.Renderable {
	return h.Html(
		h.HxExtension("path-deps, response-targets, mutation-error"),
		h.Head(
			h.Link("/public/main.css", "stylesheet"),
			h.Script("/public/mhtml.js"),
		),
		h.Body(
			partials.NavBar(),
			sheet.Closed(),
			h.Div(
				h.Class("flex flex-col gap-2 bg-white h-full"),
				h.Fragment(children...),
			),
		),
	)
}
