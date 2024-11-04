package h

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagFunctions(t *testing.T) {
	tests := []struct {
		name     string
		element  *Element
		expected string
	}{
		{"Div", Div(), "div"},
		{"Span", Span(), "span"},
		{"Button", Button(), "button"},
		{"P", P(), "p"},
		{"H1", H1(), "h1"},
		{"H2", H2(), "h2"},
		{"H3", H3(), "h3"},
		{"H4", H4(), "h4"},
		{"H5", H5(), "h5"},
		{"H6", H6(), "h6"},
		{"Img", Img(), "img"},
		{"Video", Video(), "video"},
		{"Form", Form(), "form"},
		{"A", A(), "a"},
		{"Nav", Nav(), "nav"},
		{"Section", Section(), "section"},
		{"Aside", Aside(), "aside"},
		{"Header", Header(), "header"},
		{"Footer", Footer(), "footer"},
		{"Main", Main(), "main"},
		{"Ul", Ul(), "ul"},
		{"Li", Li(), "li"},
		{"Br", Br(), "br"},
		{"Hr", Hr(), "hr"},
		{"Ol", Ol(), "ol"},
		{"Table", Table(), "table"},
		{"Tr", Tr(), "tr"},
		{"Td", Td(), "td"},
		{"Th", Th(), "th"},
		{"THead", THead(), "thead"},
		{"TBody", TBody(), "tbody"},
		{"TFoot", TFoot(), "tfoot"},
		{"Abbr", Abbr(), "abbr"},
		{"Strong", Strong(), "strong"},
		{"Code", Code(), "code"},
		{"Title", Title(), "title"},
		{"Dialog", Dialog(), "dialog"},
		{"FieldSet", FieldSet(), "fieldset"},
		{"Option", Option(), "option"},
		{"Select", Select(), "select"},
		{"Template", Template(), "template"},
		{"Label", Label(), "label"},
		{"Address", Address(), "address"},
		{"Pre", Pre(), "pre"},
		{"Article", Article(), "article"},
		{"Summary", Summary(), "summary"},
		{"Details", Details(), "details"},
		{"Svg", Svg(), "svg"},
		{"Path", Path(), "path"},
		{"Html", Html(), "html"},
		{"Head", Head(), "head"},
		{"Body", Body(), "body"},
		{"Meta", Meta("description", "test"), "meta"},
		{"Link", Link("style.css", "stylesheet"), "link"},
		{"LinkWithVersion", LinkWithVersion("style.css", "stylesheet", "1.0"), "link"},
		{"Script", Script("script.js"), "script"},
		{"ScriptWithVersion", ScriptWithVersion("script.js", "1.0"), "script"},
		{"Style", Style("body {background: #000;}"), "style"},
		{"Pre", Pre(), "pre"},
		{"Article", Article(), "article"},
		{"Checkbox", Checkbox(), "input"},
		{"TextArea", TextArea(), "textarea"},
		{"TextInput", TextInput(), "input"},
		{"NumberInput", NumberInput(), "input"},
		{"FileInput", FileInput(), "input"},
		{"Radio", Radio(), "input"},
		{"Fragment", Fragment(), ""},
		{"Template", Template(), "template"},
		{"IFrame", IFrame("example.com"), "iframe"},
		{"LabelFor", LabelFor("input-id", "Input Label"), "label"},
		{"I", I(), "i"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.element.tag)
		})
	}
}

func TestList(t *testing.T) {
	t.Parallel()
	items := []string{"hello", "world"}
	list := Ul(List(items, func(item string, index int) *Element {
		return Li(Text(item))
	}))
	assert.Equal(t, `<ul><li>hello</li><li>world</li></ul>`, Render(list))
}

func TestIterMap(t *testing.T) {
	t.Parallel()
	items := map[string]string{
		"hello": "world",
	}
	list := Ul(IterMap(items, func(key string, value string) *Element {
		return Li(Text(key), Text(value))
	}))
	assert.Equal(t, `<ul><li>helloworld</li></ul>`, Render(list))
}

func TestUnsafeRaw(t *testing.T) {
	t.Parallel()
	element := UnsafeRaw("<div>Hello World</div>")
	assert.Equal(t, "<div>Hello World</div>", Render(element))
}

func TestUnsafeRawScript(t *testing.T) {
	t.Parallel()
	element := UnsafeRawScript("alert('Hello World')")
	assert.Equal(t, "<script>alert('Hello World')</script>", Render(element))
}

func TestUnsafeRawF(t *testing.T) {
	t.Parallel()
	element := UnsafeRawF("Hello %s", "World")
	assert.Equal(t, "Hello World", Render(element))
}

func TestMultiLineQuotes(t *testing.T) {
	t.Parallel()
	element := MultiLineQuotes("Hello World")
	assert.Equal(t, "`Hello World`", element)
}

func TestValue(t *testing.T) {
	t.Parallel()
	assert.Equal(t, ` value="Hello World"`, Render(Value("Hello World")))
	assert.Equal(t, ` value="1"`, Render(Value(1)))
	assert.Equal(t, ` value="true"`, Render(Value(true)))
}

func TestAppendChildren(t *testing.T) {
	t.Parallel()
	element := Div()
	element.AppendChildren(Div(), Div())
	assert.Equal(t, "<div><div></div><div></div></div>", Render(element))
}

func TestTagF(t *testing.T) {
	t.Parallel()
	element := TagF("div", "Hello %s", "World")
	assert.Equal(t, "<div>Hello World</div>", Render(element))

	element2 := TagF("div", "Hello World", Class("my-class"))
	assert.Equal(t, "<div class=\"my-class\">Hello World</div>", Render(element2))

	element3 := TagF("div", "Value", P(Text("Hello World")))
	assert.Equal(t, "<div>Value<p>Hello World</p></div>", Render(element3))
}

func TestTag(t *testing.T) {
	t.Parallel()
	element := Tag("div")
	assert.Equal(t, "<div></div>", Render(element))
}
