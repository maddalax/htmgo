package h

type AttributeR struct {
	Name  string
	Value string
}

type TextContent struct {
	Content string
}

type RawContent struct {
	Content string
}

type ChildList struct {
	Children []Ren
}

func NewAttribute(name string, value string) *AttributeR {
	return &AttributeR{
		Name:  name,
		Value: value,
	}
}

func NewRawContent(content string) *RawContent {
	return &RawContent{
		Content: content,
	}
}

func NewTextContent(content string) *TextContent {
	return &TextContent{
		Content: content,
	}
}

func NewChildList(children ...Ren) *ChildList {
	return &ChildList{
		Children: children,
	}
}
