package partials

import (
	"chat/chat"
	"chat/components"
	"github.com/maddalax/htmgo/framework/h"
	"net/http"
	"time"
)

func CreateOrJoinRoom(ctx *h.RequestContext) *h.Partial {
	locator := ctx.ServiceLocator()
	service := chat.NewService(locator)

	chatRoomId := ctx.Request.FormValue("join-chat-room")
	username := ctx.Request.FormValue("username")

	if username == "" {
		return h.SwapPartial(ctx, components.FormError("Username is required"))
	}

	user, err := service.CreateUser(username)

	if err != nil {
		return h.SwapPartial(ctx, components.FormError("Failed to create user"))
	}

	var redirect = func(path string) *h.Partial {
		cookie := &http.Cookie{
			Name:    "session_id",
			Value:   user.SessionID,
			Path:    "/",
			Expires: time.Now().Add(24 * 30 * time.Hour),
		}
		return h.SwapManyPartialWithHeaders(
			ctx,
			h.NewHeaders(
				"Set-Cookie", cookie.String(),
				"HX-Redirect", path,
			),
			h.Fragment(),
		)
	}

	if chatRoomId != "" {
		room, _ := service.GetRoom(chatRoomId)
		if room == nil {
			return h.SwapPartial(ctx, components.FormError("Room not found"))
		} else {
			return redirect("/chat/" + chatRoomId)
		}
	}

	chatRoomName := ctx.Request.FormValue("new-chat-room")
	if chatRoomName != "" {
		room, _ := service.CreateRoom(chatRoomName)
		if room == nil {
			return h.SwapPartial(ctx, components.FormError("Failed to create room"))
		} else {
			return redirect("/chat/" + room.ID)
		}
	}

	return h.SwapPartial(ctx, components.FormError("Create a new room or join an existing one"))
}
