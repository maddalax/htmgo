package pages

import (
	"github.com/maddalax/htmgo/framework/h"
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
			),
			h.Div(),
		),
	)
}

func TestPartial(ctx *h.RequestContext) *h.Partial {
	return h.NewPartial(
		h.Div(
			h.Text("Hello World"),
		),
	)
}
