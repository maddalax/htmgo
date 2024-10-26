package parser

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/maddalax/htmgo/tools/html-to-htmgo/internal/domain"
	"golang.org/x/net/html"
	"strings"
)

type Parser struct {
}

var ParseErr = errors.New("parse error")

func (p Parser) FromBytes(in []byte) (*domain.CustomNode, error) {
	hNode, err := html.Parse(bytes.NewReader(in))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ParseErr, err)
	}
	var findBody func(n *html.Node) *html.Node
	findBody = func(n *html.Node) *html.Node {
		if n.Data == "body" {
			return n
		}
		var e *html.Node
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			e = findBody(c)
		}
		return e
	}

	body := findBody(hNode)
	if body == nil {
		return nil, fmt.Errorf("%w", ParseErr)
	}

	var f func(*html.Node, *domain.CustomNode) *domain.CustomNode
	f = func(n *html.Node, cNode *domain.CustomNode) *domain.CustomNode {
		if n.Type == html.ElementNode {
			cNode.SetType(n.Data)

			for _, attr := range n.Attr {
				cNode.AddAttr(attr.Key, attr.Val)
			}
		}

		if n.Type == html.TextNode && len(strings.TrimSpace(n.Data)) > 0 {
			cNode.ParentNode.AddAttr("h.Text", strings.TrimSpace(n.Data))
		}

		var i uint
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			cNode.Nodes = append(cNode.Nodes, &domain.CustomNode{ParentNode: cNode, Level: cNode.Level + 1})
			cNode.Nodes[i] = f(c, cNode.Nodes[i])
			i++
		}
		return cNode
	}
	output := &domain.CustomNode{}
	out := f(body, output)

	return out, nil
}

func New() Parser {
	return Parser{}
}
