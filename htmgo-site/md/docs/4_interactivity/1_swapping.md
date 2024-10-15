## Interactivity / Swapping

1. Adding interactivity to your website is done through [htmx](http://htmx.org) by utilizing various attributes/headers. This should cover most use cases.
   htmgo offers utility methods to make this process a bit easier

Here are a few methods we offer:

Partial Response methods

```go
SwapManyPartialWithHeaders(ctx *RequestContext, headers *Headers, swaps ...*Element) *Partial
SwapPartial(ctx *RequestContext, swap *Element) *Partial
SwapManyPartial(ctx *RequestContext, swaps ...*Element) *Partial
SwapManyXPartial(ctx *RequestContext, swaps ...SwapArg) *Partial
GetPartialPath(partial PartialFunc) string
GetPartialPathWithQs(partial PartialFunc, qs *Qs) string
```

Swapping can also be done by adding a child to an element

```go
OobSwapWithSelector(ctx *RequestContext, selector string, content *Element, option ...SwapOption) *Element
OobSwap(ctx *RequestContext, content *Element, option ...SwapOption) *Element
SwapMany(ctx *RequestContext, elements ...*Element)
```



Usage:

1. I have a Card component that renders a list of tasks. I want to add a new button that completes all the tasks and updates the Card component with the completed tasks.
   

**/components/task.go**

```go
func Card(ctx *h.RequestContext) *h.Element {
	service := tasks.NewService(ctx.ServiceLocator())
	list, _ := service.List()

	return h.Div(
		h.Id("task-card"),
		h.Class("bg-white w-full rounded shadow-md"),
		CardBody(list, getActiveTab(ctx)),
    CompleteAllButton(list)
	)
}
```

```go
func CompleteAllButton(list []*ent.Task) *h.Element {
	notCompletedCount := len(h.Filter(list, func(item *ent.Task) bool {
		return item.CompletedAt == nil
	}))

	return h.Button(
		h.TextF("Complete %s tasks", notCompletedCount),
		h.PostPartialWithQs(CompleteAll,
			h.NewQs("complete",
				h.Ternary(notCompletedCount > 0, "true", "false"),
			)),
	)
}
```

**/partials/task.go**

```go
func CompleteAll(ctx *h.RequestContext) *h.Partial {
	service := tasks.NewService(ctx.ServiceLocator())
	service.SetAllCompleted(ctx.QueryParam("complete") == "true")
	return h.SwapPartial(ctx,
		Card(ctx),
	)
}
```

When the **CompleteAll** button is clicked, a **POST** will be sent to the **CompleteAll** partial, which will complete all the tasks and then swap out the Card content with the updated list of tasks. Pretty cool right?

**SwapManyPartial** can be used to swap out multiple items on the page instead of a single one.

Note: These partial swap methods use https://htmx.org/attributes/hx-swap-oob/ behind the scenes, so it must match 
the swap target by id.

**If** you are only wanting to swap the element that made the xhr request for the partial in the first place, just use `h.NewPartial` instead, it will use the default htmx swapping, and not hx-swap-oob.
