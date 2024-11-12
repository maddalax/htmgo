package ws

import (
	"github.com/maddalax/htmgo/framework/h"
)

type Metrics struct {
	Manager ManagerMetrics
	Handler HandlerMetrics
}

func MetricsFromCtx(ctx *h.RequestContext) Metrics {
	manager := ManagerFromCtx(ctx)
	return Metrics{
		Manager: manager.Metrics(),
		Handler: GetHandlerMetics(),
	}
}
