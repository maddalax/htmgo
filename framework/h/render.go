package h

import (
	"fmt"
	"time"
)

func Render(node Ren) string {
	start := time.Now()
	html := node.Render()
	duration := time.Since(start)
	fmt.Printf("render took %d microseconds\n", duration.Microseconds())
	return html
}
