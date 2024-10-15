package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/pages/base"
	"htmgo-site/partials"
)

func HtmlToGoPage(ctx *h.RequestContext) *h.Page {
	return h.NewPage(
		base.PageWithNav(ctx,
			h.Div(
				h.Class("flex flex-col h-full items-center justify-center w-full pt-6"),
				h.H3(
					h.Text("Convert raw html to htmgo code"),
					h.Class("text-2xl font-bold"),
				),
				h.Div(
					h.Class("h-full w-full flex gap-4 p-8"),
					partials.HtmlInput(),
					partials.GoOutput(""),
				),
			),
		),
	)
}
