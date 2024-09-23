package partials

import (
	"github.com/maddalax/htmgo/framework/h"
)

func SamplePartial(ctx *h.RequestContext) *h.Partial {
	return h.NewPartial(h.Div(h.P(h.Text(" asdasasds"))))
}

func NewPartial(ctx *h.RequestContext) *h.Partial {
	return h.NewPartial(h.Div(h.P(h.Text("This sadsl."))))
}

func NewPartial2(ctx *h.RequestContext) *h.Partial {
	return h.NewPartial(h.Div(h.P(h.Text("This sasdsadasdwl."))))
}
