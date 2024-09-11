package h

import (
	"fmt"
	"strings"
	"time"
)

const FlagSkip = "skip"
const FlagText = "text"
const FlagRaw = "raw"
const FlagAttributeList = "attribute-list"

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

		for _, child := range node.children {

			if child == nil {
				continue
			}

			if child.tag == "class" {
				insertAttribute(node, "class", child.value)
				child.tag = FlagSkip
			}

			if child.tag == FlagAttributeList {
				for _, gc := range child.children {
					for key, value := range gc.attributes {
						insertAttribute(node, key, value)
					}
					gc.tag = FlagSkip
				}
				child.tag = FlagSkip
			}

			if child.tag == "attribute" {
				for key, value := range child.attributes {
					insertAttribute(node, key, value)
				}
				child.tag = FlagSkip
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

		if child.tag == FlagText {
			page.builder.WriteString(child.text)
			continue
		}
		if child.tag == FlagRaw {
			page.builder.WriteString(child.value)
			continue
		}
		if child.tag != FlagSkip {
			page.renderNode(child)
		}
	}
	if node.tag != "" {
		page.builder.WriteString(fmt.Sprintf("</%s>", node.tag))
	}
}

func Render(node *Node) string {
	start := time.Now()
	builder := strings.Builder{}
	page := Builder{
		builder: &builder,
		root:    node,
	}
	page.render()
	d := page.builder.String()
	duration := time.Since(start)
	fmt.Printf("render took %s\n", duration)
	return d
}
