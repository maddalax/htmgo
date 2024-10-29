package snippets

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/js"
)

func SetTextOnClick(ctx *h.RequestContext) *h.Partial {
	return h.NewPartial(
		h.Button(
			h.Text("Click to set text"),
			h.Class("bg-slate-900 text-white py-2 px-4 rounded"),
			h.OnClick(
				js.SetText("Hello World"),
			),
		),
	)
}
