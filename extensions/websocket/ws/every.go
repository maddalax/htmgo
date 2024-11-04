package ws

import (
	"github.com/maddalax/htmgo/extensions/websocket/internal/wsutil"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"github.com/maddalax/htmgo/framework/session"
	"log/slog"
	"time"
)

// Every executes the given callback every interval, until the socket is disconnected, or the callback returns false.
func Every(ctx *h.RequestContext, interval time.Duration, cb func() bool) {
	socketId := session.GetSessionId(ctx)
	socketIdSlog := slog.String("socketId", string(socketId))

	slog.Debug("ws-extension: starting every loop", socketIdSlog, slog.Duration("duration", interval))

	go func() {
		tries := 0
		for {
			locator := ctx.ServiceLocator()
			socketManager := service.Get[wsutil.SocketManager](locator)
			socket := socketManager.Get(string(socketId))
			// This can run before the socket is established, lets try a few times and kill it if socket isn't connected after a bit.
			if socket == nil {
				if tries > 5 {
					slog.Debug("ws-extension: socket disconnected, killing goroutine", socketIdSlog)
					return
				} else {
					time.Sleep(time.Second)
					tries++
					slog.Debug("ws-extension: socket not connected yet, trying again", socketIdSlog, slog.Int("attempt", tries))
					continue
				}
			}
			success := cb()
			if !success {
				return
			}
			time.Sleep(interval)
		}
	}()
}
