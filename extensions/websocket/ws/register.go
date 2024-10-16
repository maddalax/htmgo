package ws

import (
	"github.com/maddalax/htmgo/extensions/websocket/internal/wsutil"
	"github.com/maddalax/htmgo/extensions/websocket/state"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/puzpuzpuz/xsync/v3"
	"sync"
	"sync/atomic"
)

type HandlerData struct {
	SessionId state.SessionId
	Socket    *wsutil.SocketConnection
	Manager   *wsutil.SocketManager
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
var hashesToSessionId = xsync.NewMapOf[KeyHash, state.SessionId]()
var serverEventNamesToHash = xsync.NewMapOf[string, map[KeyHash]bool]()

var socketMessageListener = make(chan wsutil.SocketEvent, 100)
var serverSideMessageListener = make(chan ServerSideEvent, 100)
var lock = sync.Mutex{}
var callingHandler = atomic.Bool{}

func makeId() string {
	return h.GenId(30)
}

func AddServerSideHandler(ctx *h.RequestContext, event string, handler Handler) *h.AttributeMapOrdered {
	// If we are already in a handler, we don't want to add another handler
	// this can happen if the handler renders another element that has a handler
	if callingHandler.Load() {
		return h.NewAttributeMap()
	}
	sessionId := state.GetSessionId(ctx)
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
	sessionId := state.GetSessionId(ctx)
	storeHashForSession(sessionId, hash)
	storeSessionIdForHash(sessionId, hash)
	return h.AttributePairs("data-handler-id", hash, "data-handler-event", event)
}

func storeHashForSession(sessionId state.SessionId, hash KeyHash) {
	m, _ := sessionIdToHashes.LoadOrCompute(sessionId, func() map[KeyHash]bool {
		return make(map[KeyHash]bool)
	})
	m[hash] = true
}

func storeSessionIdForHash(sessionId state.SessionId, hash KeyHash) {
	hashesToSessionId.Store(hash, sessionId)
}
