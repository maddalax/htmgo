package pages

import (
	"fmt"
	"github.com/maddalax/htmgo/extensions/websocket/session"
	"github.com/maddalax/htmgo/extensions/websocket/ws"
	"github.com/maddalax/htmgo/framework/h"
	"ws-example/partials"
)

func IndexPage(ctx *h.RequestContext) *h.Page {
	sessionId := session.GetSessionId(ctx)

	return h.NewPage(
		RootPage(
			ctx,
			h.Div(
				h.Attribute("ws-connect", fmt.Sprintf("/ws?sessionId=%s", sessionId)),
				h.Class("flex flex-col gap-4 items-center pt-24 min-h-screen bg-neutral-100"),
				h.H3(
					h.Id("intro-text"),
					h.Text("Repeater Example"),
					h.Class("text-2xl"),
				),
				h.Div(
					h.Id("ws-metrics"),
				),
				partials.CounterForm(ctx, partials.CounterProps{Id: "counter-1"}),
				partials.Repeater(ctx, partials.RepeaterProps{
					Id: "repeater-1",
					OnAdd: func(data ws.HandlerData) {
						//ws.BroadcastServerSideEvent("increment", map[string]any{})
					},
					OnRemove: func(data ws.HandlerData, index int) {
						//ws.BroadcastServerSideEvent("decrement", map[string]any{})
					},
					AddButton: h.Button(
						h.Text("+ Add Item"),
					),
					RemoveButton: func(index int, children ...h.Ren) *h.Element {
						return h.Button(
							h.Text("Remove"),
							h.Children(children...),
						)
					},
					Item: func(index int) *h.Element {
						return h.Input(
							"text",
							h.Class("border border-gray-300 rounded p-2"),
							h.Value(fmt.Sprintf("item %d", index)),
						)
					},
				}),
			),
		),
	)
}
