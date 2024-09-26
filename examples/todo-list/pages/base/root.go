package base

import (
	"github.com/maddalax/htmgo/framework/h"
)

func RootPage(children ...h.Ren) h.Ren {
	return h.Html(
		h.HxExtension(h.BaseExtensions()),
		h.Head(
			h.Meta("viewport", "width=device-width, initial-scale=1"),
			h.Meta("title", "htmgo todo mvc"),
			h.Meta("description", "an example of how to build a todo mvc app with htmgo"),
			h.Meta("charset", "utf-8"),
			h.Link("/public/main.css", "stylesheet"),
			h.Script("/public/htmgo.js"),
		),
		h.Fragment(children...),
	)
}
