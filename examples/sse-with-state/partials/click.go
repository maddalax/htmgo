package partials

import (
	"github.com/maddalax/htmgo/framework/h"
	"sse-with-state/event"
)

func OnClick(ctx *h.RequestContext, handler event.Handler) *h.AttributeMapOrdered {
	return event.AddHandler(ctx, "click", handler)
}

func OnServerSideEvent(ctx *h.RequestContext, id string, eventName string, handler event.Handler) h.Ren {
	event.AddServerSideHandler(ctx, id, eventName, handler)
	return h.Empty()
}

func OnMouseOver(ctx *h.RequestContext, handler event.Handler) *h.AttributeMapOrdered {
	return event.AddHandler(ctx, "mouseover", handler)
}
