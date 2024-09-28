**Caching Components**

You may want to cache components to improve performance. This is especially useful for components that are expensive to render 
or make external requests for data.

To cache a component in htmgo, we offer:

```go
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
```

The `duration` parameter is the time the component should be cached for. The `cb` parameter is a function that returns the component.

When a request is made for a cached component, the component is rendered and stored in memory. Subsequent requests for the same component within the cache duration will return the cached component instead of rendering it again.

**Usage:**

```go
func ExpensiveComponent(ctx *h.RequestContext) *h.Element { 
  // Some expensive call
  data := http.Get("https://api.example.com/data")	
  return h.Div(
    h.Text(data),
  )
}

var CachedComponent = h.CachedT(5*time.Minute, func(ctx *h.RequestContext) *h.Element {
  return ExpensiveComponent(ctx)
})

func IndexPage(ctx *h.RequestContext) *h.Page {
  return h.NewPage(
    CachedComponent(ctx),
  )
}
```

**Real Example:**
I want to make a navbar that renders how many github stars my repository has. I don't want to make a request to the GitHub API everytime someone visits my page, so I will cache the component for 15 minutes.
```go
var CachedGithubStars = h.CachedT(time.Minute*15, func(t *h.RequestContext) *h.Element {
    return GithubStars(t)
})

func GithubStars(ctx *h.RequestContext) *h.Element {
  stars := http.Get("https://api.github.com/repos/maddalax/htmgo/stargazers")
  return h.Div(
    h.Text(stars),
  )
}
```

**Note:** We are using CachedT because the component takes one argument, the RequestContext.
If the component takes more arguments, use CachedT2, CachedT3, etc.

**Important**: 
1. The cached value is stored globally in memory, so it is shared across all requests. Do not store request-specific data in a cached component. Only cache components that you are OK with all users seeing the same data.
2. The arguments passed into cached component **DO NOT** affect the cache key. You will get the same cached component regardless of the arguments passed in. This is different from what you may be used to from something like React useMemo.
3. Ensure the declaration of the cached component is **outside the function** that uses it. This is to prevent the component from being redeclared on each request.
