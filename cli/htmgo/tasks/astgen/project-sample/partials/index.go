package partials

import "github.com/maddalax/htmgo/framework/h"

func CountersPartial(ctx *h.RequestContext) *h.Partial {
	return h.NewPartial(
		h.Div(
			h.Text("my counter"),
		),
	)
}
