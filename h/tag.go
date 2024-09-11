package h

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"html"
	"net/url"
	"strings"
)

type Node struct {
	id         string
	tag        string
	attributes map[string]string
	children   []Renderable
	text       string
	value      string
	changed    bool
}

func (node *Node) Render() *Node {
	return node
}

func (node *Node) AppendChild(child Renderable) Renderable {
	node.children = append(node.children, child)
	return node
}

func (node *Node) SetChanged(changed bool) Renderable {
	node.changed = changed
	return node
}

func Data(data map[string]any) Renderable {
	serialized, err := json.Marshal(data)
	if err != nil {
		return Empty()
	}
	return Attribute("x-data", string(serialized))
}

func ClassIf(condition bool, value string) Renderable {
	if condition {
		return Class(value)
	}
	return Empty()
}

func Class(value ...string) Renderable {
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

func Id(value string) Renderable {
	if strings.HasPrefix(value, "#") {
		value = value[1:]
	}
	return Attribute("id", value)
}

func Attributes(attrs map[string]string) Renderable {
	return &Node{
		tag:        "attribute",
		attributes: attrs,
	}
}

func Attribute(key string, value string) Renderable {
	return Attributes(map[string]string{key: value})
}

func Disabled() Renderable {
	return Attribute("disabled", "")
}

func Get(path string) Renderable {
	return Attribute("hx-get", path)
}

func GetPartial(partial func(ctx *fiber.Ctx) *Partial) Renderable {
	return Get(GetPartialPath(partial))
}

func GetPartialWithQs(partial func(ctx *fiber.Ctx) *Partial, qs string) Renderable {
	return Get(GetPartialPathWithQs(partial, qs))
}

func CreateTriggers(triggers ...string) []string {
	return triggers
}

type ReloadParams struct {
	Triggers []string
}

func ViewOnLoad(partial func(ctx *fiber.Ctx) *Partial) Renderable {
	return View(partial, ReloadParams{
		Triggers: CreateTriggers("load"),
	})
}

func View(partial func(ctx *fiber.Ctx) *Partial, params ReloadParams) Renderable {
	return Div(Attributes(map[string]string{
		"hx-get":     GetPartialPath(partial),
		"hx-trigger": strings.Join(params.Triggers, ", "),
	}))
}

func ViewWithTriggers(partial func(ctx *fiber.Ctx) *Partial, triggers ...string) Renderable {
	return Div(Attributes(map[string]string{
		"hx-get":     GetPartialPath(partial),
		"hx-trigger": strings.Join(triggers, ", "),
	}))
}

func GetWithQs(path string, qs map[string]string) Renderable {
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

func Post(url string) Renderable {
	return Attribute("hx-post", url)
}

func Trigger(trigger string) Renderable {
	return Attribute("hx-trigger", trigger)
}

func Text(text string) Renderable {
	return &Node{
		tag:  "text",
		text: text,
	}
}

func Pf(format string, args ...interface{}) Renderable {
	return P(fmt.Sprintf(format, args...))
}

func Target(target string) Renderable {
	return Attribute("hx-target", target)
}

func Name(name string) Renderable {
	return Attribute("name", name)
}

func Confirm(message string) Renderable {
	return Attribute("hx-confirm", message)
}

func Href(path string) Renderable {
	return Attribute("href", path)
}

func Type(name string) Renderable {
	return Attribute("type", name)
}

func Placeholder(placeholder string) Renderable {
	return Attribute("placeholder", placeholder)
}

func OutOfBandSwap(selector string) Renderable {
	return Attribute("hx-swap-oob",
		Ternary(selector == "", "true", selector))
}

func Click(value string) Renderable {
	return Attribute("onclick", value)
}

func Tag(tag string, children ...Renderable) Renderable {
	return &Node{
		tag:      tag,
		children: children,
	}
}

func Html(children ...Renderable) Renderable {
	return Tag("html", children...)
}

func Head(children ...Renderable) Renderable {
	return Tag("head", children...)
}

func Body(children ...Renderable) Renderable {
	return Tag("body", children...)
}

func Script(url string) Renderable {
	return &Node{
		tag: "script",
		attributes: map[string]string{
			"src": url,
		},
		children: make([]Renderable, 0),
	}
}

func Raw(text string) Renderable {
	return &Node{
		tag:      "raw",
		children: make([]Renderable, 0),
		value:    text,
	}
}

func RawScript(text string) Renderable {
	return Raw("<script>" + text + "</script>")
}

func Div(children ...Renderable) Renderable {
	return Tag("div", children...)
}

func Input(inputType string, children ...Renderable) Renderable {
	return &Node{
		tag: "input",
		attributes: map[string]string{
			"type": inputType,
		},
		children: children,
	}
}

func List[T any](items []T, mapper func(item T, index int) Renderable) Renderable {
	node := &Node{
		tag:      "",
		children: make([]Renderable, len(items)),
	}
	for index, value := range items {
		node.children[index] = mapper(value, index)
	}
	return node
}

func Fragment(children ...Renderable) Renderable {
	return &Node{
		tag:      "",
		children: children,
	}
}

func AttributeList(children ...Renderable) Renderable {
	return &Node{
		tag:      FlagAttributeList,
		children: children,
	}
}

func AppendChildren(node *Node, children ...Renderable) Renderable {
	node.children = append(node.children, children...)
	return node

}

func Button(children ...Renderable) Renderable {
	return Tag("button", children...)
}

func Indicator(tag string) Renderable {
	return Attribute("hx-indicator", tag)
}

func P(text string, children ...Renderable) Renderable {
	return &Node{
		tag:      "p",
		children: children,
		text:     text,
	}
}

func Form(children ...Renderable) Renderable {
	return Tag("form", children...)
}

func A(text string, children ...Renderable) Renderable {
	return &Node{
		tag:      "a",
		children: children,
		text:     text,
	}
}

func Nav(children ...Renderable) Renderable {
	return Tag("nav", children...)
}

func Empty() Renderable {
	return &Node{
		tag: "",
	}
}

func BeforeRequestSetHtml(children ...Renderable) Renderable {
	serialized := Render(Fragment(children...))
	return Attribute("hx-on::before-request", `this.innerHTML = '`+html.EscapeString(serialized)+`'`)
}

func BeforeRequestSetAttribute(key string, value string) Renderable {
	return Attribute("hx-on::before-request", `this.setAttribute('`+key+`', '`+value+`')`)
}

func BeforeRequestSetText(text string) Renderable {
	return Attribute("hx-on::before-request", `this.innerText = '`+text+`'`)
}

func AfterRequestRemoveAttribute(key string, value string) Renderable {
	return Attribute("hx-on::after-request", `this.removeAttribute('`+key+`')`)
}

func IfQueryParam(key string, node *Node) Renderable {
	return Fragment(Attribute("hx-if-qp:"+key, "true"), node)
}

func Hidden() Renderable {
	return Attribute("style", "display:none")
}

func MatchQueryParam(defaultValue string, active string, m map[string]*Node) Renderable {

	rendered := make(map[string]string)
	for s, node := range m {
		rendered[s] = Render(node)
	}

	root := Tag("span",
		m[active],
		Trigger("url"),
		Attribute("hx-match-qp", "true"),
		Attribute("hx-match-qp-default", defaultValue),
	)

	for s, node := range rendered {
		root = AppendChildren(root.Render(), Attribute("hx-match-qp-mapping:"+s, ``+html.EscapeString(node)+``))
	}

	return root
}

func AfterRequestSetHtml(children ...Renderable) Renderable {
	serialized := Render(Fragment(children...))
	return Attribute("hx-on::after-request", `this.innerHTML = '`+html.EscapeString(serialized)+`'`)
}

func Children(children []Renderable) Renderable {
	return &Node{
		tag:      FlagChildrenList,
		children: children,
	}
}

func Label(text string) Renderable {
	return Tag("label", Text(text))
}

func If(condition bool, node Renderable) Renderable {
	if condition {
		return node
	} else {
		return Empty()
	}
}

func IfElse(condition bool, node Renderable, node2 Renderable) Renderable {
	if condition {
		return node
	} else {
		return node2
	}
}

func IfElseLazy(condition bool, cb1 func() Renderable, cb2 func() Renderable) Renderable {
	if condition {
		return cb1()
	} else {
		return cb2()
	}
}

func IfHtmxRequest(ctx *fiber.Ctx, node Renderable) Renderable {
	if ctx.Get("HX-Request") != "" {
		return node
	}
	return Empty()
}

type SwapArg struct {
	Selector string
	Content  *Node
}

func NewSwap(selector string, content *Node) SwapArg {
	return SwapArg{
		Selector: selector,
		Content:  content,
	}
}

func Swap(ctx *fiber.Ctx, content Renderable) Renderable {
	return SwapWithSelector(ctx, "", content)
}

func SwapWithSelector(ctx *fiber.Ctx, selector string, content Renderable) Renderable {
	if ctx == nil || ctx.Get("HX-Request") == "" {
		return Empty()
	}
	c := content.Render()
	return c.AppendChild(OutOfBandSwap(selector))
}

func SwapMany(ctx *fiber.Ctx, args ...SwapArg) Renderable {
	if ctx.Get("HX-Request") == "" {
		return Empty()
	}
	for _, arg := range args {
		arg.Content.AppendChild(OutOfBandSwap(arg.Selector))
	}
	return Fragment(Map(args, func(arg SwapArg) Renderable {
		return arg.Content
	})...)
}

type OnRequestSwapArgs struct {
	Target        string
	Get           string
	Default       *Node
	BeforeRequest *Node
	AfterRequest  *Node
}

func OnRequestSwap(args OnRequestSwapArgs) Renderable {
	return Div(args.Default,
		BeforeRequestSetHtml(args.BeforeRequest),
		AfterRequestSetHtml(args.AfterRequest),
		Get(args.Get),
		Target(args.Target),
	)
}
