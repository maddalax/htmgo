package partials

import (
	"github.com/maddalax/htmgo/framework/h"
	"time"
)

func CurrentTimePartial(ctx *h.RequestContext) *h.Partial {
	now := time.Now()
	return h.NewPartial(
		h.Div(
			h.Class("flex gap-1 items-center"),
			h.Pf("The current time is "),
			h.Span(
				h.Text(now.Format(time.RFC3339)),
				h.Class("font-bold"),
			),
		),
	)
}
