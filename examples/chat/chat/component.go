package chat

import "github.com/maddalax/htmgo/framework/h"

func MessageRow(text string) *h.Element {
	return h.Div(
		h.Attribute("hx-swap-oob", "beforeend"),
		h.Class("flex flex-col gap-2 w-full"),
		h.Id("messages"),
		h.Pf(text),
	)
}
