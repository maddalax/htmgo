package examples

import "github.com/maddalax/htmgo/framework/h"

func HackerNewsExample(ctx *h.RequestContext) *h.Page {
	SetSnippet(ctx, &HackerNewsSnippet)
	return Index(ctx)
}
