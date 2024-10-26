package pages

import (
	"github.com/maddalax/htmgo/framework/h"
)

func RootPage(children ...h.Ren) h.Ren {
	banner := h.A(
		h.Class("bg-neutral-200 text-neutral-600 text-center p-2 flex items-center justify-center"),
		h.Href("https://github.com/maddalax/htmgo"),
		h.Attribute("target", "_blank"),
		h.Text("Built with htmgo.dev"),
	)

	return h.Html(
		h.HxExtensions(
			h.BaseExtensions(),
		),
		h.Head(
			h.Meta("viewport", "width=device-width, initial-scale=1"),
			h.Link("/public/favicon.ico", "icon"),
			h.Link("/public/apple-touch-icon.png", "apple-touch-icon"),
			h.Meta("title", "hackernews"),
			h.Meta("charset", "utf-8"),
			h.Meta("author", "htmgo"),
			h.Meta("description", "hacker news reader, built with htmgo"),
			h.Meta("og:title", "hacker news reader"),
			h.Meta("og:url", "https://hn.htmgo.dev"),
			h.Link("canonical", "https://hn.htmgo.dev"),
			h.Meta("og:description", "hacker news reader, built with htmgo"),
			h.Link("/public/main.css", "stylesheet"),
			h.Script("/public/htmgo.js"),
		),
		h.Body(
			banner,
			h.Div(
				h.Fragment(children...),
			),
		),
	)
}
