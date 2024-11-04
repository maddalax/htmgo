package pushing_data

import (
	"github.com/maddalax/htmgo/framework/h"
	. "htmgo-site/pages/docs"
	"htmgo-site/ui"
)

func ServerSentEvents(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Server Sent Events (SSE)"),
			Text(`
				htmgo supports server-sent events (SSE) out of the box. 
				This allows you to push data from the server to the client in real-time. 
			`),
			h.P(
				h.Text("Example of this can be found in the "),
				Link("examples/chat", "examples/chat"),
				h.Text(" project."),
			),
			SubTitle("How it works"),
			Text(`1. The client sends a request to the server to establish a connection.
	      2. The server holds the connection open and sends data (in our case, most likely elements) to the client whenever there is new data to send.
        3. The htmgo SSE extension uses hx-swap-oob to swap out the elements that the server sends.
			`),
			HelpText("Note: SSE is unidirectional (the server can only send data to the client). For the client to send data to the server, normal xhr behavior should be used (form submission, triggers, etc)."),
			Text(`<b>Usage:</b>`),
			Text("Add the SSE connection attribute and the path to the handler that will handle the connection."),
			ui.GoCodeSnippet(SseConnectAttribute),
			Text("The following <b>Event Handlers</b> can be used to react to SSE connections."),
			ui.GoCodeSnippet(SseEventHandlers),
			Text("Example: Adding an event listener handle SSE errors."),
			ui.GoCodeSnippet(SseErrorHandlingExample),
			Text("Example: Clearing the input field after sending a message."),
			ui.GoCodeSnippet(SseClearInputExample),
			NextStep(
				"mt-4",
				PrevBlock("Caching Per Key", DocPath("/performance/caching-per-key")),
				NextBlock("HTMX extensions", DocPath("/htmx-extensions/overview")),
			),
		),
	)
}

const SseConnectAttribute = `
h.Attribute("sse-connect", fmt.Sprintf("/chat/%s", roomId))
`

const SseEventHandlers = `
h.HxOnSseOpen
h.HxBeforeSseMessage
h.HxAfterSseMessage
h.HxOnSseError
h.HxOnSseClose
h.HxOnSseConnecting
`

const SseErrorHandlingExample = `
h.HxOnSseError(
    js.EvalJs(fmt.Sprintf("
			const reason = e.detail.event.data
			if(['invalid room', 'no session', 'invalid user'].includes(reason)) {
				window.location.href = '/?roomId=%s';
			} else if(e.detail.event.code === 1011) {
				window.location.reload()
			} else if (e.detail.event.code === 1008 || e.detail.event.code === 1006) {
				window.location.href = '/?roomId=%s';
			} else {
				console.error('Connection closed:', e.detail.event)
			}
", roomId, roomId)),
),
`

const SseClearInputExample = `
func MessageInput() *h.Element {
	return h.Input("text",
		h.Id("message-input"),
		h.Required(),
		h.HxAfterSseMessage(
			js.SetValue(""),
		),
	)
}`
