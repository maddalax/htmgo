package base

import (
	"github.com/maddalax/htmgo/framework/h"
)

func RootPage(children ...h.Ren) h.Ren {
	return h.Html(
		h.HxExtension(h.BaseExtensions()),
		h.Head(
			h.Link("/public/main.css", "stylesheet"),
			h.Script("/public/htmgo.js"),
		),
		h.Fragment(children...),
	)
}
