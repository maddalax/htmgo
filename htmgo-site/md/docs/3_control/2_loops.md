**Loops / Dealing With Lists**

Very commonly you will need to render a list or slice of items onto the page. Frameworks generally solve this in different ways, such as React uses regular JS .map function to solve it.

We offer the same conveniences in htmgo.

```go
h.List(items, func(item, index)) *h.Element
h.IterMap(map, mapper func(key, value) *Element) *Element 
```

**Example:**

- Render a list of tasks

```go
h.List(list, func(item *ent.Task, index int) *h.Element {
    if tab == TabComplete && item.CompletedAt == nil {
       return h.Empty()
    }
    return Task(item, false)
})
```

- Render a map

```go
  values := map[string]string{
  		"key": "value",
  	}
  
  	IterMap(values, func(key string, value string) *Element {
  		return Div(
  			Text(key),
  			Text(value),
  		)
  	})
  ```
- 



