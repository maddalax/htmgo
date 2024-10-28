package snippets

import (
	"github.com/maddalax/htmgo/framework/h"
)

func ChatExample(ctx *h.RequestContext) *h.Page {
	SetSnippet(ctx, &ChatSnippet)
	return Index(ctx)
}
