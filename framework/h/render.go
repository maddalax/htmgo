package h

import (
	"fmt"
	"strings"
	"time"
)

type Ren interface {
	Render(builder *strings.Builder)
}

func Render(node Ren) string {
	start := time.Now()
	builder := &strings.Builder{}
	node.Render(builder)
	duration := time.Since(start)
	fmt.Printf("render took %d microseconds\n", duration.Microseconds())
	return builder.String()
}
