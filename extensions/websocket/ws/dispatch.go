package ws

import (
	"github.com/maddalax/htmgo/extensions/websocket/internal/wsutil"
	"github.com/maddalax/htmgo/extensions/websocket/session"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
)

// PushServerSideEvent sends a server side event this specific session
func PushServerSideEvent(data HandlerData, event string, value map[string]any) {
	serverSideMessageListener <- ServerSideEvent{
		Event:     event,
		Payload:   value,
		SessionId: data.SessionId,
	}
}

// BroadcastServerSideEvent sends a server side event to all clients that have a handler for the event, not just the current session
func BroadcastServerSideEvent(event string, value map[string]any) {
	serverSideMessageListener <- ServerSideEvent{
		Event:     event,
		Payload:   value,
		SessionId: "*",
	}
}

// PushElement sends an element to the current session and swaps it into the page
func PushElement(data HandlerData, el *h.Element) bool {
	return data.Manager.SendHtml(data.Socket.Id, h.Render(el))
}

// PushElementCtx sends an element to the current session and swaps it into the page
func PushElementCtx(ctx *h.RequestContext, el *h.Element) bool {
	locator := ctx.ServiceLocator()
	socketManager := service.Get[wsutil.SocketManager](locator)
	socketId := session.GetSessionId(ctx)
	socket := socketManager.Get(string(socketId))
	if socket == nil {
		return false
	}
	return PushElement(HandlerData{
		Socket:    socket,
		Manager:   socketManager,
		SessionId: socketId,
	}, el)
}
