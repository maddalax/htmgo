package partials

import (
	"chat/chat"
	"chat/components"
	"github.com/maddalax/htmgo/framework/h"
)

func CreateOrJoinRoom(ctx *h.RequestContext) *h.Partial {
	locator := ctx.ServiceLocator()
	service := chat.NewService(locator)

	chatRoomId := ctx.FormValue("join-chat-room")

	if chatRoomId != "" {
		room, _ := service.GetRoom(chatRoomId)
		if room == nil {
			return h.SwapPartial(ctx, components.FormError("Room not found"))
		} else {
			return h.RedirectPartial("/chat/" + chatRoomId)
		}
	}

	chatRoomName := ctx.FormValue("chat-room-name")
	if chatRoomName != "" {
		// create room
	}

	return h.SwapPartial(ctx, components.FormError("Create a new room or join an existing one"))
}
