package event

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"github.com/puzpuzpuz/xsync/v3"
	"runtime"
	"sse-with-state/internal"
	"sse-with-state/sse"
	"sse-with-state/state"
	"strings"
	"time"
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

type Metrics struct {
	Sessions []MetricPerSession
}

type MetricPerSession struct {
	SessionId       state.SessionId
	ClientListeners []MetricListener
	ServerListeners []MetricListener
}

type MetricListener struct {
	Event     string
	HandlerId string
}

func GetMetrics() *Metrics {
	metrics := &Metrics{
		Sessions: make([]MetricPerSession, 0),
	}

	Map.Range(func(key state.SessionId, value *Events) bool {
		clientListeners := make([]MetricListener, 0)
		value.client.Range(func(key string, value handlerWrapper) bool {
			clientListeners = append(clientListeners, MetricListener{
				Event:     "",
				HandlerId: key,
			})
			return true
		})

		serverListeners := make([]MetricListener, 0)
		value.server.Range(func(event string, value *xsync.MapOf[string, handlerWrapper]) bool {
			value.Range(func(handlerId string, value handlerWrapper) bool {
				serverListeners = append(serverListeners, MetricListener{
					Event:     event,
					HandlerId: handlerId,
				})
				return true
			})
			return true
		})

		metrics.Sessions = append(metrics.Sessions, MetricPerSession{
			SessionId:       key,
			ClientListeners: clientListeners,
			ServerListeners: serverListeners,
		})
		return true
	})

	return metrics
}

type Events struct {
	SessionId state.SessionId
	server    *xsync.MapOf[string, *xsync.MapOf[string, handlerWrapper]]
	client    *xsync.MapOf[string, handlerWrapper]
}

func NewEvents(sessionId state.SessionId) *Events {
	return &Events{
		SessionId: sessionId,
		server:    xsync.NewMapOf[string, *xsync.MapOf[string, handlerWrapper]](),
		client:    xsync.NewMapOf[string, handlerWrapper](),
	}
}

func (e *Events) AddServerSideHandler(event string, id string, handler Handler) {
	wrapper := handlerWrapper{
		handler:   handler,
		sessionId: e.SessionId,
	}

	if _, ok := e.server.Load(event); !ok {
		e.server.Store(event, xsync.NewMapOf[string, handlerWrapper]())
	}

	handlers, _ := e.server.Load(event)
	handlers.Store(id, wrapper)
	e.server.Store(event, handlers)
}

func (e *Events) AddClientSideHandler(event string, handler Handler) *h.AttributeMapOrdered {
	handlerId := fmt.Sprintf("event_%s_%s", event, internal.RandSeq(30))
	fmt.Printf("adding client side handler %s\n", handlerId)
	e.client.Store(handlerId, handlerWrapper{
		handler:   handler,
		sessionId: e.SessionId,
	})
	return h.AttributePairs(
		"data-handler-id", handlerId,
		"data-handler-event", event,
	)
}

func (e *Events) OnServerSideEvent(manager *sse.SocketManager, eventName string) {
	handlers, ok := e.server.Load(eventName)
	if ok {
		socket := manager.Get(string(e.SessionId))
		if socket == nil {
			fmt.Printf("socket not found, must be disconnected: %s", e.SessionId)
			e.OnSocketDisconnected()
			Map.Delete(e.SessionId)
			return
		}
		handlers.Range(func(key string, value handlerWrapper) bool {
			go value.handler(HandlerData{
				SessionId: e.SessionId,
				Socket:    socket,
				Manager:   manager,
			})
			return true
		})
	}
}

func (e *Events) OnClientSideEvent(manager *sse.SocketManager, handlerId string) {
	handlers, ok := e.client.Load(handlerId)
	if ok {
		go handlers.handler(HandlerData{
			SessionId: e.SessionId,
			Socket:    manager.Get(string(e.SessionId)),
			Manager:   manager,
		})
	}
}

func (e *Events) OnDomElementRemoved(id string) {
	e.server.Range(func(key string, value *xsync.MapOf[string, handlerWrapper]) bool {
		value.Delete(id)
		return true
	})
	e.client.Delete(id)
}

func (e *Events) OnSocketDisconnected() {
	e.client.Clear()
	e.server.Clear()
}

var Map = xsync.NewMapOf[state.SessionId, *Events]()

var socketMessageListener = make(chan sse.SocketEvent, 100)
var serverSideMessageListener = make(chan ServerSideEvent, 100)

func getCallerHash() string {
	now := time.Now()
	pc := make([]uintptr, 1000) // Adjust the size if you need a deeper stack trace
	n := runtime.Callers(2, pc) // Skip 2: runtime.Callers() and printCallers()

	frames := runtime.CallersFrames(pc[:n])
	calls := make([]string, 0)

	for {
		frame, more := frames.Next()
		calls = append(calls, fmt.Sprintf("%s:%d", frame.Function, frame.Line))
		if !more {
			break
		}
	}

	took := time.Since(now)
	fmt.Printf("took: %dms\n", took.Milliseconds())
	return strings.Join(calls, ",")
}

func AddServerSideHandler(ctx *h.RequestContext, event string, handler Handler) {
	hash := getCallerHash()
	fmt.Printf("adding server side handler %s\n", hash)
	sessionId := state.GetSessionId(ctx)
	events, ok := Map.Load(sessionId)

	if !ok {
		events = NewEvents(sessionId)
		Map.Store(sessionId, events)
	}

	events.AddServerSideHandler(event, "test", handler)
}

func AddHandler(ctx *h.RequestContext, event string, handler Handler) *h.AttributeMapOrdered {
	sessionId := state.GetSessionId(ctx)
	events, ok := Map.Load(sessionId)

	if !ok {
		events = NewEvents(sessionId)
		Map.Store(sessionId, events)
	}

	return events.AddClientSideHandler(event, handler)
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
				Map.Range(func(key state.SessionId, value *Events) bool {
					value.OnServerSideEvent(
						manager,
						sevent.Event,
					)
					return true
				})

			case event := <-socketMessageListener:
				if event.Type == sse.DisconnectedEvent {
					sessionId := state.SessionId(event.SessionId)
					handler, ok := Map.Load(sessionId)
					if ok {
						handler.OnSocketDisconnected()
						Map.Delete(sessionId)
					}
					continue
				}

				if event.Type == sse.MessageEvent {
					handlerId := event.Payload["id"].(string)
					eventName := event.Payload["event"].(string)
					sessionId := state.SessionId(event.SessionId)

					fmt.Printf("received eventName: %s, handlerId: %s, sessionId: %s\n", eventName, handlerId, sessionId)

					handler, ok := Map.Load(sessionId)

					if !ok {
						return
					}

					if eventName == "dom-element-removed" {
						handler.OnDomElementRemoved(handlerId)
						continue
					}

					handler.OnClientSideEvent(manager, handlerId)
				}
			}
		}
	}()
}
