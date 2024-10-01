package chat

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"strings"
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

func ConnectedUsers(username string) *h.Element {
	return h.Ul(
		h.Attribute("hx-swap", "none"),
		h.Attribute("hx-swap-oob", "beforeend"),
		h.Id("connected-users"),
		h.Class("flex flex-col gap-2"),
		// This would be populated dynamically with connected users
		ConnectedUser(username, false),
	)
}

func ConnectedUser(username string, remove bool) *h.Element {
	id := fmt.Sprintf("connected-user-%s", strings.ReplaceAll(username, "#", "-"))
	if remove {
		return h.Div(h.Id(id), h.Attribute("hx-swap-oob", "delete"))
	}
	return h.Li(
		h.Id(id),
		h.Class("text-slate-700"),
		h.Text(username),
	)
}
