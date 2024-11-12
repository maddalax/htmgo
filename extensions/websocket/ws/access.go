package ws

import (
	"github.com/maddalax/htmgo/framework/h"
)

func ManagerFromCtx(ctx *h.RequestContext) *SocketManager {
	return SocketManagerFromCtx(ctx)
}
