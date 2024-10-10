### HTMX Extensions - Mutation Error

The `mutation-error` extension allows you to trigger an event when a request returns a >= 400 status code.

This is useful for things such as:
1. Letting a child element (such as a button) inside a form know there was an error.

<br>

**Example:**
```go
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
```

It can also be used on children elements that do not make an xhr request, if you combine it with the `hx-trigger-children` extension.
