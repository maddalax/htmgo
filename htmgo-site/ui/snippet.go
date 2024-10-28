package ui

import (
	"bytes"
	"fmt"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/google/uuid"
	"github.com/maddalax/htmgo/framework/h"
	"strings"
)

func FormatCode(code string, customStyles ...string) string {
	var buf bytes.Buffer
	lexer := lexers.Get("go")
	style := styles.Get("github")
	formatter := html.New(
		html.WrapLongLines(true),
		html.WithLineNumbers(true),
		html.WithCustomCSS(map[chroma.TokenType]string{
			chroma.PreWrapper: fmt.Sprintf("font-size: 14px; padding: 12px; overflow: auto; background-color: rgb(245, 245, 245) !important; %s", strings.Join(customStyles, ";")),
		}))
	iterator, err := lexer.Tokenise(nil, code)
	if err != nil {
		return ""
	}
	err = formatter.Format(&buf, style, iterator)
	return buf.String()
}

func CodeSnippet(code string, customStyles ...string) *h.Element {
	id := fmt.Sprintf("code-snippet-%s", uuid.NewString())
	return h.Div(
		h.Class("relative"),
		h.Div(
			h.UnsafeRaw(code),
			h.Class("hidden"),
			h.Id(id),
		),
		CopyButton("#"+id),
		h.UnsafeRaw(
			FormatCode(code, customStyles...),
		),
	)
}
