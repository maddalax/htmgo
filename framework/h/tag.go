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

func TextF(format string, args ...interface{}) *TextContent {
	return Text(fmt.Sprintf(format, args...))
}

func Text(text string) *TextContent {
	return NewTextContent(text)
}

func Pf(format string, args ...interface{}) Ren {
	return P(Text(fmt.Sprintf(format, args...)))
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

func Meta(name string, content string) Ren {
	return &Element{
		tag: "meta",
		attributes: map[string]string{
			"name":    name,
			"content": content,
		},
		children: make([]Ren, 0),
	}
}

func Link(href string, rel string) Ren {
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

func Script(url string) Ren {
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

func Checkbox(children ...Ren) Ren {
	return Input("checkbox", children...)
}

func Value(value any) Ren {
	switch v := value.(type) {
	case string:
		return Attribute("value", v)
	case int:
		return Attribute("value", strconv.Itoa(v))
	default:
		return Attribute("value", fmt.Sprintf("%v", v))
	}
}

func Input(inputType string, children ...Ren) Ren {
	attributeMap := AttributeMap{
		"type": inputType,
	}
	return &Element{
		tag:        "input",
		attributes: attributeMap.ToMap(),
		children:   children,
	}
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

func Fragment(children ...Ren) *ChildList {
	return Children(children...)
}

func Template(children ...Ren) *Element {
	return Tag("template", children...)
}

func AppendChildren(node *Element, children ...Ren) Ren {
	node.children = append(node.children, children...)
	return node

}

func Button(children ...Ren) *Element {
	return Tag("button", children...)
}

func P(children ...Ren) *Element {
	return Tag("p", children...)
}

func H1(children ...Ren) *Element {
	return Tag("h1", children...)
}

func H2(children ...Ren) *Element {
	return Tag("h2", children...)
}

func H3(children ...Ren) *Element {
	return Tag("h3", children...)
}

func H4(children ...Ren) *Element {
	return Tag("h4", children...)
}

func H5(children ...Ren) *Element {
	return Tag("h5", children...)
}

func H6(children ...Ren) *Element {
	return Tag("h6", children...)
}

func Img(children ...Ren) *Element {
	return Tag("img", children...)
}

func Src(src string) Ren {
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

func Label(text string) *Element {
	return Tag("label", Text(text))
}
