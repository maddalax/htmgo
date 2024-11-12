package ws

import (
	"github.com/maddalax/htmgo/extensions/websocket/session"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/puzpuzpuz/xsync/v3"
	"sync"
	"sync/atomic"
)

type HandlerData struct {
	SessionId session.Id
	Socket    *SocketConnection
	Manager   *SocketManager
}

type Handler func(data HandlerData)

type ServerSideEvent struct {
	Event     string
	Payload   map[string]any
	SessionId session.Id
}
type KeyHash = string

var handlers = xsync.NewMapOf[KeyHash, Handler]()
var sessionIdToHashes = xsync.NewMapOf[session.Id, map[KeyHash]bool]()
var hashesToSessionId = xsync.NewMapOf[KeyHash, session.Id]()
var serverEventNamesToHash = xsync.NewMapOf[string, map[KeyHash]bool]()

var socketMessageListener = make(chan SocketEvent, 100)
var serverSideMessageListener = make(chan ServerSideEvent, 100)
var lock = sync.Mutex{}
var callingHandler = atomic.Bool{}

type HandlerMetrics struct {
	TotalHandlers               int
	ServerEventNamesToHashCount int
	SessionIdToHashesCount      int
}

func GetHandlerMetics() HandlerMetrics {
	metrics := HandlerMetrics{
		TotalHandlers:               handlers.Size(),
		ServerEventNamesToHashCount: serverEventNamesToHash.Size(),
		SessionIdToHashesCount:      sessionIdToHashes.Size(),
	}
	return metrics
}

func makeId() string {
	return h.GenId(30)
}

func AddServerSideHandler(ctx *h.RequestContext, event string, handler Handler) *h.AttributeMapOrdered {
	// If we are already in a handler, we don't want to add another handler
	// this can happen if the handler renders another element that has a handler
	if callingHandler.Load() {
		return h.NewAttributeMap()
	}
	sessionId := session.GetSessionId(ctx)
	hash := makeId()
	handlers.LoadOrStore(hash, handler)
	m, _ := serverEventNamesToHash.LoadOrCompute(event, func() map[KeyHash]bool {
		return make(map[KeyHash]bool)
	})
	m[hash] = true
	storeHashForSession(sessionId, hash)
	storeSessionIdForHash(sessionId, hash)
	return h.AttributePairs("data-handler-id", hash, "data-handler-event", event)
}

func AddClientSideHandler(ctx *h.RequestContext, event string, handler Handler) *h.AttributeMapOrdered {
	hash := makeId()
	handlers.LoadOrStore(hash, handler)
	sessionId := session.GetSessionId(ctx)
	storeHashForSession(sessionId, hash)
	storeSessionIdForHash(sessionId, hash)
	return h.AttributePairs("data-handler-id", hash, "data-handler-event", event)
}

func storeHashForSession(sessionId session.Id, hash KeyHash) {
	m, _ := sessionIdToHashes.LoadOrCompute(sessionId, func() map[KeyHash]bool {
		return make(map[KeyHash]bool)
	})
	m[hash] = true
}

func storeSessionIdForHash(sessionId session.Id, hash KeyHash) {
	hashesToSessionId.Store(hash, sessionId)
}
