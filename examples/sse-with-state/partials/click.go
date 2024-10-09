package partials

import (
	"github.com/maddalax/htmgo/framework/h"
	"sse-with-state/event"
	"sse-with-state/internal"
)

func OnClick(ctx *h.RequestContext, handler event.Handler) *h.AttributeMapOrdered {
	return event.AddHandler(ctx, "click", handler)
}

func OnServerSideEvent(ctx *h.RequestContext, eventName string, handler event.Handler) h.Ren {
	id := internal.RandSeq(30)
	event.AddServerSideHandler(ctx, id, eventName, handler)
	return h.Attribute("data-handler-id", id)
}

func OnMouseOver(ctx *h.RequestContext, handler event.Handler) *h.AttributeMapOrdered {
	return event.AddHandler(ctx, "mouseover", handler)
}
