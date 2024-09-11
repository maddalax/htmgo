package base

import (
	"mhtml/h"
	"mhtml/partials"
	"mhtml/partials/sheet"
)

func RootPage(children ...h.Renderable) h.Renderable {
	return h.Html(
		h.Head(
			h.Script("https://cdn.tailwindcss.com"),
			h.Script("https://unpkg.com/htmx.org@2.0.2"),
			h.Script("/js/mhtml.js"),
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
