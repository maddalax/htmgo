package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/htmgo/service"
	"htmgo-site/internal/markdown"
	"htmgo-site/pages/base"
)

func MarkdownHandler(ctx *h.RequestContext, path string) error {
	return h.HtmlView(ctx, h.NewPage(MarkdownPage(ctx, path)))
}

func MarkdownPage(ctx *h.RequestContext, path string) *h.Element {
	return base.RootPage(
		h.Div(
			h.Class("flex flex-col p-4 justify-center items-center"),
			MarkdownContent(ctx, path),
			h.Div(
				h.Class("min-h-12"),
			),
		),
	)
}

func MarkdownContent(ctx *h.RequestContext, path string) *h.Element {
	renderer := service.Get[markdown.Renderer](ctx.ServiceLocator())
	return h.Div(
		h.Article(
			h.Class("prose max-w-sm pt-3 p-4 md:p-4 md:max-w-2xl prose-code:text-black"),
			h.Raw(renderer.RenderFile(path)),
		),
	)
}
