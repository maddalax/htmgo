package h

import (
	"github.com/maddalax/htmgo/framework/hx"
	"html"
	"net/http"
	"reflect"
	"runtime"
)

type Partial struct {
	Headers *Headers
	Root    *Element
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

func SwapManyPartialWithHeaders(ctx *RequestContext, headers *Headers, swaps ...*Element) *Partial {
	return NewPartialWithHeaders(
		headers,
		SwapMany(ctx, swaps...),
	)
}

func SwapPartial(ctx *RequestContext, swap *Element) *Partial {
	return NewPartial(
		SwapMany(ctx, swap))
}

func RedirectPartial(url string) *Partial {
	return NewPartialWithHeaders(NewHeaders(hx.RedirectHeader, url), Fragment())
}

func SwapManyPartial(ctx *RequestContext, swaps ...*Element) *Partial {
	return NewPartial(
		SwapMany(ctx, swaps...),
	)
}

func SwapManyXPartial(ctx *RequestContext, swaps ...SwapArg) *Partial {
	return NewPartial(
		SwapManyX(ctx, swaps...),
	)
}

func GetPartialPath(partial PartialFunc) string {
	return runtime.FuncForPC(reflect.ValueOf(partial).Pointer()).Name()
}

func GetPartialPathWithQs(partial func(ctx *RequestContext) *Partial, qs *Qs) string {
	return html.EscapeString(GetPartialPath(partial) + "?" + qs.ToString())
}
