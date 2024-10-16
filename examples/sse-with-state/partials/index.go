package partials

import (
	"github.com/maddalax/htmgo/extensions/websocket/ws"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/session"
)

type Counter struct {
	Count     func() int
	Increment func()
	Decrement func()
}

func UseCounter(sessionId session.Id, id string) Counter {
	get, set := session.UseState(sessionId, id, 0)

	var increment = func() {
		set(get() + 1)
	}

	var decrement = func() {
		set(get() - 1)
	}

	return Counter{
		Count:     get,
		Increment: increment,
		Decrement: decrement,
	}
}

type CounterProps struct {
	Id string
}

func CounterForm(ctx *h.RequestContext, props CounterProps) *h.Element {
	if props.Id == "" {
		props.Id = h.GenId(6)
	}
	counter := UseCounter(session.GetSessionId(ctx), props.Id)

	return h.Div(
		h.Attribute("hx-swap", "none"),
		h.Class("flex flex-col gap-3 items-center"),
		h.Id(props.Id),
		h.P(
			h.Id("counter-text-"+props.Id),
			h.AttributePairs(
				"id", "counter",
				"class", "text-xl",
				"name", "count",
				"text", "count",
			),
			h.TextF("Count: %d", counter.Count()),
		),
		h.Button(
			h.Class("bg-rose-400 hover:bg-rose-500 text-white font-bold py-2 px-4 rounded"),
			h.Type("submit"),
			h.Text("Increment"),
			ws.OnServerSideEvent(ctx, "increment", func(data ws.HandlerData) {
				counter.Increment()
				ws.PushElement(data, CounterForm(ctx, props))
			}),
			ws.OnServerSideEvent(ctx, "decrement", func(data ws.HandlerData) {
				counter.Decrement()
				ws.PushElement(data, CounterForm(ctx, props))
			}),
		),
	)
}
