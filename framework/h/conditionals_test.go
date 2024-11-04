package h

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIf(t *testing.T) {
	t.Parallel()
	result := If(true, Pf("hello"))
	assert.Equal(t, "<p>hello</p>", Render(result))

	result2 := If(false, Pf("hello"))
	assert.Equal(t, "", Render(result2)) // Expect an empty element
}

func TestIfElse(t *testing.T) {
	t.Parallel()
	result := IfElse(true, Pf("hello"), Pf("world"))
	assert.Equal(t, "<p>hello</p>", Render(result))

	result2 := IfElse(false, Pf("hello"), Pf("world"))
	assert.Equal(t, "<p>world</p>", Render(result2))
}

func TestTernary(t *testing.T) {
	t.Parallel()
	result := Ternary(true, Pf("hello"), Pf("world"))
	assert.Equal(t, "<p>hello</p>", Render(result))

	result2 := Ternary(false, Pf("hello"), Pf("world"))
	assert.Equal(t, "<p>world</p>", Render(result2))
}

func TestIfElseLazy(t *testing.T) {
	t.Parallel()
	result := IfElseLazy(true, func() *Element { return Pf("hello") }, func() *Element { return Pf("world") })
	assert.Equal(t, "<p>hello</p>", Render(result))

	result2 := IfElseLazy(false, func() *Element { return Pf("hello") }, func() *Element { return Pf("world") })
	assert.Equal(t, "<p>world</p>", Render(result2))
}

func TestElementIf(t *testing.T) {
	t.Parallel()
	element := Pf("hello")
	result := ElementIf(true, element)
	assert.Equal(t, "<p>hello</p>", Render(result))

	result2 := ElementIf(false, element)
	assert.Equal(t, "", Render(result2)) // Expect an empty element
}

func TestIfElseE(t *testing.T) {
	t.Parallel()
	element1 := Pf("hello")
	element2 := Pf("world")
	result := IfElseE(true, element1, element2)
	assert.Equal(t, "<p>hello</p>", Render(result))

	result2 := IfElseE(false, element1, element2)
	assert.Equal(t, "<p>world</p>", Render(result2))
}

func TestIfHtmxRequest(t *testing.T) {
	t.Parallel()
	ctx := &RequestContext{isHxRequest: true}
	result := IfHtmxRequest(ctx, Pf("hello"))
	assert.Equal(t, "<p>hello</p>", Render(result))

	ctx2 := &RequestContext{isHxRequest: false}
	result2 := IfHtmxRequest(ctx2, Pf("hello"))
	assert.Equal(t, "", Render(result2)) // Expect an empty element
}

func TestClassIf(t *testing.T) {
	t.Parallel()
	result := ClassIf(true, "my-class")
	assert.Equal(t, ` class="my-class"`, Render(result))

	result2 := ClassIf(false, "my-class")
	assert.Equal(t, "", Render(result2)) // Expect an empty element
}

func TestAttributeIf(t *testing.T) {
	t.Parallel()
	result := AttributeIf(true, "data-test", "value")
	assert.Equal(t, ` data-test="value"`, Render(result))

	result2 := AttributeIf(false, "data-test", "value")
	assert.Equal(t, "", Render(result2)) // Expect an empty element
}
