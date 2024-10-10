package pages

import (
	"github.com/maddalax/htmgo/framework/h"
)

func RootPage(children ...h.Ren) h.Ren {
	return h.Html(
		h.HxExtensions(h.BaseExtensions()),
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
			h.Div(
				h.Fragment(children...),
			),
		),
	)
}
