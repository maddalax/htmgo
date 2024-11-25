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
	"io"
	"net/http"
	"strings"
)

func FormatCode(props CodeSnippetProps) string {

	if props.SingleLine {
		props.CustomStyles = append(props.CustomStyles, "height: 50px; width: 100%;")
	}

	var buf bytes.Buffer
	lexer := lexers.Get(props.Lang)
	style := styles.Get("github")
	formatter := html.New(
		html.WrapLongLines(props.WrapLines),
		html.WithLineNumbers(!props.SingleLine && !props.HideLineNumbers),
		html.WithCustomCSS(map[chroma.TokenType]string{
			chroma.PreWrapper: fmt.Sprintf("border-radius: 0.2rem; line-height: 24px; font-size: 14px; padding: 12px; overflow: auto; background-color: rgb(245, 245, 245) !important; %s", strings.Join(props.CustomStyles, ";")),
		}))
	iterator, err := lexer.Tokenise(nil, props.Code)
	if err != nil {
		return ""
	}
	err = formatter.Format(&buf, style, iterator)
	return buf.String()
}

type CodeSnippetProps struct {
	Code            string
	Lang            string
	CustomStyles    []string
	HideLineNumbers bool
	SingleLine      bool
	WrapLines       bool
}

func CodeSnippet(props CodeSnippetProps) *h.Element {
	id := fmt.Sprintf("code-snippet-%s", uuid.NewString())

	props.Code = strings.TrimPrefix(props.Code, "\n")
	props.Code = strings.TrimSuffix(props.Code, "\n")

	if props.SingleLine {
		return h.Div(
			h.Class("flex items-center w-full"),
			h.Div(
				h.UnsafeRaw(props.Code),
				h.Class("hidden"),
				h.Id(id),
			),
			h.UnsafeRaw(
				FormatCode(props),
			),
			CopyButton("#"+id, "h-[50px] rounded-sm"),
		)
	}

	return h.Div(
		h.Class("relative max-w-[90vw]"),
		h.Div(
			h.UnsafeRaw(props.Code),
			h.Class("hidden"),
			h.Id(id),
		),
		AbsoluteCopyButton("#"+id),
		h.UnsafeRaw(
			FormatCode(props),
		),
	)
}

func BashCodeSnippet(code string, customStyles ...string) *h.Element {
	return CodeSnippet(CodeSnippetProps{
		Code:         code,
		Lang:         "bash",
		CustomStyles: customStyles,
	})
}

func SingleLineBashCodeSnippet(code string, customStyles ...string) *h.Element {
	return CodeSnippet(CodeSnippetProps{
		Code:         code,
		Lang:         "bash",
		CustomStyles: customStyles,
		SingleLine:   true,
	})
}

func GoCodeSnippet(code string, customStyles ...string) *h.Element {
	return CodeSnippet(CodeSnippetProps{
		Code:         code,
		Lang:         "go",
		CustomStyles: customStyles,
	})
}

func GoCodeSnippetSingleLine(code string, customStyles ...string) *h.Element {
	return CodeSnippet(CodeSnippetProps{
		Code:         code,
		Lang:         "go",
		CustomStyles: customStyles,
		SingleLine:   true,
	})
}

func CodeSnippetFromUrl(url string, props CodeSnippetProps) *h.Element {
	data, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return h.Empty()
	}
	defer data.Body.Close()
	b, err := io.ReadAll(data.Body)
	if err != nil {
		return h.Empty()
	}
	props.Code = string(b)
	return CodeSnippet(props)
}
