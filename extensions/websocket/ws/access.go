package ws

import (
	"github.com/maddalax/htmgo/extensions/websocket/internal/wsutil"
	"github.com/maddalax/htmgo/framework/h"
)

func ManagerFromCtx(ctx *h.RequestContext) *wsutil.SocketManager {
	return wsutil.SocketManagerFromCtx(ctx)
}
