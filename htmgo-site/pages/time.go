package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/pages/base"
	"htmgo-site/partials"
)

func CurrentTimePage(ctx *h.RequestContext) *h.Page {
	return h.NewPage(
		base.RootPage(
			ctx,
			h.GetPartial(
				partials.CurrentTimePartial,
				"load, every 1s"),
		))
}
