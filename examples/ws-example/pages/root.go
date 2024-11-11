package pages

import (
	"github.com/maddalax/htmgo/framework/h"
)

func RootPage(ctx *h.RequestContext, children ...h.Ren) h.Ren {
	return h.Html(
		h.JoinExtensions(
			h.HxExtension(
				h.BaseExtensions(),
			),
			h.HxExtension("ws"),
		),
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
