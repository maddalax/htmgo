package pages

import (
	"github.com/maddalax/htmgo/framework/h"
)

func IndexPage(ctx *h.RequestContext) *h.Page {
	return h.NewPage(MarkdownPage(ctx, "md/index.md"))
}
