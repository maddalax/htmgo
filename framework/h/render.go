package h

import (
	"strings"
)

type Ren interface {
	Render(context *RenderContext)
}

func Render(node Ren) string {
	builder := &strings.Builder{}
	context := &RenderContext{
		builder: builder,
	}
	node.Render(context)
	return builder.String()
}
