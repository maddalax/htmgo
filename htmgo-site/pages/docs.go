package pages

import (
	"embed"
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/internal/dirwalk"
	"htmgo-site/pages/base"
	"htmgo-site/partials"
)

func DocsPage(ctx *h.RequestContext) *h.Page {
	assets := ctx.Get("embeddedMarkdown").(*embed.FS)
	pages := dirwalk.WalkPages("md/docs", assets)

	return h.NewPage(base.RootPage(
		h.Div(
			h.Class("flex flex-col md:flex-row gap-4 justify-center mb-12"),
			partials.DocSidebar(pages),
			h.Div(
				h.Class("flex flex-col justify-center items-center md:mt-6"),
				h.List(pages, func(page *dirwalk.Page, index int) *h.Element {
					anchor := partials.CreateAnchor(page.Parts)
					return h.Div(
						h.Class("border-b border-b-slate-300 w-full pb-8 mb-8 p-4 md:px-0 -mb-2"),
						MarkdownContent(ctx, page.FilePath, anchor),
						h.Div(
							h.Class("ml-4 pl-1 mt-2 bg-rose-200"),
							h.If(anchor == "core-concepts-partials",
								h.GetPartial(partials.CurrentTimePartial, "load, every 1s"),
							),
						),
					)
				}),
			),
			h.Div(
				h.Class("min-h-12"),
			),
		),
	))
}
