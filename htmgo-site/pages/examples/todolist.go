package examples

import (
	"github.com/maddalax/htmgo/framework/h"
)

func TodoListExample(ctx *h.RequestContext) *h.Page {
	SetSnippet(ctx, &TodoListSnippet)
	return Index(ctx)
}
