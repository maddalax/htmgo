## Caching Components Per User

If you need to cache a component per user, you can use the `CachedPerKey` functions. 
These functions allow you to cache a component by a specific key. This key can be any string that uniquely identifies the user.

Note: I'm using the term 'user' to simply mean a unique identifier. This could be a user ID, session ID, or any other unique identifier.

To cache a component by unique identifier / key in htmgo, we offer:

```go
// No arguments passed to the component, the component can be cached by a specific key
h.CachedPerKey(duration time.Duration, cb GetElementFuncWithKey)
// One argument passed to the component, the component can be cached by a specific key
h.CachedPerKeyT1(duration time.Duration, cb GetElementFuncWithKey)
// Two argument passed to the component, the component can be cached by a specific key
h.CachedPerKeyT2(duration time.Duration, cb GetElementFuncWithKey)
// Three arguments passed to the component, the component can be cached by a specific key
h.CachedPerKeyT3(duration time.Duration, cb GetElementFuncWithKey)
// Four arguments passed to the component, the component can be cached by a specific key
h.CachedPerKeyT4(duration time.Duration, cb GetElementFuncWithKey)

```

The `duration` parameter is the time the component should be cached for. The `cb` parameter is a function that returns the component and the key.

When a request is made for a cached component, the component is rendered and stored in memory. Subsequent requests for the same component with the same key within the cache duration will return the cached component instead of rendering it again.

**Usage:**

```go
var CachedUserDocuments = h.CachedPerKeyT(time.Minute*15, func(ctx *h.RequestContext) (string, h.GetElementFunc) {
	userId := getUserIdFromSession(ctx)
	return userId, func() *h.Element {
		return UserDocuments(ctx)
	}
})

func UserDocuments(ctx *h.RequestContext) *h.Element {
	docService := NewDocumentService(ctx)
	// Expensive call
	docs := docService.getDocuments()
	return h.Div(
		h.Class("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4"),
		h.List(docs, func(doc Document, index int) *h.Element {
			return h.Div(
				h.Class("p-4 bg-white border border-gray-200 rounded-md"),
				h.H3(doc.Title),
				h.P(doc.Description),
			)
		}),
	)
}

func MyPage(ctx *h.RequestContext) *h.Page {
	
	// Note this is not a real way to create a context, just an example
	user1 := &h.RequestContext{
		Session: "user_1_session",
    }
	
	user2 := &h.RequestContext{
		Session: "user_2_session",
	}
	
	// Different users will get different cached components
    return h.NewPage(
        CachedUserDocuments(user1),
        CachedUserDocuments(user2),
    )
}

```

**Note:** We are using CachedPerKeyT because the component takes one argument, the RequestContext.
If the component takes more arguments, use CachedPerKeyT2, CachedPerKeyT3, etc.

**Important** 
1. The cached value is stored globally in memory by key, it is shared across all requests. Ensure if you are storing request-specific data in a cached component, you are using a unique key for each user.
2. The arguments passed into cached component **DO NOT** affect the cache key. The only thing that affects the cache key is the key returned by the `GetElementFuncWithKey` function.
3. Ensure the declaration of the cached component is **outside the function** that uses it. This is to prevent the component from being redeclared on each request.
