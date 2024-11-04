package interactivity

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/ui"
)

import . "htmgo-site/pages/docs"

func EventsAndCommands(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Events Handler / Commands"),
			Text(`
				In some cases, you need to update elements client side without having to do a network call. 
				For this you generally have to target an element with javascript and set an attribute, change the innerHTML, etc.
	      To make this work while still keeping a pure go feel, htmgo offers a few utility methods to execute various javascript on an element.
			`),
			Text("Example: When the form is submitted, set the button text to submitting and disable it, and vice versa after submit is done."),
			ui.GoCodeSnippet(EventsExample1),
			Text(`
        The structure of this comes down to:
       	1. Add an event handler to the element
       	2. Add commands (found in the <b>'js'</b> package) as children to that event handler
			`),
			Text(`The current list of event handlers we have utility methods for so far are:`),
			ui.CodeSnippet(ui.CodeSnippetProps{
				Code:            CurrentHandlersSnippet,
				Lang:            "bash",
				HideLineNumbers: true,
			}),
			h.P(
				h.Text("If there is not an existing method for the event you need, you can use the h.OnEvent method to add a handler for any "),
				Link("DOM event", "https://www.w3schools.com/jsref/dom_obj_event.asp"),
				h.Text(" or "),
				Link("htmx event.", "https://htmx.org/events/"),
			),
			Text("If there is not an existing method for the event you need, you can use the <b>h.OnEvent</b> method to add an event handler for any DOM or htmx event."),
			ui.GoCodeSnippet(OnEventBlurSnippet),
			h.P(
				h.Text(`For more details on how they work, see the source for `),
				Link("lifecycle.", "https://github.com/maddalax/htmgo/blob/master/framework/h/lifecycle.go"),
				h.Text(" Any method that returns *Lifecycle can be used as an event handler, and any method that returns *Command can be used as a command."),
			),
			h.P(
				h.Text(`The current list of commands supported can be found `),
				Link("here.", "https://github.com/maddalax/htmgo/blob/master/framework/js/commands.go"),
			),
			HelpText("Note: Each command you attach to the event handler will be passed 'self' and 'event' (if applicable) as arguments. self is the current element, and event is the event object."),
			Text("Example: Evaluating arbitrary Javascript"),
			ui.GoCodeSnippet(EvalArbitraryJavascriptSnippet),
			HelpText("Tips: If you are using Jetbrains IDE's, you can write '// language=js' as a comment above the function call (h.EvalJS) and it will automatically give you syntax highlighting on the raw JS."),
			h.P(
				h.Text("More examples and usage can be found on the "),
				Link("examples page, ", "/examples/js-set-text-on-click"),
				h.Text("in the 'Interactivity' section."),
			),
			NextStep(
				"mt-4",
				PrevBlock("Swapping", DocPath("/interactivity/swapping")),
				NextBlock("Eval Commands", DocPath("/interactivity/eval-commands")),
			),
		),
	)
}

const EventsExample1 = `
func MyForm() *h.Element {
	return h.Form(
		h.Button(
			h.Text("Submit"),
			h.HxBeforeRequest(
				js.SetDisabled(true),
				js.SetText("Submitting..."),
			),
			h.HxAfterRequest(
				js.SetDisabled(false),
				js.SetText("Submit"),
			),
		),
	)
}
`

var OnEventBlurSnippet = `
h.Input(
	h.OnEvent(
		hx.BlurEvent,
		js.SetValue("Input was blurred"),
	)
)`

var EvalArbitraryJavascriptSnippet = fmt.Sprintf(`func MyButton() *h.Element {
	return h.Button(
		h.Text("Submit"),
		h.OnClick(
        	// make sure you use 'self' instead of 'this' for referencing the current element
			h.EvalJs(%s
					if(Math.random() > 0.5) {
					self.innerHTML = "Success!";
				}%s
			),
		),
	)
}`, "`", "`")

const CurrentHandlersSnippet = `
h.OnEvent
h.OnLoad
h.HxBeforeRequest
h.HxOnLoad
h.HxOnAfterSwap
h.OnClick
h.OnSubmit
h.HxBeforeSseMessage
h.HxAfterSseMessage
h.HxOnSseError
h.HxOnSseClose
h.HxOnSseConnecting
h.HxOnSseOpen
h.HxAfterRequest
h.HxOnMutationError
`
