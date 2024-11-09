package session

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/puzpuzpuz/xsync/v3"
)

type Id string

var cache = xsync.NewMapOf[Id, *xsync.MapOf[string, any]]()

type State struct {
	SessionId Id
}

func NewState(ctx *h.RequestContext) *State {
	id := GetSessionId(ctx)
	cache.Store(id, xsync.NewMapOf[string, any]())
	return &State{
		SessionId: id,
	}
}

func CreateSession(ctx *h.RequestContext) Id {
	sessionId := fmt.Sprintf("session-id-%s", h.GenId(30))
	ctx.Set("session-id", sessionId)
	return Id(sessionId)
}

func GetSessionId(ctx *h.RequestContext) Id {
	sessionIdRaw := ctx.Get("session-id")
	sessionId := ""

	if sessionIdRaw == "" || sessionIdRaw == nil {
		panic("session id is not set, please use session.CreateSession(ctx) in middleware to create a session id")
	} else {
		sessionId = sessionIdRaw.(string)
	}

	return Id(sessionId)
}

func Update[T any](sessionId Id, key string, compute func(prev T) T) T {
	actual := Get[T](sessionId, key, *new(T))
	next := compute(actual)
	Set(sessionId, key, next)
	return next
}

func Get[T any](sessionId Id, key string, fallback T) T {
	actual, _ := cache.LoadOrCompute(sessionId, func() *xsync.MapOf[string, any] {
		return xsync.NewMapOf[string, any]()
	})
	value, exists := actual.Load(key)
	if exists {
		return value.(T)
	}
	return fallback
}

func Set(sessionId Id, key string, value any) {
	actual, _ := cache.LoadOrCompute(sessionId, func() *xsync.MapOf[string, any] {
		return xsync.NewMapOf[string, any]()
	})
	actual.Store(key, value)
}

func UseState[T any](sessionId Id, key string, initial T) (func() T, func(T)) {
	var get = func() T {
		return Get[T](sessionId, key, initial)
	}
	var set = func(value T) {
		Set(sessionId, key, value)
	}
	return get, set
}
