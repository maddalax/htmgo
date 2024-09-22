package h

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/hx"
	"strings"
)

type RenderContext struct {
	builder *strings.Builder
	scripts []string
}

func (ctx *RenderContext) AddScript(funcName string, body string) {
	script := fmt.Sprintf(`
	<script id="%s">
		function %s() {
				%s
		}
	</script>`, funcName, funcName, body)
	ctx.scripts = append(ctx.scripts, script)
}

func (node *Element) Render(context *RenderContext) {
	// some elements may not have a tag, such as a Fragment

	if node.tag != "" {
		context.builder.WriteString("<" + node.tag)
		context.builder.WriteString(" ")

		for name, value := range node.attributes {
			NewAttribute(name, value).Render(context)
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
			child.Render(context)
		case *LifeCycle:
			child.Render(context)
		}
	}

	// close the tag
	if node.tag != "" {
		context.builder.WriteString(">")
	}

	// render the children elements that are not attributes
	for _, child := range node.children {
		switch child.(type) {
		case *AttributeMap:
			continue
		case *LifeCycle:
			continue
		default:
			child.Render(context)
		}
	}

	if node.tag != "" {
		renderScripts(context)
		context.builder.WriteString("</" + node.tag + ">")
	}
}

func renderScripts(context *RenderContext) {
	for _, script := range context.scripts {
		context.builder.WriteString(script)
	}
	context.scripts = []string{}
}

func (a *AttributeR) Render(context *RenderContext) {
	context.builder.WriteString(fmt.Sprintf(`%s='%s'`, a.Name, a.Value))
}

func (t *TextContent) Render(context *RenderContext) {
	context.builder.WriteString(t.Content)
}

func (r *RawContent) Render(context *RenderContext) {
	context.builder.WriteString(r.Content)
}

func (c *ChildList) Render(context *RenderContext) {
	for _, child := range c.Children {
		child.Render(context)
	}
}

func (j SimpleJsCommand) Render(context *RenderContext) {
	context.builder.WriteString(j.Command)
}

func (j ComplexJsCommand) Render(context *RenderContext) {
	context.builder.WriteString(j.Command)
}

func (p *Partial) Render(context *RenderContext) {
	p.Root.Render(context)
}

func (m *AttributeMap) Render(context *RenderContext) {
	m2 := m.ToMap()

	for k, v := range m2 {
		context.builder.WriteString(" ")
		NewAttribute(k, v).Render(context)
	}
}

func (l *LifeCycle) fromAttributeMap(event string, key string, value string, context *RenderContext) {

	if key == hx.GetAttr || key == hx.PatchAttr || key == hx.PostAttr {
		TriggerString(hx.ToHtmxTriggerName(event)).Render(context)
	}

	Attribute(key, value).Render(context)
}

func (l *LifeCycle) Render(context *RenderContext) {
	m := make(map[string]string)

	for event, commands := range l.handlers {
		m[event] = ""
		for _, command := range commands {
			switch c := command.(type) {
			case SimpleJsCommand:
				m[event] += fmt.Sprintf("%s;", c.Command)
			case ComplexJsCommand:
				context.AddScript(c.TempFuncName, c.Command)
				m[event] += fmt.Sprintf("%s();", c.TempFuncName)
			case *AttributeMap:
				for k, v := range c.ToMap() {
					l.fromAttributeMap(event, k, v, context)
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

	Children(children...).Render(context)
}
