package state

import (
	"github.com/google/uuid"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/puzpuzpuz/xsync/v3"
	"net/http"
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
	stateCookie, err := ctx.Request.Cookie("state")
	sessionId := ""
	if err == nil {
		sessionId = stateCookie.Value
	} else {
		sessionId = uuid.NewString()
	}

	c := http.Cookie{
		Name:  "state",
		Value: sessionId,
	}
	ctx.Response.Header().Set("Set-Cookie", c.String())

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
