package examples

import (
	"github.com/maddalax/htmgo/framework/h"
)

func JsHideChildrenOnClickPage(ctx *h.RequestContext) *h.Page {
	SetSnippet(ctx, &JsHideChildrenOnClick)
	return Index(ctx)
}
