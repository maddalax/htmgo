package interactivity

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/ui"
)

import . "htmgo-site/pages/docs"

func EvalCommands(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Eval Commands"),
			Text(`
				Now that we've learned how about events/commands, I want to highlight a few useful commands.
		    One in particular is EvalCommands, which allows you to evaluate any command against any element just by referencing it in Go.
			`),
			SubTitle("Referencing an element directly"),
			Text("Example: Setting the text of an element on click of another element"),
			ui.GoCodeSnippet(EvalCommandsSnippet),
			Text(
				`We are calling <b>js.EvalCommands</b> with the text variable and the command to toggle the text of the element.
				This always you to run any commands against any element, without having to query for it via a selector.
			`,
			),
			h.P(
				h.A(
					h.Class("underline text-blue-500"),
					h.Href("/examples/js-hide-children-on-click"),
					h.Text("View the demo"),
				),
				h.Text(" for more details on what this could be used for."),
			),
			h.Div(
				h.Class("mt-4"),
			),
			SubTitle("Using a selector"),
			Text("If needed, you can query by selector"),
			ui.GoCodeSnippet(EvalCommandsSnippetWithSelector),
			NextStep(
				"mt-4",
				PrevBlock("Events / Commands", DocPath("/interactivity/events")),
				NextBlock("Caching Components", DocPath("/performance/caching-globally")),
			),
		),
	)
}

const EvalCommandsSnippetWithSelector = `	
func MyComponent(ctx *h.RequestContext) *h.Element {
	text := h.Pf("Text Before", h.Id("my-element"))
	return h.Div(
		h.Button(
			h.Text("Toggle Text"),
			h.OnClick(
				js.EvalCommandsOnSelector(
					"#my-element",
					js.ToggleText("Text Before", "Text After"),
				),
			),
		),
		text,
	)
}	
`

var EvalCommandsSnippet = `func MyComponent(ctx *h.RequestContext) *h.Element {
	text := h.Pf("Text Before")
	return h.Div(
		h.Button(
			h.Text("Toggle Text"),
			h.OnClick(
				js.EvalCommands(
					text,
					js.ToggleText("Text Before", "Text After"),
				),
			),
		),
		text,
	)
}`
