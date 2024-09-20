package h

import (
	"fmt"
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

func (l *LifeCycle) Render(builder *strings.Builder) {
	m := make(map[string]string)

	for event, commands := range l.handlers {
		m[event] = ""
		for _, command := range commands {
			m[event] += fmt.Sprintf("%s;", command.Command)
		}
	}

	children := make([]Ren, 0)

	for event, js := range m {
		children = append(children, Attribute(event, js))
	}

	Children(children...).Render(builder)
}
