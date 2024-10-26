package h

import (
	"strings"
)

type Ren interface {
	Render(context *RenderContext)
}

// Render renders the given node recursively, and returns the resulting string.
func Render(node Ren) string {
	builder := &strings.Builder{}
	context := &RenderContext{
		builder: builder,
	}
	node.Render(context)
	return builder.String()
}
