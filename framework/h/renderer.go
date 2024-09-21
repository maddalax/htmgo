package h

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/hx"
	"strings"
)

func (node *Element) Render(builder *strings.Builder) {
	// some elements may not have a tag, such as a Fragment

	if node.tag != "" {
		builder.WriteString("<" + node.tag)
		builder.WriteString(" ")

		for name, value := range node.attributes {
			NewAttribute(name, value).Render(builder)
		}
	}

	// first pass, flatten the children
	flatChildren := make([]Ren, 0)
	for _, child := range node.children {
		switch child.(type) {
		case *ChildList:
			flatChildren = append(flatChildren, child.(*ChildList).Children...)
		default:
			flatChildren = append(flatChildren, child)
		}
	}

	node.children = flatChildren

	// second pass, render any attributes within the tag
	for _, child := range node.children {
		switch child.(type) {
		case *AttributeMap:
			child.Render(builder)
		case *LifeCycle:
			child.Render(builder)
		}
	}

	// close the tag
	if node.tag != "" {
		builder.WriteString(">")
	}

	// render the children elements that are not attributes
	for _, child := range node.children {
		switch child.(type) {
		case *AttributeMap:
			continue
		case *LifeCycle:
			continue
		default:
			child.Render(builder)
		}
	}

	if node.tag != "" {
		builder.WriteString("</" + node.tag + ">")
	}
}

func (a *AttributeR) Render(builder *strings.Builder) {
	builder.WriteString(fmt.Sprintf(`%s="%s"`, a.Name, a.Value))
}

func (t *TextContent) Render(builder *strings.Builder) {
	builder.WriteString(t.Content)
}

func (r *RawContent) Render(builder *strings.Builder) {
	builder.WriteString(r.Content)
}

func (c *ChildList) Render(builder *strings.Builder) {
	for _, child := range c.Children {
		child.Render(builder)
	}
}

func (m *AttributeMap) Render(builder *strings.Builder) {
	m2 := m.ToMap()

	for k, v := range m2 {
		builder.WriteString(" ")
		NewAttribute(k, v).Render(builder)
	}
}

func (l *LifeCycle) fromAttributeMap(event string, key string, value string, builder *strings.Builder) {

	if key == hx.GetAttr || key == hx.PatchAttr || key == hx.PostAttr {
		TriggerString(hx.ToHtmxTriggerName(event)).Render(builder)
	}

	Attribute(key, value).Render(builder)
}

func (l *LifeCycle) Render(builder *strings.Builder) {
	m := make(map[string]string)

	for event, commands := range l.handlers {
		m[event] = ""
		for _, command := range commands {
			switch c := command.(type) {
			case JsCommand:
				eventName := hx.FormatEventName(event, true)
				m[eventName] += fmt.Sprintf("%s;", c.Command)
			case *AttributeMap:
				for k, v := range c.ToMap() {
					l.fromAttributeMap(event, k, v, builder)
				}
			}
		}
	}

	children := make([]Ren, 0)

	for event, value := range m {
		if value != "" {
			children = append(children, Attribute(event, value))
		}
	}

	if len(children) == 0 {
		return
	}

	Children(children...).Render(builder)
}
