package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"starter-template/partials"
)

func IndexPage(ctx *h.RequestContext) *h.Page {
	return RootPage(
		h.Div(
			h.Class("flex flex-col gap-4 items-center pt-24 min-h-screen bg-neutral-100"),
			h.H3(
				h.Id("intro-text"),
				h.Text("hello htmgo"),
				h.Class("text-5xl"),
			),
			h.Div(
				h.Class("mt-3"),
				partials.CounterForm(0),
			),
		),
	)
}
