package docs

import "github.com/maddalax/htmgo/framework/h"

func Index(ctx *h.RequestContext) *h.Page {
	ctx.Redirect("/docs/introduction", 302)
	return h.EmptyPage()
}
