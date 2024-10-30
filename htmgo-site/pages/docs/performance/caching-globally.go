package performance

import (
	"github.com/maddalax/htmgo/framework/h"
	. "htmgo-site/pages/docs"
	"htmgo-site/ui"
)

func CachingGlobally(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Caching Components Globally"),
			Text(`
				You may want to cache components to improve performance. This is especially useful for components that are expensive to render or make external requests for data.

				When a request is made for a cached component, the component is rendered and stored in memory. Subsequent requests for the same component within the cache duration will return the cached component instead of rendering it again.

				To cache a component in htmgo, we offer two ways, caching globally or caching per key, this section will focus on caching globally, you will learn more about caching per key in the next section:
			`),
			Text("<b>Methods for caching globally:</b>"),
			ui.GoCodeSnippet(CachingMethods),
			h.P(
				h.Text("For caching components per unique identifier, see "),
				Link("Caching Components Per Key", "/docs/performance/caching-per-key"),
				h.Text("."),
			),
			Text(`<b>Usage:</b>`),
			ui.GoCodeSnippet(CachedGloballyExample),
			Text(`
				We are using CachedT because the component takes one argument, the RequestContext.
				If the component takes more arguments, use CachedT2, CachedT3, etc.
			`),
			Text(
				`<b>Important Note:</b> When using h.CachedT and not <b>CachedPerKey</b>, the cached value is stored globally in memory, so it is shared across all requests. 
				      Do not store request-specific data in a cached component. Only cache components that you are OK with all users seeing the same data.

							The arguments passed into cached component <b>DO NOT</b> affect the cache key. You will get the same cached component regardless of the arguments passed in. This is different from what you may be used to from something like React useMemo.

							Ensure the declaration of the cached component is outside the function that uses it. This is to prevent the component from being redeclared on each request.
			`),
			NextStep(
				"mt-4",
				PrevBlock("Events", DocPath("/interactivity/events")),
				NextBlock("Caching Per Key", DocPath("/performance/caching-per-key")),
			),
		),
	)
}

const CachingMethods = `
// No arguments passed to the component
h.Cached(duration time.Duration, cb GetElementFunc)
// One argument passed to the component
h.CachedT(duration time.Duration, cb GetElementFunc)
// Two arguments passed to the component
h.CachedT2(duration time.Duration, cb GetElementFunc)
// Three arguments passed to the component
h.CachedT3(duration time.Duration, cb GetElementFunc)
// Four arguments passed to the component
h.CachedT4(duration time.Duration, cb GetElementFunc)
`

const CachedGloballyExample = `
func ExpensiveComponent(ctx *h.RequestContext) *h.Element { 
  // Some expensive call
  data := http.Get("https://api.example.com/data")	
  return h.Div(
    h.Text(data),
  )
}

var CachedComponent = h.CachedT(time.Minute*15, func(ctx *h.RequestContext) *h.Element {
  return ExpensiveComponent(ctx)
})

func IndexPage(ctx *h.RequestContext) *h.Page {
  return h.NewPage(
    CachedComponent(ctx),
  )
}
`
