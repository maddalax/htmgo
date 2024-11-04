package examples

import (
	"github.com/maddalax/htmgo/framework/h"
)

func FormWithBlurValidation(ctx *h.RequestContext) *h.Page {
	SetSnippet(ctx, &FormWithBlurValidationSnippet)
	return Index(ctx)
}
