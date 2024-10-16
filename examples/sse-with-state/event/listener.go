package event

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"github.com/puzpuzpuz/xsync/v3"
	"sse-with-state/sse"
	"sse-with-state/state"
	"sync"
	"sync/atomic"
	"time"
)

type HandlerData struct {
	SessionId state.SessionId
	Socket    *sse.SocketConnection
	Manager   *sse.SocketManager
}

type Handler func(data HandlerData)

type ServerSideEvent struct {
	Event     string
	Payload   map[string]any
	SessionId state.SessionId
}

type KeyHash = string

var handlers = xsync.NewMapOf[KeyHash, Handler]()
var sessionIdToHashes = xsync.NewMapOf[state.SessionId, map[KeyHash]bool]()
var serverEventNamesToHash = xsync.NewMapOf[string, map[KeyHash]bool]()

var socketMessageListener = make(chan sse.SocketEvent, 100)
var serverSideMessageListener = make(chan ServerSideEvent, 100)
var lock = sync.Mutex{}
var callingHandler = atomic.Bool{}

func AddServerSideHandler(ctx *h.RequestContext, event string, handler Handler) *h.AttributeMapOrdered {
	// If we are already in a handler, we don't want to add another handler
	// this can happen if the handler renders another element that has a handler
	if callingHandler.Load() {
		return h.NewAttributeMap()
	}
	sessionId := state.GetSessionId(ctx)
	hash := uuid.NewString()
	fmt.Printf("adding server side handler %s\n", hash)
	handlers.LoadOrStore(hash, handler)
	m, _ := serverEventNamesToHash.LoadOrCompute(event, func() map[KeyHash]bool {
		return make(map[KeyHash]bool)
	})
	m[hash] = true
	storeHashForSession(sessionId, hash)
	return h.AttributePairs("data-handler-id", hash, "data-handler-event", event)
}

func AddClientSideHandler(ctx *h.RequestContext, event string, handler Handler) *h.AttributeMapOrdered {
	hash := uuid.NewString()
	fmt.Printf("adding client side handler %s\n", hash)
	handlers.LoadOrStore(hash, handler)
	sessionId := state.GetSessionId(ctx)
	storeHashForSession(sessionId, hash)
	return h.AttributePairs("data-handler-id", hash, "data-handler-event", event)
}

func storeHashForSession(sessionId state.SessionId, hash KeyHash) {
	m, _ := sessionIdToHashes.LoadOrCompute(sessionId, func() map[KeyHash]bool {
		return make(map[KeyHash]bool)
	})
	m[hash] = true
}

func PushServerSideEvent(data HandlerData, event string, value map[string]any) {
	serverSideMessageListener <- ServerSideEvent{
		Event:     event,
		Payload:   value,
		SessionId: data.SessionId,
	}
}

func PushElement(data HandlerData, el *h.Element) {
	data.Manager.SendHtml(data.Socket.Id, h.Render(el))
}

func StartListener(locator *service.Locator) {
	manager := service.Get[sse.SocketManager](locator)
	manager.Listen(socketMessageListener)

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
			case sevent := <-serverSideMessageListener:
				fmt.Printf("received server side event: %s\n", sevent.Event)
				// TODO optimize this
				hashes, ok := serverEventNamesToHash.Load(sevent.Event)
				hashesForSession, ok2 := sessionIdToHashes.Load(sevent.SessionId)

				subset := make(map[KeyHash]bool)
				for hash := range hashes {
					if _, ok := hashesForSession[hash]; ok {
						subset[hash] = true
					}
				}

				if ok && ok2 {
					lock.Lock()
					callingHandler.Store(true)
					wg := sync.WaitGroup{}
					for hash := range subset {
						cb, ok := handlers.Load(hash)
						if ok {
							wg.Add(1)
							go func(e ServerSideEvent) {
								defer wg.Done()
								cb(HandlerData{
									SessionId: e.SessionId,
									Socket:    manager.Get(string(e.SessionId)),
									Manager:   manager,
								})
							}(sevent)
						}
					}
					wg.Wait()
					callingHandler.Store(false)
					lock.Unlock()
				}

			case event := <-socketMessageListener:

				if event.Type == sse.DisconnectedEvent {
					sessionId := state.SessionId(event.SessionId)
					fmt.Printf("disconnected sessionId: %s\n", sessionId)
					hashes, ok := sessionIdToHashes.Load(sessionId)
					if ok {
						for hash := range hashes {
							handlers.Delete(hash)
						}
						sessionIdToHashes.Delete(sessionId)
					}
				}

				if event.Type == sse.MessageEvent {
					handlerId := event.Payload["id"].(string)
					eventName := event.Payload["event"].(string)
					sessionId := state.SessionId(event.SessionId)
					fmt.Printf("received eventName: %s, handlerId: %s, sessionId: %s\n", eventName, handlerId, sessionId)

					if eventName == "dom-element-removed" {
						handlers.Delete(handlerId)
						continue
					}

					cb, ok := handlers.Load(handlerId)
					if ok {
						cb(HandlerData{
							SessionId: sessionId,
							Socket:    manager.Get(event.SessionId),
							Manager:   manager,
						})
					}
				}
			}
		}
	}()
}
