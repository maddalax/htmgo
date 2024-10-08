package event

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"github.com/puzpuzpuz/xsync/v3"
	"sse-with-state/internal"
	"sse-with-state/sse"
	"sse-with-state/state"
)

type HandlerData struct {
	SessionId state.SessionId
	Socket    *sse.SocketConnection
	Manager   *sse.SocketManager
}

type Handler func(data HandlerData)

type handlerWrapper struct {
	handler   Handler
	sessionId state.SessionId
}

type ServerSideEvent struct {
	Event     string
	Payload   map[string]any
	SessionId state.SessionId
}

var Map = xsync.NewMapOf[string, handlerWrapper]()
var ServerSideEventMap = xsync.NewMapOf[string, *xsync.MapOf[string, handlerWrapper]]()
var socketMessageListener = make(chan sse.SocketEvent, 100)
var serverSideMessageListener = make(chan ServerSideEvent, 100)

func AddServerSideHandler(ctx *h.RequestContext, id string, event string, handler Handler) {
	sessionId := state.GetSessionId(ctx)

	wrapper := handlerWrapper{
		handler:   handler,
		sessionId: sessionId,
	}

	handlers, ok := ServerSideEventMap.Load(event)
	if !ok {
		ServerSideEventMap.Store(event, xsync.NewMapOf[string, handlerWrapper]())
	}

	handlers, _ = ServerSideEventMap.Load(event)
	handlers.Store(id, wrapper)

	fmt.Printf("added server side handler for %s, %v\n", event, handlers)
}

func AddHandler(ctx *h.RequestContext, event string, handler Handler) *h.AttributeMapOrdered {
	handlerId := fmt.Sprintf("event_%s_%s", event, internal.RandSeq(30))
	for {
		_, ok := Map.Load(handlerId)
		if ok {
			handlerId = fmt.Sprintf("event_%s_%s", event, internal.RandSeq(30))
		} else {
			break
		}
	}
	sessionId := state.GetSessionId(ctx)
	Map.Store(handlerId, handlerWrapper{
		handler:   handler,
		sessionId: sessionId,
	})
	return h.AttributePairs(
		"data-handler-id", handlerId,
		"data-handler-event", event,
	)
}

func PushServerSideEvent(sessionId state.SessionId, event string) {
	serverSideMessageListener <- ServerSideEvent{
		Event:     event,
		Payload:   make(map[string]any),
		SessionId: sessionId,
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
			select {
			case sevent := <-serverSideMessageListener:
				handlers, ok := ServerSideEventMap.Load(sevent.Event)
				if ok {
					handlers.Range(func(key string, value handlerWrapper) bool {
						go value.handler(HandlerData{
							SessionId: sevent.SessionId,
							Socket:    manager.Get(string(sevent.SessionId)),
							Manager:   manager,
						})
						return true
					})
				}
			case event := <-socketMessageListener:
				if event.Type == sse.MessageEvent {
					handlerId := event.Payload["id"].(string)
					eventName := event.Payload["event"].(string)
					cb, ok := Map.Load(handlerId)
					if ok {
						fmt.Printf("calling %s handler for session: %s\n", eventName, cb.sessionId)
						go cb.handler(HandlerData{
							SessionId: cb.sessionId,
							Socket:    manager.Get(event.SocketId),
							Manager:   manager,
						})
					}
				}
			}
		}
	}()
}
