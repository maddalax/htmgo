package event

import "github.com/maddalax/htmgo/framework/h"

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
func PushElement(data HandlerData, el *h.Element) {
	data.Manager.SendHtml(data.Socket.Id, h.Render(el))
}
