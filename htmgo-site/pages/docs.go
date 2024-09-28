package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/internal/dirwalk"
	"htmgo-site/pages/base"
	"htmgo-site/partials"
	"io/fs"
)

func DocsPage(ctx *h.RequestContext) *h.Page {
	assets := ctx.Get("embeddedMarkdown").(fs.FS)
	pages := dirwalk.WalkPages("md/docs", assets)

	return h.NewPage(base.RootPage(
		ctx,
		h.Div(
			h.Class("flex flex-col md:flex-row gap-6 justify-center"),
			h.Aside(
				h.Class("md:h-screen md:sticky md:top-0 md:w-42"), // Applied sticky positioning here
				partials.DocSidebar(pages),
			),
			h.Main(
				h.Class("md:flex gap-4 justify-center mb-6"),
				h.Div(
					h.Class("flex flex-col"),
					h.Div(
						h.Class("flex flex-col justify-center items-center md:mt-6 "),
						h.List(pages, func(page *dirwalk.Page, index int) *h.Element {
							anchor := partials.CreateAnchor(page.Parts)
							return h.Div(
								h.Class("border-b border-b-slate-300 w-full pb-8 p-4 md:px-0 -mb-2"),
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
						h.Class("flex justify-center items-center mt-6"),
						h.A(
							h.Text("Back to Top"),
							h.Class("py-2 px-3 bg-slate-800 rounded text-white"),
							h.Href("#"),
						),
					),
				),
			),
		),
	))
}
