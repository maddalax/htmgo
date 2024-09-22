package h

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/hx"
	"strings"
)

type AttributeMap map[string]any

func (m *AttributeMap) ToMap() map[string]string {
	result := make(map[string]string)
	for k, v := range *m {
		switch v.(type) {
		case AttributeMap:
			m2 := v.(*AttributeMap).ToMap()
			for _, a := range m2 {
				result[k] = a
			}
		case string:
			result[k] = v.(string)
		default:
			result[k] = fmt.Sprintf("%v", v)
		}
	}
	return result
}

func Attribute(key string, value string) *AttributeMap {
	return Attributes(&AttributeMap{key: value})
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

func Id(value string) Ren {
	if strings.HasPrefix(value, "#") {
		value = value[1:]
	}
	return Attribute("id", value)
}

func Disabled() Ren {
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

func HxIndicator(tag string) *AttributeMap {
	return Attribute(hx.IndicatorAttr, tag)
}

func TriggerChildren() Ren {
	return HxExtension("trigger-children")
}

func TriggerString(triggers ...string) *AttributeMap {
	trigger := hx.NewStringTrigger(strings.Join(triggers, ", "))
	return Attribute(hx.TriggerAttr, trigger.ToString())
}

func HxTrigger(opts ...hx.TriggerEvent) *AttributeMap {
	return Attribute(hx.TriggerAttr, hx.NewTrigger(opts...).ToString())
}

func HxTriggerClick(opts ...hx.Modifier) *AttributeMap {
	return HxTrigger(hx.OnClick(opts...))
}

func HxExtension(value string) Ren {
	return Attribute(hx.ExtAttr, value)
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

func Hidden() Ren {
	return Attribute("style", "display:none")
}

func Class(value ...string) *AttributeMap {
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

func Boost() Ren {
	return Attribute(hx.BoostAttr, "true")
}

func IfQueryParam(key string, node *Element) Ren {
	return Fragment(Attribute("hx-if-qp:"+key, "true"), node)
}
