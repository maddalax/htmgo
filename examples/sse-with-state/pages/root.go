package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"starter-template/state"
)

func RootPage(ctx *h.RequestContext, children ...h.Ren) h.Ren {
	s := state.NewState(ctx)
	return h.Html(
		h.Attribute("data-session-id", s.SessionId),
		h.HxExtension(h.BaseExtensions()),
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
