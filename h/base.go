package h

import (
	"github.com/gofiber/fiber/v2"
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

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func GetPartialPath(partial func(ctx *fiber.Ctx) *Partial) string {
	return GetFunctionName(partial)
}

func GetPartialPathWithQs(partial func(ctx *fiber.Ctx) *Partial, qs string) string {
	return html.EscapeString(GetPartialPath(partial) + "?" + qs)
}
