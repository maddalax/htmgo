package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/pages/base"
)

func IndexPage(ctx *h.RequestContext) *h.Page {
	return h.NewPage(
		base.RootPage(h.Div(
			h.Class("flex items-center justify-center"),
			h.Div(
				h.Class("w-full px-4 flex flex-col prose max-w-[95vw] md:max-w-3xl mt-6"),
				h.Div(
					h.Class("flex flex-col mb-6 md:mb-0 md:flex-row justify-between items-center"),
					h.Div(
						h.H1(
							h.Class("text-center md:text-left"),
							h.Text("htmgo"),
						),
						h.H3(
							h.Class("-mt-4"),
							h.TextF("build simple and scalable systems with go + htmx"),
						),
					),
					h.Div(
						h.A(
							h.Href("/docs"),
							h.Class("not-prose p-3 bg-slate-900 text-white rounded-md"),
							h.Text("Get Started"),
						),
					),
				),
				h.Div(
					h.Class("border-b border-b-slate-200 h-1"),
					h.Div(
						h.Class("mt-4"),
						MarkdownPage(ctx, "md/index.md", ""),
					),
				),
			)),
		),
	)
}
