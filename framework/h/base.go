package h

import (
	"github.com/labstack/echo/v4"
	"html"
	"net/http"
	"reflect"
	"runtime"
)

type Headers = map[string]string

type Partial struct {
	Headers *Headers
	Root    *Node
}

func (p *Partial) Render() *Node {
	return p.Root
}

type Page struct {
	Root       Renderable
	HttpMethod string
}

func NewPage(root Renderable) *Page {
	return &Page{
		HttpMethod: http.MethodGet,
		Root:       root,
	}
}

func NewPageWithHttpMethod(httpMethod string, root Renderable) *Page {
	return &Page{
		HttpMethod: httpMethod,
		Root:       root,
	}
}

func NewPartialWithHeaders(headers *Headers, root Renderable) *Partial {
	return &Partial{
		Headers: headers,
		Root:    root.Render(),
	}
}

func NewPartial(root Renderable) *Partial {
	return &Partial{
		Root: root.Render(),
	}
}

func GetPartialPath(partial func(ctx echo.Context) *Partial) string {
	return runtime.FuncForPC(reflect.ValueOf(partial).Pointer()).Name()
}

func GetPartialPathWithQs(partial func(ctx echo.Context) *Partial, qs string) string {
	return html.EscapeString(GetPartialPath(partial) + "?" + qs)
}
