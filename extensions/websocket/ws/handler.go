package ws

import (
	"fmt"
	"github.com/maddalax/htmgo/extensions/websocket/internal/wsutil"
	"github.com/maddalax/htmgo/extensions/websocket/session"
	"sync"
)

type MessageHandler struct {
	manager *wsutil.SocketManager
}

func NewMessageHandler(manager *wsutil.SocketManager) *MessageHandler {
	return &MessageHandler{manager: manager}
}

func (h *MessageHandler) OnServerSideEvent(e ServerSideEvent) {
	fmt.Printf("received server side event: %s\n", e.Event)
	hashes, ok := serverEventNamesToHash.Load(e.Event)

	// If we are not broadcasting to everyone, filter it down to just the current session that invoked the event
	// TODO optimize this
	if e.SessionId != "*" {
		hashesForSession, ok2 := sessionIdToHashes.Load(e.SessionId)

		if ok2 {
			subset := make(map[KeyHash]bool)
			for hash := range hashes {
				if _, ok := hashesForSession[hash]; ok {
					subset[hash] = true
				}
			}
			hashes = subset
		}
	}

	if ok {
		lock.Lock()
		callingHandler.Store(true)
		wg := sync.WaitGroup{}
		for hash := range hashes {
			cb, ok := handlers.Load(hash)
			if ok {
				wg.Add(1)
				go func(e ServerSideEvent) {
					defer wg.Done()
					sessionId, ok2 := hashesToSessionId.Load(hash)
					if ok2 {
						cb(HandlerData{
							SessionId: sessionId,
							Socket:    h.manager.Get(string(sessionId)),
							Manager:   h.manager,
						})
					}
				}(e)
			}
		}
		wg.Wait()
		callingHandler.Store(false)
		lock.Unlock()
	}
}

func (h *MessageHandler) OnClientSideEvent(handlerId string, sessionId session.Id) {
	cb, ok := handlers.Load(handlerId)
	if ok {
		cb(HandlerData{
			SessionId: sessionId,
			Socket:    h.manager.Get(string(sessionId)),
			Manager:   h.manager,
		})
	}
}

func (h *MessageHandler) OnDomElementRemoved(handlerId string) {
	handlers.Delete(handlerId)
}

func (h *MessageHandler) OnSocketDisconnected(event wsutil.SocketEvent) {
	sessionId := session.Id(event.SessionId)
	hashes, ok := sessionIdToHashes.Load(sessionId)
	if ok {
		for hash := range hashes {
			hashesToSessionId.Delete(hash)
			handlers.Delete(hash)
		}
		sessionIdToHashes.Delete(sessionId)
	}
}
