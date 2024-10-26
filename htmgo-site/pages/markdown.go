package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"htmgo-site/internal/markdown"
	"io/fs"
)

func MarkdownPage(ctx *h.RequestContext, path string, id string) *h.Element {
	return h.Div(
		MarkdownContent(ctx, path, id),
		h.Div(
			h.Class("min-h-12"),
		),
	)
}

func MarkdownContent(ctx *h.RequestContext, path string, id string) *h.Element {
	embeddedMd := ctx.Get("embeddedMarkdown").(fs.FS)
	renderer := service.Get[markdown.Renderer](ctx.ServiceLocator())
	return h.Div(
		h.If(
			id != "",
			h.Id(id),
		),
		h.Div(
			h.Class("w-full flex flex-col prose max-w-[90vw] md:max-w-xl md:px-4 lg:max-w-4xl xl:max-w-5xl prose-code:text-black prose-p:my-1 prose:p-0 prose-li:m-0 prose-ul:m-0 prose-ol:m-0"),
			h.UnsafeRaw(renderer.RenderFile(path, embeddedMd)),
		),
	)
}
