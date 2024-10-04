**Server Sent Events (SSE)**

htmgo supports server-sent events (SSE) out of the box. 
This allows you to push data from the server to the client in real-time. 

Example of this can be found in the [chat-app](https://github.com/maddalax/htmgo/tree/master/examples/chat) example.
Demo: https://chat-example.htmgo.dev

## How it works ##
1. The client sends a request to the server to establish a connection.
2. The server holds the connection open and sends data (in our case, most likely elements) to the client whenever there is new data to send.
3. The htmgo SSE extension uses https://htmx.org/attributes/hx-swap-oob/ to swap out the elements that the server sends.


**Note**: SSE is **unidirectional** (the server can only send data to the client).
For the client to send data to the server, normal xhr behavior should be used (form submission, triggers, etc).

## Usage
1. Add the SSE connection attribute and the path to the handler that will handle the connection.

```go
h.Attribute("sse-connect", fmt.Sprintf("/chat/%s", roomId))
```

The following **Event Handlers** can be used to react to SSE connections.
```go
h.HxOnSseOpen
h.HxBeforeSseMessage
h.HxAfterSseMessage
h.HxOnSseError
h.HxOnSseClose
h.HxOnSseConnecting
```

**Example:** Adding an event listener handle SSE errors.

```go
h.HxOnSseError(
    js.EvalJs(fmt.Sprintf(`
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
    `, roomId, roomId)),
),
```

**Example:** Clearing the input field after sending a message.
```go
func MessageInput() *h.Element {
	return h.Input("text",
		h.Id("message-input"),
		h.Required(),
		h.HxAfterSseMessage(
			js.SetValue(""),
		),
	)
}
```
