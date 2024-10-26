package pages

import (
	"github.com/maddalax/htmgo/framework/h"
)

func RootPage(children ...h.Ren) h.Ren {
	return h.Html(
		h.HxExtensions(
			h.BaseExtensions(),
		),
		h.Head(
			h.Meta("viewport", "width=device-width, initial-scale=1"),
			h.Link("/public/favicon.ico", "icon"),
			h.Link("/public/apple-touch-icon.png", "apple-touch-icon"),
			h.Meta("title", "htmgo template"),
			h.Meta("charset", "utf-8"),
			h.Meta("author", "htmgo"),
			h.Meta("description", "this is a template"),
			h.Meta("og:title", "htmgo template"),
			h.Meta("og:url", "https://htmgo.dev"),
			h.Link("canonical", "https://htmgo.dev"),
			h.Meta("og:description", "this is a template"),
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
