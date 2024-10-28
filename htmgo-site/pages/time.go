package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/partials"
)

func CurrentTimePage(ctx *h.RequestContext) *h.Page {
	return RootPage(
		ctx,
		h.GetPartial(
			partials.CurrentTimePartial,
			"load, every 1s"),
	)
}
