package ws

import "github.com/maddalax/htmgo/framework/h"

func OnClick(ctx *h.RequestContext, handler Handler) *h.AttributeMapOrdered {
	return AddClientSideHandler(ctx, "click", handler)
}

func OnClientEvent(ctx *h.RequestContext, eventName string, handler Handler) *h.AttributeMapOrdered {
	return AddClientSideHandler(ctx, eventName, handler)
}

func OnServerEvent(ctx *h.RequestContext, eventName string, handler Handler) h.Ren {
	AddServerSideHandler(ctx, eventName, handler)
	return h.Attribute("data-handler-id", "")
}

func OnMouseOver(ctx *h.RequestContext, handler Handler) *h.AttributeMapOrdered {
	return AddClientSideHandler(ctx, "mouseover", handler)
}
