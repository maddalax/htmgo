package event

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/service"
	"sse-with-state/sse"
	"sse-with-state/state"
	"time"
)

func StartListener(locator *service.Locator) {
	manager := service.Get[sse.SocketManager](locator)
	manager.Listen(socketMessageListener)
	handler := NewMessageHandler(manager)

	go func() {
		for {
			fmt.Printf("total handlers: %d\n", handlers.Size())
			fmt.Printf("total serverEventNamesToHash: %d\n", serverEventNamesToHash.Size())
			fmt.Printf("total sessionIdToHashes: %d\n", sessionIdToHashes.Size())
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case event := <-serverSideMessageListener:
				handler.OnServerSideEvent(event)
			case event := <-socketMessageListener:
				switch event.Type {
				case sse.DisconnectedEvent:
					handler.OnSocketDisconnected(event)
				case sse.MessageEvent:
					handlerId := event.Payload["id"].(string)
					eventName := event.Payload["event"].(string)
					sessionId := state.SessionId(event.SessionId)
					if eventName == "dom-element-removed" {
						handler.OnDomElementRemoved(handlerId)
						continue
					} else {
						handler.OnClientSideEvent(handlerId, sessionId)
					}
				}
			}
		}
	}()
}
