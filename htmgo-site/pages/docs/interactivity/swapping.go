package interactivity

import . "htmgo-site/pages/docs"
import "htmgo-site/ui"
import "github.com/maddalax/htmgo/framework/h"

func Swapping(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Swapping"),
			Text(`
				Swapping is the process of swapping out the content of an element with another element.
				This is the primary way htmgo allows you to add interactivity to your website through htmx.
			`),
			h.P(
				h.Text("The swapping examples below utilize "),
				Link("hx-swap-oob", "https://htmx.org/attributes/hx-swap-oob/"),
				h.Text(" behind the scenes to swap out the content of an element."),
			),
			Text("Example: A simple counter"),
			ui.GoCodeSnippet(SwapExample),
			Text(`
				In this example, when the form is submitted, an HTTP POST will be sent to the server and call <b>CounterPartial</b>.
				CounterPartial will then update the count and return it back to the client via <b>h.SwapManyPartial</b>.
				The h.SwapManyPartial function is a helper function that allows you to swap out multiple elements on the page.
			`),
			Text(`
				All the routing is handled behind the scenes by htmgo, so you can reference partials directly by their function reference, 
				instead of having to wire up routes for each partial.
			`),
			Text(`
          Sometimes you may need to pass additional information when calling the partial, such as an id of the current entity you are working with.
					This can be done by like so:
		   `),
			Text("Example: Getting the http path to the partial with extra qs parameters"),
			ui.GoCodeSnippet(SwapGetPartialPathWithQsExample),
			Text("Example: Posting to the partial path on blur"),
			ui.GoCodeSnippet(SwapGetPartialPathExampleOnBlur),
			h.P(
				h.Text("Note: if your swapping is not working as expected, make sure the element you are swapping has an id and it matches. "),
				h.Text("For further details on how oob works behind the scenes, see the "),
				Link("hx-swap-oob", "https://htmx.org/attributes/hx-swap-oob/"),
				h.Text(" docs."),
			),
			NextStep(
				"mt-4",
				PrevBlock("Loops / Dealing With Lists", DocPath("/control/loops")),
				NextBlock("Events / Commands", DocPath("/interactivity/events")),
			),
		),
	)
}

const SwapExample = `
func CounterPartial(ctx *h.RequestContext) *h.Partial {
	count, _ := strconv.ParseInt(ctx.FormValue("count"), 10, 64)

	count++

	return h.SwapManyPartial(
		ctx,
		CounterForm(int(count)),
		h.ElementIf(count > 10, SubmitButton("New record!")),
	)
}

func CounterForm(count int) *h.Element {
	return h.Form(
		h.Id("counter-form"),
		h.PostPartial(CounterPartial),
		h.Input(
			"text",
			h.Class("hidden"),
			h.Value(count),
			h.Name("count"),
		),
		h.P(
			h.Id("counter"),
          	h.Name("count"),
			h.TextF("Count: %d", count),
		),
		h.Button(
			h.Type("submit"),
			h.Text("Increment"),
		),
	)
}
`

const SwapGetPartialPathWithQsExample = `
func MyComponent() *h.Element {
	return h.Div(
		h.GetPartialPathWithQs(
			CounterPartial,
			h.NewQs("count", count),
		),
	)
}
`

const SwapGetPartialPathExampleOnBlur = `
func MyComponent() *h.Element {
	path := h.GetPartialPath(CounterPartial)
	return h.Input(
		h.Post(path, hx.BlurEvent),
	)
}
`
