package h

import (
	"github.com/maddalax/htmgo/framework/hx"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSwap(t *testing.T) {
	content := Div()
	option := SwapOption{Selector: "#myDiv", SwapType: hx.SwapTypeInnerHtml, Modifier: "test-mod"}
	swapArg := NewSwap(content, option)

	assert.Equal(t, content, swapArg.Content)
	assert.Equal(t, option, swapArg.Option)
}

func TestOobSwap(t *testing.T) {
	ctx := &RequestContext{isHxRequest: true}
	content := Div()

	result := OobSwap(ctx, content)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result.children))
	assert.Equal(t, Attribute(hx.SwapOobAttr, hx.SwapTypeTrue), result.children[0])

	// Test with a nil context
	result = OobSwap(nil, content)
	assert.Equal(t, Empty(), result)

	// Test with non-HX request
	ctx.isHxRequest = false
	result = OobSwap(ctx, content)
	assert.Equal(t, Empty(), result)
}

func TestOobSwapWithSelector(t *testing.T) {
	ctx := &RequestContext{isHxRequest: true}
	content := Div()

	result := OobSwapWithSelector(ctx, "#selector", content)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result.children))
	assert.Equal(t, Attribute(hx.SwapOobAttr, "#selector"), result.children[0])
}

func TestSwapMany(t *testing.T) {
	ctx := &RequestContext{isHxRequest: true}
	element1 := Div()
	element2 := Span()

	result := SwapMany(ctx, element1, element2)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result.children))
	assert.Equal(t, element1, result.children[0])
	assert.Equal(t, element2, result.children[1])
	assert.Equal(t, Attribute(hx.SwapOobAttr, hx.SwapTypeTrue), element1.children[0])
	assert.Equal(t, Attribute(hx.SwapOobAttr, hx.SwapTypeTrue), element2.children[0])

	// Test with non-HX request
	ctx.isHxRequest = false
	result = SwapMany(ctx, element1, element2)
	assert.Equal(t, Empty(), result)
}

func TestSwapManyX(t *testing.T) {
	ctx := &RequestContext{isHxRequest: true}
	content1 := Div()
	content2 := Span()
	option := SwapOption{SwapType: hx.SwapTypeBeforeEnd}

	arg1 := NewSwap(content1)
	arg2 := NewSwap(content2, option)

	result := SwapManyX(ctx, arg1, arg2)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result.children))
	assert.Equal(t, content1, result.children[0])
	assert.Equal(t, content2, result.children[1])
	assert.Equal(t, Attribute(hx.SwapOobAttr, hx.SwapTypeTrue), content1.children[0])
	assert.Equal(t, Attribute(hx.SwapOobAttr, hx.SwapTypeBeforeEnd), content2.children[0])

	// Test with non-HX request
	ctx.isHxRequest = false
	result = SwapManyX(ctx, arg1, arg2)
	assert.Equal(t, Empty(), result)
}
