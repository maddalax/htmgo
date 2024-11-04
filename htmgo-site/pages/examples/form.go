package examples

import (
	"github.com/maddalax/htmgo/framework/h"
)

func FormWithLoadingState(ctx *h.RequestContext) *h.Page {
	SetSnippet(ctx, &FormWithLoadingStateSnippet)
	return Index(ctx)
}
