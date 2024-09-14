package base

import (
	"github.com/maddalax/htmgo/framework/h"
	"starter-template/partials"
	"starter-template/partials/sheet"
)

func RootPage(children ...h.Renderable) h.Renderable {
	return h.Html(
		h.HxExtension("path-deps, response-targets, mutation-error"),
		h.Head(
			h.Link("/public/main.css", "stylesheet"),
			h.Script("/public/htmgo.js"),
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
