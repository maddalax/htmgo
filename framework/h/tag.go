package h

import (
	"fmt"
	"strconv"
)

type ClassMap map[string]bool

type PartialFunc = func(ctx *RequestContext) *Partial

type Element struct {
	tag        string
	attributes map[string]string
	children   []Ren
}

func (node *Element) AppendChild(child Ren) *Element {
	node.children = append(node.children, child)
	return node
}

func (node *Element) AppendChildren(children ...Ren) *Element {
	node.children = append(node.children, children...)
	return node
}

func Svg(children ...Ren) *Element {
	return Tag("svg", children...)
}

func Path(children ...Ren) *Element {
	return Tag("path", children...)
}

func TextF(format string, args ...interface{}) *TextContent {
	return Text(fmt.Sprintf(format, args...))
}

func Text(text string) *TextContent {
	return NewTextContent(text)
}

func Pf(format string, args ...interface{}) *Element {
	return TagF("p", format, args...)
}

func Tag(tag string, children ...Ren) *Element {
	return &Element{
		tag:        tag,
		children:   children,
		attributes: make(map[string]string),
	}
}

func Html(children ...Ren) *Element {
	return Tag("html", children...)
}

func Head(children ...Ren) *Element {
	return Tag("head", children...)
}

func Body(children ...Ren) *Element {
	return Tag("body", children...)
}

func Meta(name string, content string) *Element {
	return &Element{
		tag: "meta",
		attributes: map[string]string{
			"name":    name,
			"content": content,
		},
		children: make([]Ren, 0),
	}
}

func LinkWithVersion(href string, rel string, version string) *Element {
	attributeMap := AttributeMap{
		"href": href + "?v=" + version,
		"rel":  rel,
	}
	return &Element{
		tag:        "link",
		attributes: attributeMap.ToMap(),
		children:   make([]Ren, 0),
	}
}

func Link(href string, rel string) *Element {
	attributeMap := AttributeMap{
		"href": href,
		"rel":  rel,
	}
	return &Element{
		tag:        "link",
		attributes: attributeMap.ToMap(),
		children:   make([]Ren, 0),
	}
}

func ScriptWithVersion(url string, version string) *Element {
	attributeMap := AttributeMap{
		"src": url + "?v=" + version,
	}
	return &Element{
		tag:        "script",
		attributes: attributeMap.ToMap(),
		children:   make([]Ren, 0),
	}
}

func Script(url string) *Element {
	attributeMap := AttributeMap{
		"src": url,
	}
	return &Element{
		tag:        "script",
		attributes: attributeMap.ToMap(),
		children:   make([]Ren, 0),
	}
}

func Raw(text string) *RawContent {
	return NewRawContent(text)
}

func Style(text string) *Element {
	return Tag("style", Text(text))
}

func MultiLineQuotes(text string) string {
	return "`" + text + "`"
}

func RawF(text string, args any) *RawContent {
	return Raw(fmt.Sprintf(text, args))
}

func RawScript(text string) *RawContent {
	return Raw("<script>" + text + "</script>")
}

func Pre(children ...Ren) *Element {
	return Tag("pre", children...)
}

func Div(children ...Ren) *Element {
	return Tag("div", children...)
}

func Article(children ...Ren) *Element {
	return Tag("article", children...)
}

func Checkbox(children ...Ren) *Element {
	return Input("checkbox", children...)
}

func Value(value any) *AttributeR {
	switch v := value.(type) {
	case string:
		return Attribute("value", v)
	case int:
		return Attribute("value", strconv.Itoa(v))
	default:
		return Attribute("value", fmt.Sprintf("%v", v))
	}
}

func Input(inputType string, children ...Ren) *Element {
	attributeMap := AttributeMap{
		"type": inputType,
	}
	return &Element{
		tag:        "input",
		attributes: attributeMap.ToMap(),
		children:   children,
	}
}

func TextInput(children ...Ren) *Element {
	return Input("text", children...)
}

func NumberInput(children ...Ren) *Element {
	return Input("number", children...)
}

func FileInput(children ...Ren) *Element {
	return Input("file", children...)
}

func Radio(children ...Ren) *Element {
	return Input("radio", children...)
}

func IterMap[T any](m map[string]T, mapper func(key string, value T) *Element) *Element {
	node := &Element{
		tag:      "",
		children: make([]Ren, len(m)),
	}
	index := 0
	for key, value := range m {
		node.children[index] = mapper(key, value)
		index++
	}
	return node
}

func List[T any](items []T, mapper func(item T, index int) *Element) *Element {
	node := &Element{
		tag:      "",
		children: make([]Ren, len(items)),
	}
	for index, value := range items {
		node.children[index] = mapper(value, index)
	}
	return node
}

func Fragment(children ...Ren) *Element {
	e := Empty()
	e.children = children
	return e
}

func Template(children ...Ren) *Element {
	return Tag("template", children...)
}

func AppendChildren(node *Element, children ...Ren) *Element {
	node.children = append(node.children, children...)
	return node

}

func Button(children ...Ren) *Element {
	return Tag("button", children...)
}

func P(children ...Ren) *Element {
	return Tag("p", children...)
}

func TagF(tag string, format string, args ...interface{}) *Element {

	/*
		An invocation can look like
		h.H3F("build simple and scalable systems with %s", "go + htmx", h.Class("-mt-4")),

		where the args may be a mix of strings, *Element, *AttributeMap, *ChildList, *AttributeR
		We need to separate the children from the format arguments
	*/
	children := make([]Ren, 0)
	fmtArgs := make([]interface{}, 0)
	for _, arg := range args {
		switch d := arg.(type) {
		case *Element:
			children = append(children, d)
		case *AttributeMap:
			children = append(children, d)
		case *ChildList:
			for _, child := range d.Children {
				children = append(children, child)
			}
		case *AttributeR:
			children = append(children, d)
		default:
			fmtArgs = append(fmtArgs, d)
		}
	}

	combined := Children(Text(fmt.Sprintf(format, fmtArgs...)))
	combined.Children = append(combined.Children, children...)

	return Tag(tag, combined)
}

func H1(children ...Ren) *Element {
	return Tag("h1", children...)
}

func H1F(format string, args ...interface{}) *Element {
	return TagF("h1", format, args...)
}

func H2(children ...Ren) *Element {
	return Tag("h2", children...)
}

func H2F(format string, args ...interface{}) *Element {
	return TagF("h2", format, args...)
}

func H3(children ...Ren) *Element {
	return Tag("h3", children...)
}

func H3F(format string, args ...interface{}) *Element {
	return TagF("h3", format, args...)
}

func H4(children ...Ren) *Element {
	return Tag("h4", children...)
}

func H4F(format string, args ...interface{}) *Element {
	return TagF("h4", format, args...)
}

func H5(children ...Ren) *Element {
	return Tag("h5", children...)
}

func H5F(format string, args ...interface{}) *Element {
	return TagF("h5", format, args...)
}

func H6(children ...Ren) *Element {
	return Tag("h6", children...)
}

func H6F(format string, args ...interface{}) *Element {
	return TagF("h6", format, args...)
}

func Img(children ...Ren) *Element {
	return Tag("img", children...)
}

func Src(src string) *AttributeR {
	return Attribute("src", src)
}

func Form(children ...Ren) *Element {
	return Tag("form", children...)
}

func A(children ...Ren) *Element {
	return Tag("a", children...)
}

func Nav(children ...Ren) *Element {
	return Tag("nav", children...)
}

func Empty() *Element {
	return &Element{
		tag: "",
	}
}

func Children(children ...Ren) *ChildList {
	return NewChildList(children...)
}

func Label(children ...Ren) *Element {
	return Tag("label", children...)
}

func IFrame(src string) *Element {
	return Tag("iframe", Src(src))
}

func Address(children ...Ren) *Element {
	return Tag("address", children...)
}

func Span(children ...Ren) *Element {
	return Tag("span", children...)
}

func Aside(children ...Ren) *Element {
	return Tag("aside", children...)
}

func Section(children ...Ren) *Element {
	return Tag("section", children...)
}

func Code(children ...Ren) *Element {
	return Tag("code", children...)
}

func Dialog(children ...Ren) *Element {
	return Tag("dialog", children...)
}

func FieldSet(children ...Ren) *Element {
	return Tag("fieldset", children...)
}

func Footer(children ...Ren) *Element {
	return Tag("footer", children...)
}

func Header(children ...Ren) *Element {
	return Tag("header", children...)
}

func Hr(children ...Ren) *Element {
	return Tag("hr", children...)
}

func LabelFor(id string, text string) *Element {
	return Tag("label", Attribute("for", id), Text(text))
}

func Main(children ...Ren) *Element {
	return Tag("main", children...)
}

func Ol(children ...Ren) *Element {
	return Tag("ol", children...)
}

func Ul(children ...Ren) *Element {
	return Tag("ul", children...)
}

func Li(children ...Ren) *Element {
	return Tag("li", children...)
}

func Select(children ...Ren) *Element {
	return Tag("select", children...)
}

func Option(children ...Ren) *Element {
	return Tag("option", children...)
}

func Strong(children ...Ren) *Element {
	return Tag("strong", children...)
}

func Table(children ...Ren) *Element {
	return Tag("table", children...)
}

func TBody(children ...Ren) *Element {
	return Tag("tbody", children...)
}

func Td(children ...Ren) *Element {
	return Tag("td", children...)
}

func Th(children ...Ren) *Element {
	return Tag("th", children...)
}

func Tr(children ...Ren) *Element {
	return Tag("tr", children...)
}

func THead(children ...Ren) *Element {
	return Tag("thead", children...)
}

func TFoot(children ...Ren) *Element {
	return Tag("tfoot", children...)
}

func Abbr(children ...Ren) *Element {
	return Tag("abbr", children...)
}

func Title(children ...Ren) *Element {
	return Tag("title", children...)
}

func Br() *Element {
	return Tag("br")
}
