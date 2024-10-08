package pages

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"sse-with-state/partials"
	"sse-with-state/state"
)

func IndexPage(ctx *h.RequestContext) *h.Page {
	state.NewState(ctx)
	return h.NewPage(
		RootPage(
			h.Div(
				h.Attribute("ws-connect", fmt.Sprintf("/ws/test")),
				h.Class("flex flex-col gap-4 items-center pt-24 min-h-screen bg-neutral-100"),
				h.H3(h.Id("intro-text"), h.Text("Repeater Example"), h.Class("text-2xl")),

				partials.CounterForm(ctx, partials.CounterProps{Id: "counter-1"}),

				partials.Repeater(ctx, partials.RepeaterProps{
					Id: "repeater-1",
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
						return h.Input("text",
							h.Class("border border-gray-300 rounded p-2"),
							h.Value(fmt.Sprintf("item %d", index)))
					},
				}),
			),
		),
	)
}
