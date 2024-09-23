package h

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRender(t *testing.T) {
	div := Div(
		Id("my-div"),
		Attribute("data-attr-2", "value"),
		Attributes(&AttributeMap{
			"data-attr-3": "value",
			"data-attr-4": "value",
		}),
		HxBeforeRequest(
			SetText("before request"),
		),
		HxAfterRequest(
			SetText("after request"),
		),
		Children(
			Div(Text("hello, world")),
		),
		Text("hello, child"),
	)

	div.attributes["data-attr-1"] = "value"

	result := Render(div)

	assert.Equal(t,
		`<div data-attr-1="value" id="my-div" data-attr-2="value" data-attr-3="value" data-attr-4="value" hx-on::before-request="this.innerText = 'before request';" hx-on::after-request="this.innerText = 'after request';"><div >hello, world</div>hello, child</div>`,
		result)
}

func TestRawContent(t *testing.T) {
	str := "<div>hello, world</div>"
	raw := Raw(str)
	assert.Equal(t, str, Render(raw))
}

func TestConditional(t *testing.T) {
	result := Render(
		Div(
			Ternary(true, Text("true"), Text("false")),
		),
	)
	assert.Equal(t, "<div >true</div>", result)

	result = Render(
		Div(
			If(false, Text("true")),
		),
	)
	assert.Equal(t, "<div ></div>", result)
}
