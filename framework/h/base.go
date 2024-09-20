package h

import (
	"html"
	"net/http"
	"reflect"
	"runtime"
)

type Headers = map[string]string

type Partial struct {
	Headers *Headers
	Root    string
}

func (p *Partial) Render() string {
	return p.Root
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

func NewPageWithHttpMethod(httpMethod string, root Ren) *Page {
	return &Page{
		HttpMethod: httpMethod,
		Root:       root,
	}
}

func NewPartialWithHeaders(headers *Headers, root Ren) *Partial {
	return &Partial{
		Headers: headers,
		Root:    root.Render(),
	}
}

func NewPartial(root Ren) *Partial {
	return &Partial{
		Root: root.Render(),
	}
}

func GetPartialPath(partial func(ctx *RequestContext) *Partial) string {
	return runtime.FuncForPC(reflect.ValueOf(partial).Pointer()).Name()
}

func GetPartialPathWithQs(partial func(ctx *RequestContext) *Partial, qs string) string {
	return html.EscapeString(GetPartialPath(partial) + "?" + qs)
}
