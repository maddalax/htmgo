## Events Handlers / Commands

Sometimes you need to update elements client side without having to do a network call. For this you generally have to target an element with javascript and set an attribute, change the innerHTML, etc.

To make this work while still keeping a pure go feel, we offer a few utility methods to execute various javascript on an element.

**Example:** When the form is submitted, set the button text to submitting and disable it, and vice versa after submit is done.

```go
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
```

The structure of this comes down to:

1. Add an event handler to the element
2. Add commands (found in the **js** package) as children to that event handler

<br>

**Event Handlers:**

```go
HxBeforeRequest(cmd ...Command) *LifeCycle
HxAfterRequest(cmd ...Command) *LifeCycle
HxOnMutationError(cmd ...Command) *LifeCycle
OnEvent(event hx.Event, cmd ...Command) *LifeCycle
OnClick(cmd ...Command) *LifeCycle
HxOnAfterSwap(cmd ...Command) *LifeCycle
HxOnLoad(cmd ...Command) *LifeCycle
```

