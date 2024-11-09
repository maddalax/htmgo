package ws

import (
	"github.com/maddalax/htmgo/extensions/websocket/internal/wsutil"
	"github.com/maddalax/htmgo/extensions/websocket/session"
	"github.com/maddalax/htmgo/framework/service"
)

func StartListener(locator *service.Locator) {
	manager := service.Get[wsutil.SocketManager](locator)
	manager.Listen(socketMessageListener)
	handler := NewMessageHandler(manager)
	go func() {
		for {
			handle(handler)
		}
	}()
}

func handle(handler *MessageHandler) {
	select {
	case event := <-serverSideMessageListener:
		handler.OnServerSideEvent(event)
	case event := <-socketMessageListener:
		switch event.Type {
		case wsutil.DisconnectedEvent:
			handler.OnSocketDisconnected(event)
		case wsutil.MessageEvent:

			handlerId, ok := event.Payload["id"].(string)
			eventName, ok2 := event.Payload["event"].(string)

			if !ok || !ok2 {
				return
			}

			sessionId := session.Id(event.SessionId)
			if eventName == "dom-element-removed" {
				handler.OnDomElementRemoved(handlerId)
				return
			} else {
				handler.OnClientSideEvent(handlerId, sessionId)
			}
		}
	}
}
