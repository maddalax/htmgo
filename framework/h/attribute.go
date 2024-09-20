package h

type AttributeR struct {
	Name  string
	Value string
}

func NewAttribute(name string, value string) *AttributeR {
	return &AttributeR{
		Name:  name,
		Value: value,
	}
}

type TextContent struct {
	Content string
}

func NewTextContent(content string) *TextContent {
	return &TextContent{
		Content: content,
	}
}

type RawContent struct {
	Content string
}

func NewRawContent(content string) *RawContent {
	return &RawContent{
		Content: content,
	}
}

type ChildList struct {
	Children []Ren
}

func NewChildList(children ...Ren) *ChildList {
	return &ChildList{
		Children: children,
	}
}
