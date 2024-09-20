package h

import (
	"html"
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

type Headers = map[string]string

type Partial struct {
	Headers *Headers
	Root    *Element
}

func (p *Partial) Render(builder *strings.Builder) {
	p.Root.Render(builder)
}

type Page struct {
	Root       Ren
	HttpMethod string
}

func NewPage(root Ren) *Page {
	return &Page{
		HttpMethod: http.MethodGet,
		Root:       root,
	}
}

func NewPageWithHttpMethod(httpMethod string, root *Element) *Page {
	return &Page{
		HttpMethod: httpMethod,
		Root:       root,
	}
}

func NewPartialWithHeaders(headers *Headers, root *Element) *Partial {
	return &Partial{
		Headers: headers,
		Root:    root,
	}
}

func NewPartial(root *Element) *Partial {
	return &Partial{
		Root: root,
	}
}

func GetPartialPath(partial func(ctx *RequestContext) *Partial) string {
	return runtime.FuncForPC(reflect.ValueOf(partial).Pointer()).Name()
}

func GetPartialPathWithQs(partial func(ctx *RequestContext) *Partial, qs string) string {
	return html.EscapeString(GetPartialPath(partial) + "?" + qs)
}
