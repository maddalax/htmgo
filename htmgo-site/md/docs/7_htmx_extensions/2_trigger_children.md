### HTMX Extensions - Trigger Children

The `trigger-children` extension allows you to trigger an event on all children and siblings of an element.

This is useful for things such as:
1. Letting a child element (such as a button) inside a form know the form was submitted

<br>

**Example:** https://github.com/maddalax/htmgo/blob/master/htmgo-site/pages/form.go#L17

In this example: The trigger-children extension will trigger **hx-before-request** and **hx-after-request** 
on all children of the form when the form is submitted, and the button reacts to that by showing a loading state.
