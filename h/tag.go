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
	children   []*Node
	text       string
	value      string
	changed    bool
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
	if strings.HasPrefix(value, "#") {
		value = value[1:]
	}
	return Attribute("id", value)
}

func Attributes(attrs map[string]string) *Node {
	return &Node{
		tag:        "attribute",
		attributes: attrs,
	}
}

func Attribute(key string, value string) *Node {
	return Attributes(map[string]string{key: value})
}

func Disabled() *Node {
	return Attribute("disabled", "")
}

func Get(path string) *Node {
	return Attribute("hx-get", path)
}

func GetPartial(partial func(ctx *fiber.Ctx) *Partial) *Node {
	return Get(GetPartialPath(partial))
}

func GetPartialWithQs(partial func(ctx *fiber.Ctx) *Partial, qs string) *Node {
	return Get(GetPartialPathWithQs(partial, qs))
}

func CreateTriggers(triggers ...string) []string {
	return triggers
}

type ReloadParams struct {
	Triggers []string
}

func ViewOnLoad(partial func(ctx *fiber.Ctx) *Partial) *Node {
	return View(partial, ReloadParams{
		Triggers: CreateTriggers("load"),
	})
}

func View(partial func(ctx *fiber.Ctx) *Partial, params ReloadParams) *Node {
	return Div(Attributes(map[string]string{
		"hx-get":     GetPartialPath(partial),
		"hx-trigger": strings.Join(params.Triggers, ", "),
	}))
}

func ViewWithTriggers(partial func(ctx *fiber.Ctx) *Partial, triggers ...string) *Node {
	return Div(Attributes(map[string]string{
		"hx-get":     GetPartialPath(partial),
		"hx-trigger": strings.Join(triggers, ", "),
	}))
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

func Pf(format string, args ...interface{}) *Node {
	return P(fmt.Sprintf(format, args...))
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

func OutOfBandSwap(selector string) *Node {
	return Attribute("hx-swap-oob",
		Ternary(selector == "", "true", selector))
}

func Click(value string) *Node {
	return Attribute("onclick", value)
}

func Tag(tag string, children ...*Node) *Node {
	return &Node{
		tag:      tag,
		children: children,
	}
}

func Html(children ...*Node) *Node {
	return Tag("html", children...)
}

func Head(children ...*Node) *Node {
	return Tag("head", children...)
}

func Body(children ...*Node) *Node {
	return Tag("body", children...)
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
	return Tag("div", children...)
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
	return Tag("button", children...)
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

func Nav(children ...*Node) *Node {
	return Tag("nav", children...)
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

func BeforeRequestSetAttribute(key string, value string) *Node {
	return Attribute("hx-on::before-request", `this.setAttribute('`+key+`', '`+value+`')`)
}

func BeforeRequestSetText(text string) *Node {
	return Attribute("hx-on::before-request", `this.innerText = '`+text+`'`)
}

func AfterRequestRemoveAttribute(key string, value string) *Node {
	return Attribute("hx-on::after-request", `this.removeAttribute('`+key+`')`)
}

func IfQueryParam(key string, node *Node) *Node {
	return Fragment(Attribute("hx-if-qp:"+key, "true"), node)
}

func Hidden() *Node {
	return Attribute("style", "display:none")
}

func MatchQueryParam(defaultValue string, active string, m map[string]*Node) *Node {

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
		root = AppendChildren(root, Attribute("hx-match-qp-mapping:"+s, ``+html.EscapeString(node)+``))
	}

	return root
}

func AfterRequestSetHtml(children ...*Node) *Node {
	serialized := Render(Fragment(children...))
	return Attribute("hx-on::after-request", `this.innerHTML = '`+html.EscapeString(serialized)+`'`)
}

func Children(children []*Node) *Node {
	return &Node{
		tag:      FlagChildrenList,
		children: children,
	}
}

func If(condition bool, node *Node) *Node {
	if condition {
		return node
	} else {
		return Empty()
	}
}

func IfElse(condition bool, node *Node, node2 *Node) *Node {
	if condition {
		return node
	} else {
		return node2
	}
}

func IfElseLazy(condition bool, cb1 func() *Node, cb2 func() *Node) *Node {
	if condition {
		return cb1()
	} else {
		return cb2()
	}
}

func IfHtmxRequest(ctx *fiber.Ctx, node *Node) *Node {
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

func Swap(ctx *fiber.Ctx, content *Node) *Node {
	return SwapWithSelector(ctx, "", content)
}

func SwapWithSelector(ctx *fiber.Ctx, selector string, content *Node) *Node {
	if ctx == nil || ctx.Get("HX-Request") == "" {
		return Empty()
	}
	return content.AppendChild(OutOfBandSwap(selector))
}

func SwapMany(ctx *fiber.Ctx, args ...SwapArg) *Node {
	if ctx.Get("HX-Request") == "" {
		return Empty()
	}
	for _, arg := range args {
		arg.Content.AppendChild(OutOfBandSwap(arg.Selector))
	}
	return Fragment(Map(args, func(arg SwapArg) *Node {
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

func OnRequestSwap(args OnRequestSwapArgs) *Node {
	return Div(args.Default,
		BeforeRequestSetHtml(args.BeforeRequest),
		AfterRequestSetHtml(args.AfterRequest),
		Get(args.Get),
		Target(args.Target),
	)
}
