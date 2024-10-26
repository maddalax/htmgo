package htmltogo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIdentHRen(t *testing.T) {
	input := `
		package main
		import (
			"github.com/maddalax/htmgo/framework/h"
		)
		func Button(props ButtonProps) h.Ren {
			return h.Div(
				h.Div(h.Div(),h.P(),h.P(),
				),
			)
		}
	`
	indented := Indent(input)
	assert.Equal(t, `package main

import (
	"github.com/maddalax/htmgo/framework/h"
)

func Button(props ButtonProps) h.Ren {
	return h.Div(
		h.Div(
			h.Div(),
			h.P(),
			h.P(),
		),
	)
}
`, indented)
}
