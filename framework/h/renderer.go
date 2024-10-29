package h

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/hx"
	"html"
	"html/template"
	"strings"
)

type CustomElement = string

var (
	CachedNodeTag        CustomElement = "htmgo_cache_node"
	CachedNodeByKeyEntry CustomElement = "htmgo_cached_node_by_key_entry"
)

/*
*
void tags are tags that cannot have children
*/
var voidTags = map[string]bool{
	"area":   true,
	"base":   true,
	"br":     true,
	"col":    true,
	"embed":  true,
	"hr":     true,
	"img":    true,
	"input":  true,
	"link":   true,
	"meta":   true,
	"source": true,
	"track":  true,
	"wbr":    true,
}

type ScriptEntry struct {
	Body    string
	ChildOf *Element
}

type RenderContext struct {
	builder        *strings.Builder
	scripts        []ScriptEntry
	currentElement *Element
}

func (ctx *RenderContext) AddScript(funcName string, body string) {
	script := fmt.Sprintf(`
	<script id="%s">
		function %s(self, event) {
				let e = event;
				%s
		}
	</script>`, funcName, funcName, body)

	ctx.scripts = append(ctx.scripts, ScriptEntry{
		Body:    script,
		ChildOf: ctx.currentElement,
	})
}

func (node *Element) Render(context *RenderContext) {
	if node == nil {
		return
	}

	context.currentElement = node

	if node.tag == CachedNodeTag {
		meta := node.meta.(*CachedNode)
		meta.Render(context)
		return
	}

	if node.tag == CachedNodeByKeyEntry {
		meta := node.meta.(*ByKeyEntry)
		meta.Render(context)
		return
	}

	// some elements may not have a tag, such as a Fragment
	if node.tag != "" {
		context.builder.WriteString("<")
		context.builder.WriteString(node.tag)
		node.attributes.Each(func(key string, value string) {
			NewAttribute(key, value).Render(context)
		})
	}

	totalChildren := 0
	shouldFlatten := false
	for _, child := range node.children {
		switch c := child.(type) {
		case *ChildList:
			shouldFlatten = true
			totalChildren += len(c.Children)
		default:
			totalChildren++
		}
	}

	if shouldFlatten {
		// first pass, flatten the children
		flatChildren := make([]Ren, totalChildren)
		index := 0
		for _, child := range node.children {
			switch c := child.(type) {
			case *ChildList:
				for _, ren := range c.Children {
					flatChildren[index] = ren
					index++
				}
			default:
				flatChildren[index] = child
				index++
			}
		}

		node.children = flatChildren
	}

	// second pass, render any attributes within the tag
	for _, child := range node.children {
		switch child.(type) {
		case *AttributeMapOrdered:
			child.Render(context)
		case *AttributeR:
			child.Render(context)
		case *LifeCycle:
			child.Render(context)
		}
	}

	// close the tag
	if node.tag != "" {
		if voidTags[node.tag] {
			context.builder.WriteString("/")
		}
		context.builder.WriteString(">")
	}

	// void elements do not have children
	if !voidTags[node.tag] {
		// render the children elements that are not attributes
		for _, child := range node.children {
			switch child.(type) {
			case *AttributeMapOrdered:
				continue
			case *AttributeR:
				continue
			case *LifeCycle:
				continue
			default:
				child.Render(context)
			}
		}
	}

	if node.tag != "" {
		renderScripts(context, node)
		if !voidTags[node.tag] {
			context.builder.WriteString("</")
			context.builder.WriteString(node.tag)
			context.builder.WriteString(">")
		}
	}
}

func renderScripts(context *RenderContext, parent *Element) {
	if len(context.scripts) == 0 {
		return
	}
	notWritten := make([]ScriptEntry, 0)
	for _, script := range context.scripts {
		if script.ChildOf == parent {
			context.builder.WriteString(script.Body)
		} else {
			notWritten = append(notWritten, script)
		}
	}
	context.scripts = notWritten
}

func (a *AttributeR) Render(context *RenderContext) {
	context.builder.WriteString(" ")
	context.builder.WriteString(a.Name)
	if a.Value != "" {
		context.builder.WriteString(`=`)
		context.builder.WriteString(`"`)
		context.builder.WriteString(html.EscapeString(a.Value))
		context.builder.WriteString(`"`)
	}
}

func (t *TextContent) Render(context *RenderContext) {
	context.builder.WriteString(template.HTMLEscapeString(t.Content))
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

func (m *AttributeMapOrdered) Render(context *RenderContext) {
	m.Each(func(key string, value string) {
		NewAttribute(key, value).Render(context)
	})
}

func (l *LifeCycle) fromAttributeMap(event string, key string, value string, context *RenderContext) {

	if key == hx.GetAttr || key == hx.PatchAttr || key == hx.PostAttr {
		HxTriggerString(hx.ToHtmxTriggerName(event)).Render(context)
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
				m[event] += fmt.Sprintf("var self=this;var e=event;%s;", c.Command)
			case ComplexJsCommand:
				context.AddScript(c.TempFuncName, c.Command)
				m[event] += fmt.Sprintf("%s(this, event);", c.TempFuncName)
			case *AttributeMapOrdered:
				c.Each(func(key string, value string) {
					l.fromAttributeMap(event, key, value, context)
				})
			case *AttributeR:
				l.fromAttributeMap(event, c.Name, c.Value, context)
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
