package h

type Node struct {
	tag        string
	attributes map[string]string
	children   []*Node
	text       string
	value      string
}

func Class(value string) *Node {
	return &Node{
		tag:   "class",
		value: value,
	}
}

func Id(value string) *Node {
	return Attribute("id", value)
}

func Attribute(key string, value string) *Node {
	return &Node{
		tag: "attribute",
		attributes: map[string]string{
			key: value,
		},
	}
}

func Get(url string) *Node {
	return Attribute("hx-get", url)
}

func Post(url string) *Node {
	return Attribute("hx-post", url)
}

func Trigger(trigger string) *Node {
	return Attribute("hx-trigger", trigger)
}

func Target(target string) *Node {
	return Attribute("hx-target", target)
}

func Name(name string) *Node {
	return Attribute("name", name)
}

func Confirm(message string) *Node {
	return Attribute("hx-confirm", message)
}

func Href(path string) *Node {
	return Attribute("href", path)
}

func Type(name string) *Node {
	return Attribute("type", name)
}

func Placeholder(placeholder string) *Node {
	return Attribute("placeholder", placeholder)
}

func Swap(swap string) *Node {
	return Attribute("hx-swap", swap)
}

func Click(value string) *Node {
	return Attribute("onclick", value)
}

func Page(children ...*Node) *Node {
	return &Node{
		tag:      "html",
		children: children,
	}
}

func Head(children ...*Node) *Node {
	return &Node{
		tag:      "head",
		children: children,
	}
}

func Body(children ...*Node) *Node {
	return &Node{
		tag:      "body",
		children: children,
	}
}

func Script(url string) *Node {
	return &Node{
		tag: "script",
		attributes: map[string]string{
			"src": url,
		},
		children: make([]*Node, 0),
	}
}

func Div(children ...*Node) *Node {
	return &Node{
		tag:      "div",
		children: children,
	}
}

func Input(inputType string, children ...*Node) *Node {
	return &Node{
		tag: "input",
		attributes: map[string]string{
			"type": inputType,
		},
		children: children,
	}
}

func HStack(children ...*Node) *Node {
	return &Node{
		tag: "div",
		attributes: map[string]string{
			"class": "flex gap-2",
		},
		children: children,
	}
}

func VStack(children ...*Node) *Node {
	return &Node{
		tag: "div",
		attributes: map[string]string{
			"class": "flex flex-col gap-2",
		},
		children: children,
	}
}

func List[T any](items []T, mapper func(item T) *Node) *Node {
	node := &Node{
		tag:      "",
		children: make([]*Node, len(items)),
	}
	for index, value := range items {
		node.children[index] = mapper(value)
	}
	return node
}

func Fragment(children ...*Node) *Node {
	return &Node{
		tag:      "",
		children: children,
	}
}

func Button(children ...*Node) *Node {
	return &Node{
		tag:      "button",
		children: children,
	}
}

func P(text string, children ...*Node) *Node {
	return &Node{
		tag:      "p",
		children: children,
		text:     text,
	}
}

func A(text string, children ...*Node) *Node {
	return &Node{
		tag:      "a",
		children: children,
		text:     text,
	}
}

func Empty() *Node {
	return &Node{
		tag: "",
	}
}

func If(condition bool, node *Node) *Node {
	if condition {
		return node
	} else {
		return Empty()
	}
}

func IfElse(condition bool, node *Node, node2 *Node) *Node {
	if condition {
		return node
	} else {
		return node2
	}
}
