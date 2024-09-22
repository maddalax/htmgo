package h

import (
	"fmt"
	"strings"
	"time"
)

type Ren interface {
	Render(context *RenderContext)
}

func Render(node Ren) string {
	start := time.Now()
	builder := &strings.Builder{}
	context := &RenderContext{
		builder: builder,
	}
	node.Render(context)
	duration := time.Since(start)
	fmt.Printf("render took %d microseconds\n", duration.Microseconds())
	return builder.String()
}
