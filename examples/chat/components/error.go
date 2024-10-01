package components

import "github.com/maddalax/htmgo/framework/h"

func FormError(error string) *h.Element {
	return h.Div(
		h.Id("form-error"),
		h.Text(error),
		h.If(error != "", h.Class("p-4 bg-rose-400 text-white rounded")),
	)
}
