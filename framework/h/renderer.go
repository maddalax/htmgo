package h

import (
	"fmt"
	"strings"
)

func (node *Element) Render() string {
	builder := &strings.Builder{}

	// some elements may not have a tag, such as a Fragment
	if node.tag != "" {
		builder.WriteString("<" + node.tag)
		builder.WriteString(" ")
		for name, value := range node.attributes {
			builder.WriteString(NewAttribute(name, value).Render())
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
			builder.WriteString(child.(*AttributeMap).Render())
		case *LifeCycle:
			builder.WriteString(child.(*LifeCycle).Render())
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
			builder.WriteString(child.Render())
		}
	}

	if node.tag != "" {
		builder.WriteString("</" + node.tag + ">")
	}

	str := builder.String()
	return str
}

func (a *AttributeR) Render() string {
	return fmt.Sprintf(`%s="%s"`, a.Name, a.Value)
}

func (t *TextContent) Render() string {
	return t.Content
}

func (r *RawContent) Render() string {
	return r.Content
}

func (c *ChildList) Render() string {
	builder := &strings.Builder{}
	for _, child := range c.Children {
		builder.WriteString(child.Render())
	}
	str := builder.String()
	return str
}

func (m *AttributeMap) Render() string {
	builder := &strings.Builder{}
	m2 := m.ToMap()

	for k, v := range m2 {
		builder.WriteString(NewAttribute(k, v).Render())
	}

	str := builder.String()
	return str
}

func (l *LifeCycle) Render() string {
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

	result := Children(children...).Render()
	return result
}

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
