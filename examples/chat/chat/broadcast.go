package chat

import (
	"chat/ws"
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
)

func StartListener(loader *service.Locator) {
	manager := service.Get[ws.SocketManager](loader)

	c := make(chan ws.MessageEvent)
	manager.Listen(c)

	for {
		select {
		case event := <-c:
			fmt.Printf("Received message from %s: %v\n", event.Id, event.Message)
			message := event.Message["message"].(string)
			if message == "" {
				continue
			}

			messageEle := h.Div(
				h.Attribute("hx-swap-oob", "beforeend"),
				h.Class("flex flex-col gap-2 w-full"),
				h.Id("messages"),
				h.Pf(message),
			)

			manager.BroadcastText(
				h.Render(
					h.Fragment(
						messageEle,
					)),
			)
		}
	}
}
