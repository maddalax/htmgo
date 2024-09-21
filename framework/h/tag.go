package h

import (
	"encoding/json"
	"fmt"
	"github.com/maddalax/htmgo/framework/hx"
	"net/url"
	"strings"
)

type Qs struct {
	m map[string]string
}

func NewQs(pairs ...string) *Qs {
	q := &Qs{
		m: make(map[string]string),
	}
	if len(pairs)%2 != 0 {
		return q
	}
	for i := 0; i < len(pairs); i++ {
		q.m[pairs[i]] = pairs[i+1]
		i++
	}
	return q
}

func (q *Qs) Add(key string, value string) *Qs {
	q.m[key] = value
	return q
}

func (q *Qs) ToString() string {
	builder := strings.Builder{}
	index := 0
	for k, v := range q.m {
		builder.WriteString(k)
		builder.WriteString("=")
		builder.WriteString(v)
		if index < len(q.m)-1 {
			builder.WriteString("&")
		}
		index++
	}
	return builder.String()
}

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

func AttributePairs(pairs ...string) *AttributeMap {
	if len(pairs)%2 != 0 {
		return &AttributeMap{}
	}
	m := make(AttributeMap)
	for i := 0; i < len(pairs); i++ {
		m[pairs[i]] = pairs[i+1]
		i++
	}
	return &m
}

func Checked() Ren {
	return Attribute("checked", "true")
}

func Boost() Ren {
	return Attribute(hx.BoostAttr, "true")
}

func Attribute(key string, value string) *AttributeMap {
	return Attributes(&AttributeMap{key: value})
}

func TriggerChildren() Ren {
	return HxExtension("trigger-children")
}

func HxExtension(value string) Ren {
	return Attribute(hx.ExtAttr, value)
}

func Disabled() Ren {
	return Attribute("disabled", "")
}

func TriggerString(triggers ...string) *AttributeMap {
	trigger := hx.NewStringTrigger(strings.Join(triggers, ", "))
	return Attribute(hx.TriggerAttr, trigger.ToString())
}

func Trigger(opts ...hx.TriggerEvent) *AttributeMap {
	return Attribute(hx.TriggerAttr, hx.NewTrigger(opts...).ToString())
}

func TriggerClick(opts ...hx.Modifier) *AttributeMap {
	return Trigger(hx.OnClick(opts...))
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
	return Attribute(hx.TargetAttr, target)
}

func Name(name string) Ren {
	return Attribute("name", name)
}

func Confirm(message string) Ren {
	return Attribute(hx.ConfirmAttr, message)
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
	return Attribute(hx.SwapOobAttr,
		Ternary(selector == "", "true", selector))
}

func Click(value string) Ren {
	return Attribute("onclick", value)
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

func ReplaceUrlHeader(url string) *Headers {
	return NewHeaders(hx.ReplaceUrlHeader, url)
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
	current := ctx.Request().Header.Get(hx.CurrentUrlHeader)
	parsed, err := url.Parse(current)
	if err != nil {
		return ""
	}
	return parsed.Path
}

func PushQsHeader(ctx *RequestContext, key string, value string) *Headers {
	current := ctx.Request().Header.Get(hx.CurrentUrlHeader)
	parsed, err := url.Parse(current)
	if err != nil {
		return NewHeaders()
	}
	return NewHeaders(hx.ReplaceUrlHeader, SetQueryParams(parsed.Path, map[string]string{
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

func Fragment(children ...Ren) *ChildList {
	return Children(children...)
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

func Indicator(tag string) *AttributeMap {
	return Attribute(hx.IndicatorAttr, tag)
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

func IfQueryParam(key string, node *Element) Ren {
	return Fragment(Attribute("hx-if-qp:"+key, "true"), node)
}

func Hidden() Ren {
	return Attribute("style", "display:none")
}

func Children(children ...Ren) *ChildList {
	return NewChildList(children...)
}

func Label(text string) *Element {
	return Tag("label", Text(text))
}

func GetTriggerName(ctx *RequestContext) string {
	return ctx.Request().Header.Get("HX-Trigger-Name")
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
