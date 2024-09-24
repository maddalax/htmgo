package pages

import (
	"embed"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"htmgo-site/internal/markdown"
	"htmgo-site/pages/base"
)

func MarkdownHandler(ctx *h.RequestContext, path string, id string) error {
	return h.HtmlView(ctx, h.NewPage(MarkdownPage(ctx, path, id)))
}

func MarkdownPage(ctx *h.RequestContext, path string, id string) *h.Element {
	return base.RootPage(
		h.Div(
			h.Class("w-full p-4 flex flex-col justify-center items-center"),
			MarkdownContent(ctx, path, id),
			h.Div(
				h.Class("min-h-12"),
			),
		),
	)
}

func MarkdownContent(ctx *h.RequestContext, path string, id string) *h.Element {
	embeddedMd := ctx.Get("embeddedMarkdown").(*embed.FS)
	renderer := service.Get[markdown.Renderer](ctx.ServiceLocator())
	return h.Div(
		h.If(id != "", h.Id(id)),
		h.Article(
			h.Class("prose max-w-[95vw] md:max-w-3xl px-4 prose-code:text-black prose-p:my-1"),
			h.Raw(renderer.RenderFile(path, embeddedMd)),
		),
	)
}
