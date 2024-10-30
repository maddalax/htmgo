package htmx_extensions

import (
	"github.com/maddalax/htmgo/framework/h"
	. "htmgo-site/pages/docs2"
	"htmgo-site/ui"
)

func MutationError(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Mutation Error"),
			Text(`
				The 'mutation-error' extension allows you to trigger an event when a request returns a >= 400 status code.
				This is useful for things such as letting a child element (such as a button) inside a form know there was an error.
			`),
			Text(`<b>Example:</b>`),
			ui.GoCodeSnippet(MutationErrorExample),
			Text(`It can also be used on children elements that do not make an xhr request, if you combine it with the TriggerChildren extension.`),
			NextStep(
				"mt-4",
				PrevBlock("HTMX Extensions", "/docs2/htmx-extensions/overview"),
				NextBlock("Tailwind Intellisense", "/docs2/misc/tailwind-intellisense"),
			),
		),
	)
}

const MutationErrorExample = `
h.Form(
    h.HxTriggerChildren(),
    h.HxMutationError(
        js.Alert("An error occurred"),
    ),
    h.Button(
        h.Type("submit"),
        h.Text("Submit"),
    ),
)
`
