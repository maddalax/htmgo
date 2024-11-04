package markdown

import (
	"bytes"
	"github.com/alecthomas/chroma/v2"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"io"
	"io/fs"
	"sync"
)

type Renderer struct {
	cache map[string]string
	lock  sync.Mutex
}

func NewRenderer() *Renderer {
	return &Renderer{cache: make(map[string]string)}
}

func (r *Renderer) RenderFile(source string, system fs.FS) string {
	r.lock.Lock()
	defer r.lock.Unlock()
	if val, ok := r.cache[source]; ok {
		return val
	}

	o, err := system.Open(source)

	if o == nil {
		return ""
	}

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
			html.WithHardWraps(),
		),
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
					chromahtml.WithCustomCSS(map[chroma.TokenType]string{
						chroma.PreWrapper: "font-size: 14px; padding: 12px; overflow: auto; background-color: rgb(245, 245, 245) !important;",
					}),
				),
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
