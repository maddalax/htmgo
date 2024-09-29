**Events**

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

If you use the OnEvent directly, event names may be any [HTML DOM](https://www.w3schools.com/jsref/dom_obj_event.asp) events, or any [HTMX events](https://htmx.org/events/).

Commands:

```go
js.AddAttribute(string, value)
js.RemoveAttribute(string)
js.AddClass(string, value)
js.SetText(string)
js.Increment(count)
js.SetInnerHtml(Ren)
js.SetOuterHtml(Ren)
js.SetDisabled(bool)
js.RemoveClass(string)
js.Alert(string)
js.EvalJs(string) // eval arbitrary js, use 'self' to get the current element as a reference
js.InjectScript(string)
js.InjectScriptIfNotExist(string)
js.GetPartial(PartialFunc)
js.GetPartialWithQs(PartialFunc, Qs)
js.PostPartial(PartialFunc)
js.PostPartialWithQs(PartialFunc, Qs)
js.GetWithQs(string, Qs)
js.PostWithQs(string, Qs)
js.ToggleClass(string)
js.ToggleClassOnElement(string, string)

// The following methods are used to evaluate JS on nearby elements. 
// Use 'element' to get the element as a reference for the EvalJs methods.
js.EvalJsOnParent(string) 
js.EvalJsOnSibling(string, string)
js.EvalJsOnChildren(string, string)
js.SetClassOnParent(string)
js.RemoveClassOnParent(string)
js.SetClassOnChildren(string, string)
js.RemoveClassOnChildren(string, string)
js.SetClassOnSibling(string, string)
js.RemoveClassOnSibling(string, string)

```
For more usages: see https://github.com/maddalax/htmgo/blob/master/htmgo-site/pages/form.go


**Example:** Evaluating arbitrary JS

```go
func MyButton() *h.Element {
	return h.Button(
		h.Text("Submit"),
		h.OnClick(
        // make sure you use 'self' instead of 'this' 
        // for referencing the current element
			h.EvalJs(`
				   if(Math.random() > 0.5) {
				       self.innerHTML = "Success!";
				   }
		       `,
			),
		),
	)
}
```

tip: If you are using Jetbrains IDE's, you can write `// language=js` as a comment above the function call (h.EvalJS) and it will automatically give you syntax highlighting on the raw JS.

```go
// language=js
h.EvalJs(`
     if(Math.random() > 0.5) {
         self.innerHTML = "Success!";
     }
     `,
),
```

