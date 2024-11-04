package docs

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/ui"
)

const IntroSnippet = `func DocsPage(ctx *h.RequestContext) *h.Page {
	pages := dirwalk.WalkPages("md/docs")
	return h.NewPage(
	h.Div(
		h.Class("flex flex-col md:flex-row gap-4"),
		DocSidebar(pages),
		h.Div(
			h.Class("flex flex-col justify-center items-center mt-6"),
			h.List(pages, func(page *dirwalk.Page, index int) *h.Element {
				return h.Div(
					h.Class("border-b border-b-slate-300"),
					MarkdownContent(ctx, page),
				)
			}),
		 ),
	  ),
}`

func Introduction(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-2"),
			Title("Introduction"),
			Text(`
				htmgo is a lightweight pure go way to build interactive websites / web applications using go & htmx.
				We give you the utilities to build html using pure go code in a reusable way (go functions are components) while also providing htmx functions to add interactivity to your app.
			`),
			ui.GoCodeSnippet(IntroSnippet),
			h.P(
				Link("The site you are reading now", "https://github.com/maddalax/htmgo/tree/master/htmgo-site"),
				h.Text(" was written with htmgo!"),
			),
			NextStep(
				"mt-4",
				h.Div(),
				NextBlock("Getting Started", DocPath("/installation")),
			),
		),
	)
}
