package ui

import (
	"bytes"
	"fmt"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
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
			chroma.PreWrapper: fmt.Sprintf("padding: 12px; overflow: auto; background-color: rgb(245, 245, 245) !important; %s", strings.Join(customStyles, ";")),
		}))
	iterator, err := lexer.Tokenise(nil, code)
	if err != nil {
		return ""
	}
	err = formatter.Format(&buf, style, iterator)
	return buf.String()
}

func CodeSnippet(code string, customStyles ...string) *h.Element {
	return h.Div(
		h.UnsafeRaw(FormatCode(code, customStyles...)),
	)
}
