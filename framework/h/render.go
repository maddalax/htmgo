package h

import (
	"fmt"
	"strings"
	"time"
)

const FlagSkip = "skip"
const FlagText = "text"
const FlagRaw = "raw"
const FlagAttributeList = "x-attribute-list"
const FlagChildrenList = "x-children-list"

type Builder struct {
	builder *strings.Builder
	root    *Node
}

func (page Builder) render() {
	page.renderNode(page.root)
}

func insertAttribute(node *Node, name string, value string) {
	existing := node.attributes[name]
	if existing != "" {
		node.attributes[name] = existing + " " + value
	} else {
		node.attributes[name] = value
	}
}

func (page Builder) renderNode(node *Node) {
	if node.tag != "" {
		page.builder.WriteString(fmt.Sprintf("<%s", node.tag))
		index := 0

		if node.attributes == nil {
			node.attributes = map[string]string{}
		}

		flatChildren := make([]Renderable, 0)
		for _, child := range node.children {

			if child == nil {
				continue
			}

			c := child.Render()

			flatChildren = append(flatChildren, child)
			if c.tag == FlagChildrenList {
				for _, gc := range c.children {
					flatChildren = append(flatChildren, gc)
				}
				c.tag = FlagSkip
			}
		}

		if len(flatChildren) > 0 {
			node.children = flatChildren
		}

		for _, child := range node.children {

			if child == nil {
				continue
			}

			c := child.Render()

			if c.tag == "class" {
				insertAttribute(node, "class", c.value)
				c.tag = FlagSkip
			}

			if c.tag == FlagAttributeList {
				for _, gc := range c.children {
					gcr := gc.Render()
					for key, value := range gcr.attributes {
						insertAttribute(node, key, value)
					}
					gcr.tag = FlagSkip
				}
				c.tag = FlagSkip
			}

			if c.tag == "attribute" {
				for key, value := range c.attributes {
					insertAttribute(node, key, value)
				}
				c.tag = FlagSkip
			}
		}

		for key, value := range node.attributes {
			if index == 0 {
				page.builder.WriteString(" ")
			}
			page.builder.WriteString(key)
			page.builder.WriteString("=")
			page.builder.WriteRune('"')
			page.builder.WriteString(value)
			page.builder.WriteRune('"')
			if index < len(node.attributes) {
				page.builder.WriteRune(' ')
			}
			index += 1
		}
		page.builder.WriteString(">")
		if node.text != "" {
			page.builder.WriteString(node.text)
		}
	}
	for _, child := range node.children {

		if child == nil {
			continue
		}

		c := child.Render()

		if c.tag == FlagText {
			page.builder.WriteString(c.text)
			continue
		}
		if c.tag == FlagRaw {
			page.builder.WriteString(c.value)
			continue
		}
		if c.tag != FlagSkip {
			page.renderNode(c)
		}
	}
	if node.tag != "" {
		page.builder.WriteString(fmt.Sprintf("</%s>", node.tag))
	}
}

func Render(node Renderable) string {
	start := time.Now()
	builder := strings.Builder{}
	page := Builder{
		builder: &builder,
		root:    node.Render(),
	}
	page.render()
	d := page.builder.String()
	duration := time.Since(start)
	fmt.Printf("render took %d\n", duration.Microseconds())
	return d
}
