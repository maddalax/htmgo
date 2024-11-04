package examples

import "github.com/maddalax/htmgo/framework/h"

func InputComponentExample(ctx *h.RequestContext) *h.Page {
	SetSnippet(ctx, &InputComponentSnippet)
	return Index(ctx)
}
