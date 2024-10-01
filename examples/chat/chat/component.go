package chat

import (
	"github.com/maddalax/htmgo/framework/h"
	"time"
)

func MessageRow(message *Message) *h.Element {
	return h.Div(
		h.Attribute("hx-swap-oob", "beforeend"),
		h.Class("flex flex-col gap-2 w-full"),
		h.Id("messages"),
		h.Div(
			h.Class("flex gap-2 items-center"),
			h.Pf(message.UserName),
			h.Pf(message.CreatedAt.In(time.Local).Format("01/02 03:04 PM")),
			h.Pf(message.Message),
		),
	)
}
