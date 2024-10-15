package partials

import (
	"github.com/maddalax/htmgo/framework/h"
	"sse-with-state/event"
)

func OnClick(ctx *h.RequestContext, handler event.Handler) *h.AttributeMapOrdered {
	return event.AddHandler(ctx, "click", handler)
}

func OnServerSideEvent(ctx *h.RequestContext, eventName string, handler event.Handler) h.Ren {
	event.AddServerSideHandler(ctx, eventName, handler)
	return h.Attribute("data-handler-id", "")
}

func OnMouseOver(ctx *h.RequestContext, handler event.Handler) *h.AttributeMapOrdered {
	return event.AddHandler(ctx, "mouseover", handler)
}
