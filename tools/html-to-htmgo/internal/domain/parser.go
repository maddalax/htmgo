package domain

type HTMLParser interface {
	FromBytes(bytes []byte) (*CustomNode, error)
}
