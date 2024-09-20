package h

import (
	"encoding/json"
	"fmt"
	"html"
	"net/url"
	"strings"
)

type Element struct {
	tag        string
	attributes map[string]string
	children   []Ren
}

func (node *Element) AppendChild(child Ren) *Element {
	node.children = append(node.children, child)
	return node
}

func Data(data map[string]any) Ren {
	serialized, err := json.Marshal(data)
	if err != nil {
		return Empty()
	}
	return Attribute("x-data", string(serialized))
}

func ClassIf(condition bool, value string) Ren {
	if condition {
		return Class(value)
	}
	return Empty()
}

func Class(value ...string) Ren {
	return Attribute("class", MergeClasses(value...))
}

func ClassX(value string, m ClassMap) Ren {
	builder := strings.Builder{}
	builder.WriteString(value)
	builder.WriteString(" ")
	for k, v := range m {
		if v {
			builder.WriteString(k)
			builder.WriteString(" ")
		}
	}
	return Class(builder.String())
}

func MergeClasses(classes ...string) string {
	if len(classes) == 1 {
		return classes[0]
	}
	builder := strings.Builder{}
	for _, s := range classes {
		builder.WriteString(s)
		builder.WriteString(" ")
	}
	return builder.String()
}

func Id(value string) Ren {
	if strings.HasPrefix(value, "#") {
		value = value[1:]
	}
	return Attribute("id", value)
}

type ClassMap map[string]bool

func Attributes(attrs *AttributeMap) *AttributeMap {
	return attrs
}

func Checked() Ren {
	return Attribute("checked", "true")
}

func Boost() Ren {
	return Attribute("hx-boost", "true")
}

func Attribute(key string, value string) *AttributeMap {
	return Attributes(&AttributeMap{key: value})
}

func TriggerChildren() Ren {
	return HxExtension("trigger-children")
}

func HxExtension(value string) Ren {
	return Attribute("hx-ext", value)
}

func Disabled() Ren {
	return Attribute("disabled", "")
}

func Get(path string) Ren {
	return Attribute("hx-get", path)
}

func GetPartial(partial func(ctx *RequestContext) *Partial) Ren {
	return Get(GetPartialPath(partial))
}

func GetPartialWithQs(partial func(ctx *RequestContext) *Partial, qs string) Ren {
	return Get(GetPartialPathWithQs(partial, qs))
}

func CreateTriggers(triggers ...string) []string {
	return triggers
}

type ReloadParams struct {
	Triggers []string
	Target   string
	Children Ren
}

func ViewOnLoad(partial func(ctx *RequestContext) *Partial) Ren {
	return View(partial, ReloadParams{
		Triggers: CreateTriggers("load"),
	})
}

func View(partial func(ctx *RequestContext) *Partial, params ReloadParams) Ren {
	return Div(Attributes(&AttributeMap{
		"hx-get":     GetPartialPath(partial),
		"hx-trigger": strings.Join(params.Triggers, ", "),
		"hx-target":  params.Target,
	}), params.Children)
}

func PartialWithTriggers(partial func(ctx *RequestContext) *Partial, triggers ...string) Ren {
	return Div(Attributes(&AttributeMap{
		"hx-get":     GetPartialPath(partial),
		"hx-trigger": strings.Join(triggers, ", "),
	}))
}

func GetWithQs(path string, qs map[string]string) Ren {
	return Get(SetQueryParams(path, qs))
}

func PostPartialOnTrigger(partial func(ctx *RequestContext) *Partial, triggers ...string) Ren {
	return PostOnTrigger(GetPartialPath(partial), strings.Join(triggers, ", "))
}

func PostPartialWithQsOnTrigger(partial func(ctx *RequestContext) *Partial, qs string, trigger string) Ren {
	return PostOnTrigger(GetPartialPathWithQs(partial, qs), trigger)
}

func Post(url string) Ren {
	return Attribute("hx-post", url)
}

func PostOnTrigger(url string, trigger string) Ren {
	return AttributeList(Attribute("hx-post", url), Trigger(trigger))
}

func PostOnClick(url string) Ren {
	return PostOnTrigger(url, "click")
}

func PostPartialOnClick(partial func(ctx *RequestContext) *Partial) Ren {
	return PostOnClick(GetPartialPath(partial))
}

func PostPartialOnClickQs(partial func(ctx *RequestContext) *Partial, qs string) Ren {
	return PostOnClick(GetPartialPathWithQs(partial, qs))
}

func Trigger(trigger string) *AttributeMap {
	return Attribute("hx-trigger", trigger)
}

func TextF(format string, args ...interface{}) Ren {
	return Text(fmt.Sprintf(format, args...))
}

func Text(text string) *TextContent {
	return NewTextContent(text)
}

func Pf(format string, args ...interface{}) Ren {
	return P(Text(fmt.Sprintf(format, args...)))
}

func Target(target string) Ren {
	return Attribute("hx-target", target)
}

func Name(name string) Ren {
	return Attribute("name", name)
}

func Confirm(message string) Ren {
	return Attribute("hx-confirm", message)
}

func Href(path string) Ren {
	return Attribute("href", path)
}

func Type(name string) Ren {
	return Attribute("type", name)
}

func Placeholder(placeholder string) Ren {
	return Attribute("placeholder", placeholder)
}

func OutOfBandSwap(selector string) Ren {
	return Attribute("hx-swap-oob",
		Ternary(selector == "", "true", selector))
}

func Click(value string) Ren {
	return Attribute("onclick", value)
}

func Tag(tag string, children ...Ren) *Element {
	return &Element{
		tag:      tag,
		children: children,
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

func ReplaceUrlHeader(url string) *Headers {
	return NewHeaders("HX-Replace-Url", url)
}

func CombineHeaders(headers ...*Headers) *Headers {
	m := make(Headers)
	for _, h := range headers {
		for k, v := range *h {
			m[k] = v
		}
	}
	return &m
}

func CurrentPath(ctx *RequestContext) string {
	current := ctx.Request().Header.Get("Hx-Current-Url")
	parsed, err := url.Parse(current)
	if err != nil {
		return ""
	}
	return parsed.Path
}

func PushQsHeader(ctx *RequestContext, key string, value string) *Headers {
	current := ctx.Request().Header.Get("Hx-Current-Url")
	parsed, err := url.Parse(current)
	if err != nil {
		return NewHeaders()
	}
	return NewHeaders("HX-Replace-Url", SetQueryParams(parsed.Path, map[string]string{
		key: value,
	}))
}

func NewHeaders(headers ...string) *Headers {
	if len(headers)%2 != 0 {
		return &Headers{}
	}
	m := make(Headers)
	for i := 0; i < len(headers); i++ {
		m[headers[i]] = headers[i+1]
		i++
	}
	return &m
}

func Checkbox(children ...Ren) Ren {
	return Input("checkbox", children...)
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

func Fragment(children ...Ren) *Element {
	return &Element{
		tag:      "",
		children: children,
	}
}

func AttributeList(children ...*AttributeMap) *AttributeMap {
	m := make(AttributeMap)
	for _, child := range children {
		for k, v := range *child {
			m[k] = v
		}
	}
	return &m
}

func AppendChildren(node *Element, children ...Ren) Ren {
	node.children = append(node.children, children...)
	return node

}

func Button(children ...Ren) *Element {
	return Tag("button", children...)
}

func Indicator(tag string) Ren {
	return Attribute("hx-indicator", tag)
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

func BeforeRequestSetHtml(children ...Ren) Ren {
	serialized := Render(Fragment(children...))
	return Attribute("hx-on::before-request", `this.innerHTML = '`+html.EscapeString(serialized)+`'`)
}

func BeforeRequestSetAttribute(key string, value string) Ren {
	return Attribute("hx-on::before-request", `this.setAttribute('`+key+`', '`+value+`')`)
}

func OnMutationErrorSetText(text string) Ren {
	return Attribute("hx-on::mutation-error", `this.innerText = '`+text+`'`)
}

func BeforeRequestSetText(text string) Ren {
	return Attribute("hx-on::before-request", `this.innerText = '`+text+`'`)
}

func AfterRequestSetText(text string) Ren {
	return Attribute("hx-on::after-request", `this.innerText = '`+text+`'`)
}

func AfterRequestRemoveAttribute(key string, value string) Ren {
	return Attribute("hx-on::after-request", `this.removeAttribute('`+key+`')`)
}

func IfQueryParam(key string, node *Element) Ren {
	return Fragment(Attribute("hx-if-qp:"+key, "true"), node)
}

func Hidden() Ren {
	return Attribute("style", "display:none")
}

func AfterRequestSetHtml(children ...Ren) Ren {
	serialized := Render(Fragment(children...))
	return Attribute("hx-on::after-request", `this.innerHTML = '`+html.EscapeString(serialized)+`'`)
}

func Children(children ...Ren) *ChildList {
	return NewChildList(children...)
}

func Label(text string) *Element {
	return Tag("label", Text(text))
}

func If(condition bool, node Ren) Ren {
	if condition {
		return node
	} else {
		return Empty()
	}
}

func IfElse(condition bool, node Ren, node2 Ren) Ren {
	if condition {
		return node
	} else {
		return node2
	}
}

func IfElseLazy(condition bool, cb1 func() Ren, cb2 func() Ren) Ren {
	if condition {
		return cb1()
	} else {
		return cb2()
	}
}

func GetTriggerName(ctx *RequestContext) string {
	return ctx.Request().Header.Get("HX-Trigger-Name")
}

func IfHtmxRequest(ctx *RequestContext, node Ren) Ren {
	if ctx.Get("HX-Request") != "" {
		return node
	}
	return Empty()
}

type SwapArg struct {
	Selector string
	Content  *Element
}

func NewSwap(selector string, content *Element) SwapArg {
	return SwapArg{
		Selector: selector,
		Content:  content,
	}
}

func OobSwap(ctx *RequestContext, content *Element) *Element {
	return OobSwapWithSelector(ctx, "", content)
}

func OobSwapWithSelector(ctx *RequestContext, selector string, content *Element) *Element {
	if ctx == nil || ctx.Get("HX-Request") == "" {
		return Empty()
	}
	return content.AppendChild(OutOfBandSwap(selector))
}

func SwapMany(ctx *RequestContext, args ...SwapArg) Ren {
	if ctx.Get("HX-Request") == "" {
		return Empty()
	}
	for _, arg := range args {
		arg.Content.AppendChild(OutOfBandSwap(arg.Selector))
	}
	return Fragment(Map(args, func(arg SwapArg) Ren {
		return arg.Content
	})...)
}

type OnRequestSwapArgs struct {
	Target        string
	Get           string
	Default       *Element
	BeforeRequest *Element
	AfterRequest  *Element
}

func OnRequestSwap(args OnRequestSwapArgs) Ren {
	return Div(args.Default,
		BeforeRequestSetHtml(args.BeforeRequest),
		AfterRequestSetHtml(args.AfterRequest),
		Get(args.Get),
		Target(args.Target),
	)
}
