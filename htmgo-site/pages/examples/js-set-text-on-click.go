package examples

import (
	"github.com/maddalax/htmgo/framework/h"
)

func JsSetTextOnClickPage(ctx *h.RequestContext) *h.Page {
	SetSnippet(ctx, &JsSetTextOnClick)
	return Index(ctx)
}
