package base

import (
	"mhtml/h"
)

func RootPage(children ...*h.Node) *h.Node {
	return h.Html(
		h.Head(
			h.Script("https://cdn.tailwindcss.com"),
			h.Script("https://unpkg.com/htmx.org@1.9.12"),
			h.Script("https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js"),
			h.Script("/js/mhtml.js"),
		),
		h.Body(
			h.VStack(
				h.Class("flex flex-col gap-2 bg-gray-100 h-full"),
				h.Fragment(children...),
			),
		),
	)
}
