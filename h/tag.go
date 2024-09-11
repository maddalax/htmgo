package h

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"html"
	"net/url"
	"strings"
)

type Node struct {
	id         string
	tag        string
	attributes map[string]string
	children   []*Node
	text       string
	value      string
	changed    bool
}

func NewNode(tag string) Node {
	return Node{
		tag:        tag,
		attributes: nil,
		children:   nil,
		text:       "",
		value:      "",
		id:         "",
	}
}

type Action struct {
	Type   string
	Target *Node
	Value  any
}

func (node *Node) AppendChild(child *Node) *Node {
	node.children = append(node.children, child)
	return node
}

func (node *Node) SetChanged(changed bool) *Node {
	node.changed = changed
	return node
}

func Data(data map[string]any) *Node {
	serialized, err := json.Marshal(data)
	if err != nil {
		return Empty()
	}
	return Attribute("x-data", string(serialized))
}

func ClassIf(condition bool, value string) *Node {
	if condition {
		return Class(value)
	}
	return Empty()
}

func Class(value ...string) *Node {
	return &Node{
		tag:   "class",
		value: MergeClasses(value...),
	}
}

func MergeClasses(classes ...string) string {
	builder := ""
	for _, s := range classes {
		builder += s + " "
	}
	return builder
}

func Id(value string) *Node {
	return Attribute("id", value)
}

func Attribute(key string, value string) *Node {
	return &Node{
		tag: "attribute",
		attributes: map[string]string{
			key: value,
		},
	}
}

func Get(path string) *Node {
	return Attribute("hx-get", path)
}

func CreateTriggers(triggers ...string) []string {
	return triggers
}

type ReloadParams struct {
	Triggers []string
}

func View(partial func(ctx *fiber.Ctx) *Partial, params ReloadParams) *Node {
	return &Node{
		tag: "attribute",
		attributes: map[string]string{
			"hx-get":     GetPartialPath(partial),
			"hx-trigger": strings.Join(params.Triggers, ", "),
		},
	}
}

func GetWithQs(path string, qs map[string]string) *Node {
	u, err := url.Parse(path)
	if err != nil {
		return Empty()
	}

	q := u.Query()

	for s := range qs {
		q.Add(s, qs[s])
	}

	u.RawQuery = q.Encode()

	return Get(u.String())
}

func Post(url string) *Node {
	return Attribute("hx-post", url)
}

func Trigger(trigger string) *Node {
	return Attribute("hx-trigger", trigger)
}

func Text(text string) *Node {
	return &Node{
		tag:  "text",
		text: text,
	}
}

func Target(target string) *Node {
	return Attribute("hx-target", target)
}

func Name(name string) *Node {
	return Attribute("name", name)
}

func Confirm(message string) *Node {
	return Attribute("hx-confirm", message)
}

func Href(path string) *Node {
	return Attribute("href", path)
}

func Type(name string) *Node {
	return Attribute("type", name)
}

func Placeholder(placeholder string) *Node {
	return Attribute("placeholder", placeholder)
}

func Swap(swap string) *Node {
	return Attribute("hx-swap", swap)
}

func Click(value string) *Node {
	return Attribute("onclick", value)
}

func OnClickWs(handler string) *Node {
	return Attribute("data-ws-click", handler)
}

func Html(children ...*Node) *Node {
	return &Node{
		tag:      "html",
		children: children,
	}
}

func Head(children ...*Node) *Node {
	return &Node{
		tag:      "head",
		children: children,
	}
}

func Body(children ...*Node) *Node {
	return &Node{
		tag:      "body",
		children: children,
	}
}

func Script(url string) *Node {
	return &Node{
		tag: "script",
		attributes: map[string]string{
			"src": url,
		},
		children: make([]*Node, 0),
	}
}

func Raw(text string) *Node {
	return &Node{
		tag:      "raw",
		children: make([]*Node, 0),
		value:    text,
	}
}

func RawScript(text string) *Node {
	return Raw("<script>" + text + "</script>")
}

func Div(children ...*Node) *Node {
	return &Node{
		tag:      "div",
		children: children,
	}
}

func Input(inputType string, children ...*Node) *Node {
	return &Node{
		tag: "input",
		attributes: map[string]string{
			"type": inputType,
		},
		children: children,
	}
}

func HStack(children ...*Node) *Node {
	return &Node{
		tag: "div",
		attributes: map[string]string{
			"class": "flex gap-2",
		},
		children: children,
	}
}

func VStack(children ...*Node) *Node {
	return &Node{
		tag: "div",
		attributes: map[string]string{
			"class": "flex flex-col gap-2",
		},
		children: children,
	}
}

func List[T any](items []T, mapper func(item T) *Node) *Node {
	node := &Node{
		tag:      "",
		children: make([]*Node, len(items)),
	}
	for index, value := range items {
		node.children[index] = mapper(value)
	}
	return node
}

func Fragment(children ...*Node) *Node {
	return &Node{
		tag:      "",
		children: children,
	}
}

func AttributeList(children ...*Node) *Node {
	return &Node{
		tag:      FlagAttributeList,
		children: children,
	}
}

func AppendChildren(node *Node, children ...*Node) *Node {
	node.children = append(node.children, children...)
	return node

}

func Button(children ...*Node) *Node {
	return &Node{
		tag:      "button",
		children: children,
	}
}

func Indicator(tag string) *Node {
	return Attribute("hx-indicator", tag)
}

func P(text string, children ...*Node) *Node {
	return &Node{
		tag:      "p",
		children: children,
		text:     text,
	}
}

func A(text string, children ...*Node) *Node {
	return &Node{
		tag:      "a",
		children: children,
		text:     text,
	}
}

func Empty() *Node {
	return &Node{
		tag: "",
	}
}

func BeforeRequestSetHtml(children ...*Node) *Node {
	serialized := Render(Fragment(children...))
	return Attribute("hx-on::before-request", `this.innerHTML = '`+html.EscapeString(serialized)+`'`)
}

func AfterRequestSetHtml(children ...*Node) *Node {
	serialized := Render(Fragment(children...))
	return Attribute("hx-on::after-request", `this.innerHTML = '`+html.EscapeString(serialized)+`'`)
}

func If(condition bool, node *Node) *Node {
	if condition {
		return node
	} else {
		return Empty()
	}
}

func JsIf(condition string, node *Node) *Node {
	node1 := &Node{tag: "template"}
	node1.AppendChild(Attribute("x-if", condition))
	node1.AppendChild(node)
	return node
}

func JsIfElse(condition string, node *Node, node2 *Node) *Node {
	node1Template := Div(Attribute("x-show", condition), node)
	node2Template := Div(Attribute("x-show", "!("+condition+")"), node2)
	return Fragment(node1Template, node2Template)
}

func IfElse(condition bool, node *Node, node2 *Node) *Node {
	if condition {
		return node
	} else {
		return node2
	}
}
