package partials

import (
	"bytes"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/tools/html-to-htmgo/htmltogo"
)

func ConvertHtmlToGo(ctx *h.RequestContext) *h.Partial {
	value := ctx.FormValue("html-input")
	parsed := htmltogo.Parse([]byte(value))

	var buf bytes.Buffer

	lexer := lexers.Get("go")
	style := styles.Get("github")
	formatter := html.New(html.WithCustomCSS(map[chroma.TokenType]string{
		chroma.PreWrapper: "padding: 12px; height: 100%; overflow: auto;",
	}))
	iterator, err := lexer.Tokenise(nil, string(parsed))
	err = formatter.Format(&buf, style, iterator)

	if err != nil {
		return h.SwapPartial(ctx, GoOutput(string(parsed)))
	}

	return h.SwapPartial(ctx, GoOutput(buf.String()))
}

func HtmlInput() *h.Element {
	return h.Div(
		h.Class("h-[90%] w-1/2 min-w-1/2"),
		h.TextArea(
			h.Name("html-input"),
			h.PostPartial(ConvertHtmlToGo, "keyup delay:300ms"),
			h.Class("h-[90%] w-full p-4 rounded border border-slate-200"),
			h.Placeholder("Paste your HTML here"),
			h.Rows(10),
		),
	)
}

func GoOutput(content string) *h.Element {
	return h.Div(
		h.Class("h-[90%] w-1/2 min-w-1/2"),
		h.Id("go-output"),
		h.Div(
			h.Class("h-[90%] w-full rounded border border-slate-200"),
			h.UnsafeRaw(content),
		),
	)
}
