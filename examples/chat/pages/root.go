package pages

import (
	"github.com/maddalax/htmgo/framework/h"
)

func RootPage(children ...h.Ren) h.Ren {
	extensions := h.BaseExtensions()
	return h.Html(
		h.HxExtension(extensions),
		h.Meta("viewport", "width=device-width, initial-scale=1"),
		h.Meta("title", "htmgo chat example"),
		h.Meta("charset", "utf-8"),
		h.Meta("author", "htmgo"),
		h.Head(
			h.Link("/public/main.css", "stylesheet"),
			h.Script("/public/htmgo.js"),
		),
		h.Body(
			h.Div(
				h.Class("flex flex-col gap-2 bg-white h-full"),
				h.Fragment(children...),
			),
		),
	)
}
