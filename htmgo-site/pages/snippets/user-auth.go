package snippets

import (
	"github.com/maddalax/htmgo/framework/h"
)

func UserAuthExample(ctx *h.RequestContext) *h.Page {
	SetSnippet(ctx, &UserAuthSnippet)
	return Index(ctx)
}
