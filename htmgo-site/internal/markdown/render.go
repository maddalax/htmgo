package markdown

import (
	"bytes"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"io"
	"os"
)

type Renderer struct {
	cache map[string]string
}

func NewRenderer() *Renderer {
	return &Renderer{cache: make(map[string]string)}
}

func (r *Renderer) RenderFile(source string) string {
	if val, ok := r.cache[source]; ok {
		return val
	}

	o, err := os.Open(source)
	defer func(o *os.File) {
		_ = o.Close()
	}(o)

	if err != nil {
		return ""
	}

	buf := RenderMarkdown(o)
	r.cache[source] = buf.String()
	return r.cache[source]
}

func RenderMarkdown(reader io.Reader) bytes.Buffer {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("github"),
			),
		),
	)

	source, _ := io.ReadAll(reader)
	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		panic(err)
	}

	return buf
}
