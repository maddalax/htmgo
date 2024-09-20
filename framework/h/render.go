package h

import (
	"fmt"
	"strings"
	"time"
)

func Render(node Ren) string {
	start := time.Now()
	builder := &strings.Builder{}
	node.Render(builder)
	duration := time.Since(start)
	fmt.Printf("render took %d microseconds\n", duration.Microseconds())
	return builder.String()
}
