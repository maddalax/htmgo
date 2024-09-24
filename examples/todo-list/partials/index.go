package partials

import (
	"github.com/maddalax/htmgo/framework/h"
	"time"
)

func CurrentTimePartial(ctx *h.RequestContext) *h.Partial {
	now := time.Now()
	return h.NewPartial(
		h.Div(
			h.Pf("The current time is %s", now.Format(time.RFC3339)),
		),
	)
}

func NewPartial(ctx *h.RequestContext) *h.Partial {
	return h.NewPartial(h.Div(h.P(h.Text("This sadsl."))))
}

func NewPartial2(ctx *h.RequestContext) *h.Partial {
	return h.NewPartial(h.Div(h.P(h.Text("This sasdsadasdwl."))))
}
