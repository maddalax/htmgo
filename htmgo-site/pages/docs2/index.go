package docs2

import "github.com/maddalax/htmgo/framework/h"

func Index(ctx *h.RequestContext) *h.Page {
	ctx.Redirect("/docs2/introduction", 302)
	return h.EmptyPage()
}
