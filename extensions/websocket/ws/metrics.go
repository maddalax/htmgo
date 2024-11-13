package ws

import (
	"github.com/maddalax/htmgo/extensions/websocket/internal/wsutil"
	"github.com/maddalax/htmgo/framework/h"
)

type Metrics struct {
	Manager wsutil.ManagerMetrics
	Handler HandlerMetrics
}

func MetricsFromCtx(ctx *h.RequestContext) Metrics {
	manager := ManagerFromCtx(ctx)
	return Metrics{
		Manager: manager.Metrics(),
		Handler: GetHandlerMetics(),
	}
}
