package formatter

import (
	"go/format"
	"html-to-htmgo/internal/domain"
)

type Formatter struct {
}

func (f Formatter) Format(node *domain.CustomNode) string {
	b := []byte(`package main
import (
	"github.com/maddalax/htmgo/framework/h"
)
func MyComponent() *h.Element {
	return ` + node.String() + `
}`)
	dist, err := format.Source(b)
	if err != nil {
		return string(b)
	}

	return string(dist)
}

func New() Formatter {
	return Formatter{}
}
