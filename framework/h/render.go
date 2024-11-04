package h

import (
	"strings"
)

type Ren interface {
	Render(context *RenderContext)
}

type RenderOptions struct {
	doctype bool
}

type RenderOpt func(context *RenderContext, opt *RenderOptions)

func WithDocType() RenderOpt {
	return func(context *RenderContext, opt *RenderOptions) {
		opt.doctype = true
	}
}

// Render renders the given node recursively, and returns the resulting string.
func Render(node Ren, opts ...RenderOpt) string {
	builder := &strings.Builder{}
	context := &RenderContext{
		builder: builder,
	}
	options := &RenderOptions{}

	for _, opt := range opts {
		opt(context, options)
	}

	if options.doctype {
		builder.WriteString("<!DOCTYPE html>")
	}

	node.Render(context)
	return builder.String()
}
