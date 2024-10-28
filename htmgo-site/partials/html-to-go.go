package partials

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/tools/html-to-htmgo/htmltogo"
	"htmgo-site/ui"
)

func ConvertHtmlToGo(ctx *h.RequestContext) *h.Partial {
	value := ctx.FormValue("html-input")
	parsed := string(htmltogo.Parse([]byte(value)))

	formatted := ui.FormatCode(parsed, "height: 100%;")

	return h.SwapManyPartial(ctx,
		GoOutput(formatted),
		HiddenCopyOutput(parsed),
	)
}

func HtmlInput() *h.Element {
	return h.Div(
		h.Class("h-full w-1/2 min-w-1/2"),
		h.TextArea(
			h.Name("html-input"),
			h.MaxLength(500*1000),
			h.PostPartial(ConvertHtmlToGo, "keyup delay:300ms"),
			h.Class("h-[90%] w-full p-4 rounded border border-slate-200"),
			h.Placeholder("Paste your HTML here"),
			h.Rows(10),
		),
	)
}

func HiddenCopyOutput(content string) *h.Element {
	return h.Div(
		h.Class("hidden"),
		h.Id("go-output-raw"),
		h.UnsafeRaw(content),
	)
}

func GoOutput(content string) *h.Element {
	return h.Div(
		h.Class("h-full w-1/2 min-w-1/2"),
		h.Id("go-output"),
		h.Div(
			h.Class("h-[90%] w-full rounded border border-slate-200 relative"),
			h.Div(
				h.Class("h-full"),
				h.Id("go-output-content"),
				h.UnsafeRaw(content),
			),
			h.If(
				content != "",
				ui.CopyButton("#go-output-raw"),
			),
		),
	)
}
