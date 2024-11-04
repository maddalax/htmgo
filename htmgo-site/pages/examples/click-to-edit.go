package examples

import (
	"github.com/maddalax/htmgo/framework/h"
)

func ClickToEditExample(ctx *h.RequestContext) *h.Page {
	SetSnippet(ctx, &ClickToEditSnippet)
	return Index(ctx)
}
