package snippets

import (
	"github.com/maddalax/htmgo/framework/h"
)

func FormExample(ctx *h.RequestContext) *h.Page {
	SetSnippet(ctx, &FormWithLoadingState)
	return Index(ctx)
}
