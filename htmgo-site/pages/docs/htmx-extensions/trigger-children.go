package htmx_extensions

import (
	"github.com/maddalax/htmgo/framework/h"
	. "htmgo-site/pages/docs"
)

func TriggerChildren(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Trigger Children"),
			Text(`
				The 'trigger-children' extension allows you to trigger an event on all children and siblings of an element.
				This is useful for things such as letting a child element (such as a button) inside a form know the form was submitted
			`),
			Link("View Example", "https://htmgo.dev/examples/form"),
			HelpText(`In this example: The trigger-children extension will trigger hx-before-request and hx-after-request on all children of the form when the form is submitted, and the button reacts to that by showing a loading state.`),
			NextStep(
				"mt-4",
				PrevBlock("HTMX Extensions", DocPath("/htmx-extensions/overview")),
				NextBlock("Mutation Error", DocPath("/htmx-extensions/mutation-error")),
			),
		),
	)
}

const TriggerChildrenExample = `
func MyForm() *h.Element {
	return h.Form(
		h.Button(
			h.Text("Submit"),
		),
	)
}
`
