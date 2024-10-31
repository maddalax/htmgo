package h

import (
	"github.com/maddalax/htmgo/framework/hx"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPage(t *testing.T) {
	root := Div()
	page := NewPage(root)

	assert.Equal(t, http.MethodGet, page.HttpMethod)
	assert.Equal(t, root, page.Root)
}

func TestEmptyPage(t *testing.T) {
	page := EmptyPage()

	assert.Equal(t, http.MethodGet, page.HttpMethod)
	assert.Equal(t, Empty(), page.Root)
}

func TestNewPageWithHttpMethod(t *testing.T) {
	root := Div()
	page := NewPageWithHttpMethod(http.MethodPost, root)

	assert.Equal(t, http.MethodPost, page.HttpMethod)
	assert.Equal(t, root, page.Root)
}

func TestNewPartial(t *testing.T) {
	root := Div()
	partial := NewPartial(root)

	assert.Nil(t, partial.Headers)
	assert.Equal(t, root, partial.Root)
}

func TestNewPartialWithHeaders(t *testing.T) {
	root := Div()
	headers := NewHeaders("Content-Type", "application/json")
	partial := NewPartialWithHeaders(headers, root)

	assert.Equal(t, headers, partial.Headers)
	assert.Equal(t, root, partial.Root)
}

func TestSwapManyPartialWithHeaders(t *testing.T) {
	ctx := &RequestContext{isHxRequest: true}
	headers := NewHeaders("HX-Trigger", "reload")
	elements := []*Element{Div(), Span()}

	partial := SwapManyPartialWithHeaders(ctx, headers, elements...)

	assert.Equal(t, headers, partial.Headers)
	assert.Equal(t, SwapMany(ctx, elements...), partial.Root)
}

func TestRedirectPartial(t *testing.T) {
	partial := RedirectPartial("/new-path")
	headers := NewHeaders("HX-Redirect", "/new-path")

	assert.Equal(t, headers, partial.Headers)
	assert.Equal(t, Empty(), partial.Root)
}

func TestRedirectPartialWithHeaders(t *testing.T) {
	extraHeaders := NewHeaders("X-Custom", "value")
	partial := RedirectPartialWithHeaders("/redirect-path", extraHeaders)

	expectedHeaders := NewHeaders("HX-Redirect", "/redirect-path", "X-Custom", "value")
	assert.Equal(t, expectedHeaders, partial.Headers)
	assert.Equal(t, Empty(), partial.Root)
}

func TestIsEmptyPartial(t *testing.T) {
	emptyPartial := EmptyPartial()
	nonEmptyPartial := NewPartial(Div())

	assert.True(t, IsEmptyPartial(emptyPartial))
	assert.False(t, IsEmptyPartial(nonEmptyPartial))
}

func TestGetPartialPath(t *testing.T) {
	partial := func(ctx *RequestContext) *Partial {
		return &Partial{}
	}
	path := GetPartialPath(partial)

	expectedSegment := "github.com/maddalax/htmgo/framework/h.TestGetPartialPath.func1"
	assert.Contains(t, path, expectedSegment)
}

func TestGetPartialPathWithQs(t *testing.T) {
	partial := func(ctx *RequestContext) *Partial {
		return &Partial{}
	}
	qs := NewQs("param1", "value1", "param2", "value2")
	pathWithQs := GetPartialPathWithQs(partial, qs)

	assert.Contains(t, pathWithQs, "param1=value1&param2=value2")
}

func TestSwapManyPartial(t *testing.T) {
	ctx := &RequestContext{isHxRequest: true}
	element1 := Div()
	element2 := Span()

	partial := SwapManyPartial(ctx, element1, element2)

	// Ensuring the elements have been marked for swap
	assert.Equal(t, 1, len(element1.children))
	assert.Equal(t, 1, len(element2.children))

	assert.Equal(t, Attribute(hx.SwapOobAttr, hx.SwapTypeTrue), element1.children[0])
	assert.Equal(t, Attribute(hx.SwapOobAttr, hx.SwapTypeTrue), element2.children[0])

	// Test with non-HX request context
	ctx.isHxRequest = false
	partial = SwapManyPartial(ctx, element1, element2)
	assert.True(t, IsEmptyPartial(partial))
}

func TestSwapPartial(t *testing.T) {
	ctx := &RequestContext{isHxRequest: true}
	element := Div()

	partial := SwapPartial(ctx, element)

	// Ensuring the element has been marked for swap
	assert.Equal(t, 1, len(element.children))
	assert.Equal(t, Attribute(hx.SwapOobAttr, hx.SwapTypeTrue), element.children[0])

	// Test with non-HX request context
	ctx.isHxRequest = false
	partial = SwapPartial(ctx, element)
	assert.True(t, IsEmptyPartial(partial))
}
