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

func (p *Partial) ToNode() *Node {
	return p.Root
}

type Page struct {
	Root       *Node
	HttpMethod string
}

func NewPage(root *Node) *Page {
	return &Page{
		HttpMethod: http.MethodGet,
		Root:       root,
	}
}

func NewPageWithHttpMethod(httpMethod string, root *Node) *Page {
	return &Page{
		HttpMethod: httpMethod,
		Root:       root,
	}
}

func NewPartialWithHeaders(headers *Headers, root *Node) *Partial {
	return &Partial{
		Headers: headers,
		Root:    root,
	}
}

func NewPartial(root *Node) *Partial {
	return &Partial{
		Root: root,
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
