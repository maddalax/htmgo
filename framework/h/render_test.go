package h

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"sort"
	"strings"
	"testing"
)

// Sort attributes of a node by attribute name
func sortAttributes(node *html.Node) {
	if node.Type == html.ElementNode && len(node.Attr) > 1 {
		sort.SliceStable(node.Attr, func(i, j int) bool {
			return node.Attr[i].Key < node.Attr[j].Key
		})
	}
}

// Traverse and sort attributes in the entire HTML tree
func traverseAndSortAttributes(node *html.Node) {
	sortAttributes(node)
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		traverseAndSortAttributes(child)
	}
}

// Parse HTML, sort attributes, and render back to a string
func parseSortAndRenderHTML(input string) string {
	// Parse the HTML string into a node tree
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return ""
	}

	// Traverse and sort attributes for each node
	traverseAndSortAttributes(doc)

	// Use a buffer to capture the rendered HTML
	var buf bytes.Buffer
	err = html.Render(&buf, doc)
	if err != nil {
		return ""
	}

	// Return the rendered HTML as a string
	return buf.String()
}

func TestRender(t *testing.T) {
	t.Parallel()
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

	expectedRaw := `<div data-attr-1="value" id="my-div" data-attr-2="value" data-attr-3="value" data-attr-4="value" hx-on::before-request="this.innerText = 'before request';" hx-on::after-request="this.innerText = 'after request';"><div >hello, world</div>hello, child</div>`
	expected := parseSortAndRenderHTML(expectedRaw)
	result := parseSortAndRenderHTML(Render(div))

	assert.Equal(t,
		expected,
		result)
}

func TestRawContent(t *testing.T) {
	t.Parallel()
	str := "<div>hello, world</div>"
	raw := Raw(str)
	assert.Equal(t, str, Render(raw))
}

func TestConditional(t *testing.T) {
	t.Parallel()
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
