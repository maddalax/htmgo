package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/internal/dirwalk"
	"htmgo-site/pages/base"
)

func DocsPage(ctx *h.RequestContext) *h.Page {
	pages := dirwalk.WalkPages("md/docs")

	return h.NewPage(base.RootPage(
		h.Div(
			h.Class("flex flex-col p-4 justify-center items-center"),
			h.List(pages, func(page dirwalk.Page, index int) *h.Element {
				return MarkdownContent(ctx, page.FilePath)
			}),
		),
		h.Div(
			h.Class("min-h-12"),
		),
	),
	)
}
