package partials

import (
	"github.com/maddalax/htmgo/framework/h"
	"sse-with-state/event"
	"sse-with-state/state"
)

func UseState[T any](sessionId state.SessionId, key string, initial T) (func() T, func(T)) {
	var get = func() T {
		return state.Get[T](sessionId, key, initial)
	}
	var set = func(value T) {
		state.Set(sessionId, key, value)
	}
	return get, set
}

type Counter struct {
	Count     func() int
	Increment func()
}

func UseCounter(sessionId state.SessionId, id string) Counter {
	get, set := UseState(sessionId, id, 0)

	var increment = func() {
		set(get() + 1)
	}

	return Counter{
		Count:     get,
		Increment: increment,
	}
}

type CounterProps struct {
	Id string
}

func CounterForm(ctx *h.RequestContext, props CounterProps) *h.Element {
	if props.Id == "" {
		props.Id = h.GenId()
	}
	counter := UseCounter(state.GetSessionId(ctx), props.Id)
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
			OnServerSideEvent(ctx, props.Id, "increment", func(data event.HandlerData) {
				counter.Increment()
				event.PushElement(data, CounterForm(ctx, props))
			}),
			//OnMouseOver(ctx, func(data event.HandlerData) {
			//	counter.Increment()
			//	updated := CounterForm(ctx, props)
			//	event.PushElement(data, updated)
			//}),
			//OnClick(ctx, func(data event.HandlerData) {
			//	counter.Increment()
			//	updated := CounterForm(ctx, props)
			//	event.PushElement(data, updated)
			//}),
		),
	)
}
