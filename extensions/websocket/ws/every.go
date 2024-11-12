package ws

import (
	"github.com/maddalax/htmgo/extensions/websocket/session"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"time"
)

// Every executes the given callback every interval, until the socket is disconnected, or the callback returns false.
func Every(ctx *h.RequestContext, interval time.Duration, cb func() bool) {
	socketId := session.GetSessionId(ctx)
	locator := ctx.ServiceLocator()
	manager := service.Get[SocketManager](locator)
	manager.RunIntervalWithSocket(string(socketId), interval, cb)
}

func Once(ctx *h.RequestContext, cb func()) {
	// time is irrelevant, we just need to run the callback once, it will exit after because of the false return
	Every(ctx, time.Millisecond, func() bool {
		cb()
		return false
	})
}

func RunOnConnected(ctx *h.RequestContext, cb func()) {
	Once(ctx, cb)
}
