package pages

import (
	"embed"
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/internal/dirwalk"
	"htmgo-site/pages/base"
	"htmgo-site/partials"
	"strings"
)

func DocsPage(ctx *h.RequestContext) *h.Page {
	assets := ctx.Get("embeddedMarkdown").(*embed.FS)
	pages := dirwalk.WalkPages("md/docs", assets)

	return h.NewPage(base.RootPage(
		h.Div(
			h.Class("flex gap-4 justify-center mb-12"),
			partials.DocSidebar(pages),
			h.Div(
				h.Class("flex flex-col justify-center items-center mt-6 gap-12"),
				h.List(pages, func(page *dirwalk.Page, index int) *h.Element {
					return MarkdownContent(ctx, page.FilePath, strings.Join(page.Parts, "-"))
				}),
			),
			h.Div(
				h.Class("min-h-12"),
			),
		),
	))
}
