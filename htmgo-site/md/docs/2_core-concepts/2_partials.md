## Partials ##

Partials are where things get interesting. Partials allow you to start adding interactivity to your website by swapping in content, setting headers, redirecting, etc.

Partials have a similar structure to pages. A simple partial may look like:

```go
func CurrentTimePartial(ctx *h.RequestContext) *h.Partial {
	now := time.Now()
	return h.NewPartial(
		h.Div(
			h.Pf("The current time is %s", now.Format(time.RFC3339)),
		),
	)
}
```

This will get automatically registered in the same way that pages are registered, based on the file path. This allows you to reference partials directly via the function itself when rendering them, instead of worrying about the route.

**Example:**
I want to build a page that renders the current time, updating every second. Here is how that may look:

<br>

**pages/time.go**

```go
package pages

func CurrentTimePage(ctx *h.RequestContext) *h.Page {
	return h.NewPage(
		base.RootPage(
			h.GetPartial(
				partials.CurrentTimePartial,
				"load, every 1s"),
		))
}
```

**partials/time.go**

```go
package partials

func CurrentTimePartial(ctx *h.RequestContext) *h.Partial {
	now := time.Now()
	return h.NewPartial(
		h.Div(
			h.Pf("The current time is %s", now.Format(time.RFC3339)),
		),
	)
}
```

When the page load, the partial will be loaded in via htmx, and then swapped in every 1 second. With this
little amount of code and zero written javascript, you have a page that shows the current time and updates
every second.

