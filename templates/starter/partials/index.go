package partials

import (
	"github.com/maddalax/htmgo/framework/h"
	"strconv"
)

func CounterPartial(ctx *h.RequestContext) *h.Partial {
	count, err := strconv.ParseInt(ctx.FormValue("count"), 10, 64)

	if err != nil {
		count = 0
	}

	count++

	return h.SwapManyPartial(
		ctx,
		CounterForm(int(count)),
		h.ElementIf(count > 10, SubmitButton("New record!")),
	)
}

func CounterForm(count int) *h.Element {
	return h.Form(
		h.Class("flex flex-col gap-3 items-center"),
		h.Id("counter-form"),
		h.PostPartial(CounterPartial),
		h.Input(
			"text",
			h.Class("hidden"),
			h.Value(count),
			h.Name("count"),
		),
		h.P(
			h.AttributePairs(
				"id", "counter",
				"class", "text-xl",
				"name", "count",
				"text", "count",
			),
			h.TextF("Count: %d", count),
		),
		SubmitButton("Increment"),
	)
}

func SubmitButton(text string) *h.Element {
	return h.Button(
		h.Class("bg-rose-400 hover:bg-rose-500 text-white font-bold py-2 px-4 rounded"),
		h.Id("swap-text"),
		h.Type("submit"),
		h.Text(text),
	)
}
