package h

import (
	"fmt"
	"strings"

	"github.com/maddalax/htmgo/framework/hx"
	"github.com/maddalax/htmgo/framework/internal/datastructure"
)

type AttributeMap = map[string]any

type AttributeMapOrdered struct {
	data *datastructure.OrderedMap[string, string]
}

func (m *AttributeMapOrdered) Set(key string, value any) {
	switch v := value.(type) {
	case string:
		m.data.Set(key, v)
	case *AttributeMap:
		for k, v2 := range *v {
			m.Set(k, v2)
		}
	case *AttributeMapOrdered:
		v.Each(func(k string, v2 string) {
			m.Set(k, v2)
		})
	case *AttributeR:
		m.data.Set(v.Name, v.Value)
	default:
		m.data.Set(key, fmt.Sprintf("%v", value))
	}
}

func (m *AttributeMapOrdered) Each(cb func(key string, value string)) {
	m.data.Each(func(key string, value string) {
		cb(key, value)
	})
}

func (m *AttributeMapOrdered) Entries() []datastructure.MapEntry[string, string] {
	return m.data.Entries()
}

func NewAttributeMap(pairs ...string) *AttributeMapOrdered {
	m := datastructure.NewOrderedMap[string, string]()
	if len(pairs)%2 == 0 {
		for i := 0; i < len(pairs); i++ {
			m.Set(pairs[i], pairs[i+1])
			i++
		}
	}
	return &AttributeMapOrdered{data: m}
}

func NoSwap() *AttributeR {
	return Attribute("hx-swap", "none")
}

func Attribute(key string, value string) *AttributeR {
	return &AttributeR{
		Name:  key,
		Value: value,
	}
}

func AttributeList(children ...*AttributeR) *AttributeMapOrdered {
	m := NewAttributeMap()
	for _, c := range children {
		m.Set(c.Name, c.Value)
	}
	return m
}

func Attributes(attributes *AttributeMap) *AttributeMapOrdered {
	m := NewAttributeMap()
	for k, v := range *attributes {
		m.Set(k, v)
	}
	return m
}

func AttributePairs(pairs ...string) *AttributeMapOrdered {
	return NewAttributeMap(pairs...)
}

func Checked() Ren {
	return Attribute("checked", "")
}

func Id(value string) Ren {
	value = strings.TrimPrefix(value, "#")
	return Attribute("id", value)
}

func Disabled() *AttributeR {
	return Attribute("disabled", "")
}

func HxTarget(target string) Ren {
	return Attribute(hx.TargetAttr, target)
}

func Name(name string) Ren {
	return Attribute("name", name)
}

func HxConfirm(message string) Ren {
	return Attribute(hx.ConfirmAttr, message)
}

// HxInclude https://htmx.org/attributes/hx-include/
func HxInclude(selector string) Ren {
	return Attribute(hx.IncludeAttr, selector)
}

func HxIndicator(tag string) *AttributeR {
	return Attribute(hx.IndicatorAttr, tag)
}

func TriggerChildren() *AttributeR {
	return HxExtension("trigger-children")
}

func HxTriggerString(triggers ...string) *AttributeR {
	trigger := hx.NewStringTrigger(strings.Join(triggers, ", "))
	return Attribute(hx.TriggerAttr, trigger.ToString())
}

func HxTrigger(opts ...hx.TriggerEvent) *AttributeR {
	return Attribute(hx.TriggerAttr, hx.NewTrigger(opts...).ToString())
}

func HxTriggerClick(opts ...hx.Modifier) *AttributeR {
	return HxTrigger(hx.OnClick(opts...))
}

func HxExtension(value string) *AttributeR {
	return Attribute(hx.ExtAttr, value)
}

func HxExtensions(value ...string) Ren {
	return Attribute(hx.ExtAttr, strings.Join(value, ","))
}

func JoinExtensions(attrs ...*AttributeR) Ren {
	return JoinAttributes(", ", attrs...)
}

func JoinAttributes(sep string, attrs ...*AttributeR) *AttributeR {
	values := make([]string, 0, len(attrs))
	for _, a := range attrs {
		values = append(values, a.Value)
	}
	return Attribute(attrs[0].Name, strings.Join(values, sep))
}

func Href(path string) Ren {
	return Attribute("href", path)
}

func Target(target string) Ren {
	return Attribute("target", target)
}

func D(value string) Ren {
	return Attribute("d", value)
}

func Alt(value string) Ren {
	return Attribute("alt", value)
}

func For(value string) Ren {
	return Attribute("for", value)
}

func Type(name string) Ren {
	return Attribute("type", name)
}

func Placeholder(placeholder string) Ren {
	return Attribute("placeholder", placeholder)
}

func Hidden() Ren {
	return Attribute("style", "display:none")
}

func Class(value ...string) *AttributeR {
	return Attribute("class", MergeClasses(value...))
}

func ClassF(format string, args ...interface{}) *AttributeR {
	atr := fmt.Sprintf(format, args...)
	return Attribute("class", atr)
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

func Boost() Ren {
	return Attribute(hx.BoostAttr, "true")
}

func IfQueryParam(key string, node *Element) Ren {
	return Fragment(Attribute("hx-if-qp:"+key, "true"), node)
}

func ReadOnly() *AttributeR {
	return Attribute("readonly", "")
}

func Required() *AttributeR {
	return Attribute("required", "")
}

func Multiple() *AttributeR {
	return Attribute("multiple", "")
}

func Selected() *AttributeR {
	return Attribute("selected", "")
}

func MaxLength(value int) *AttributeR {
	return Attribute("maxlength", fmt.Sprintf("%d", value))
}

func MinLength(value int) *AttributeR {
	return Attribute("minlength", fmt.Sprintf("%d", value))
}

func Size(value int) *AttributeR {
	return Attribute("size", fmt.Sprintf("%d", value))
}

func Width(value int) *AttributeR {
	return Attribute("width", fmt.Sprintf("%d", value))
}

func Height(value int) *AttributeR {
	return Attribute("height", fmt.Sprintf("%d", value))
}

func Download(value bool) *AttributeR {
	return Attribute("download", fmt.Sprintf("%t", value))
}

func Rel(value string) *AttributeR {
	return Attribute("rel", value)
}

func Pattern(value string) *AttributeR {
	return Attribute("pattern", value)
}

func Action(value string) *AttributeR {
	return Attribute("action", value)
}

func Method(value string) *AttributeR {
	return Attribute("method", value)
}

func Enctype(value string) *AttributeR {
	return Attribute("enctype", value)
}

func AutoComplete(value string) *AttributeR {
	return Attribute("autocomplete", value)
}

func AutoFocus() *AttributeR {
	return Attribute("autofocus", "")
}

func NoValidate() *AttributeR {
	return Attribute("novalidate", "")
}

func Step(value string) *AttributeR {
	return Attribute("step", value)
}

func Max(value string) *AttributeR {
	return Attribute("max", value)
}

func Min(value string) *AttributeR {
	return Attribute("min", value)
}

func Cols(value int) *AttributeR {
	return Attribute("cols", fmt.Sprintf("%d", value))
}

func Rows(value int) *AttributeR {
	return Attribute("rows", fmt.Sprintf("%d", value))
}

func Wrap(value string) *AttributeR {
	return Attribute("wrap", value)
}

func Role(value string) *AttributeR {
	return Attribute("role", value)
}

func AriaLabel(value string) *AttributeR {
	return Attribute("aria-label", value)
}

func AriaHidden(value bool) *AttributeR {
	return Attribute("aria-hidden", fmt.Sprintf("%t", value))
}

func TabIndex(value int) *AttributeR {
	return Attribute("tabindex", fmt.Sprintf("%d", value))
}
