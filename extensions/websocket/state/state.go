package state

import (
	"fmt"
	"github.com/maddalax/htmgo/extensions/websocket/internal"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/puzpuzpuz/xsync/v3"
)

type SessionId string

var cache = xsync.NewMapOf[SessionId, *xsync.MapOf[string, any]]()

type State struct {
	SessionId SessionId
}

func NewState(ctx *h.RequestContext) *State {
	id := GetSessionId(ctx)
	cache.Store(id, xsync.NewMapOf[string, any]())
	return &State{
		SessionId: id,
	}
}

func GetSessionId(ctx *h.RequestContext) SessionId {
	sessionIdRaw := ctx.Get("session-id")
	sessionId := ""

	if sessionIdRaw == "" || sessionIdRaw == nil {
		sessionId = fmt.Sprintf("session-id-%s", internal.RandSeq(30))
		ctx.Set("session-id", sessionId)
	} else {
		sessionId = sessionIdRaw.(string)
	}

	return SessionId(sessionId)
}

func Update[T any](sessionId SessionId, key string, compute func(prev T) T) T {
	actual := Get[T](sessionId, key, *new(T))
	next := compute(actual)
	Set(sessionId, key, next)
	return next
}

func Get[T any](sessionId SessionId, key string, fallback T) T {
	actual, _ := cache.LoadOrCompute(sessionId, func() *xsync.MapOf[string, any] {
		return xsync.NewMapOf[string, any]()
	})
	value, exists := actual.Load(key)
	if exists {
		return value.(T)
	}
	return fallback
}

func Set(sessionId SessionId, key string, value any) {
	actual, _ := cache.LoadOrCompute(sessionId, func() *xsync.MapOf[string, any] {
		return xsync.NewMapOf[string, any]()
	})
	actual.Store(key, value)
}

func Use[T any](sessionId SessionId, key string, initial T) (func() T, func(T)) {
	var get = func() T {
		return Get[T](sessionId, key, initial)
	}
	var set = func(value T) {
		Set(sessionId, key, value)
	}
	return get, set
}
