package markdown

import (
	"bytes"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"io"
	"io/fs"
)

type Renderer struct {
	cache map[string]string
}

func NewRenderer() *Renderer {
	return &Renderer{cache: make(map[string]string)}
}

func (r *Renderer) RenderFile(source string, system fs.FS) string {
	if val, ok := r.cache[source]; ok {
		return val
	}

	o, err := system.Open(source)
	defer o.Close()

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
